package service

import (
	"context"
	"errors"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	pb "linkme-user/api/user/v1"
	"linkme-user/domain"
	"linkme-user/internal/biz"
	"linkme-user/internal/data"
)

const (
	emailRegexPattern    = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)

type UserService struct {
	pb.UnimplementedUserServer
	uc       *biz.UserUsecase
	Email    *regexp.Regexp
	PassWord *regexp.Regexp
	ijwt     data.Handler
}

func NewUserService(uc *biz.UserUsecase, ijwt data.Handler) *UserService {
	return &UserService{
		uc:       uc,
		Email:    regexp.MustCompile(emailRegexPattern, regexp.None),
		PassWord: regexp.MustCompile(passwordRegexPattern, regexp.None),
		ijwt:     ijwt,
	}
}

func (s *UserService) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpReply, error) {
	emailValid, err := s.Email.MatchString(req.Email)
	if err != nil {
		return &pb.SignUpReply{
			Code: 1,
			Msg:  "User registration failed",
		}, err
	}
	if !emailValid {
		return &pb.SignUpReply{
			Code: 1,
			Msg:  "Invalid email format, please check",
		}, nil
	}
	// 验证密码是否一致
	if req.Password != req.ConfirmPassword {
		return &pb.SignUpReply{
			Code: 1,
			Msg:  "The two passwords entered are different, please re-enter",
		}, nil
	}
	// 验证密码格式
	passwordValid, err := s.PassWord.MatchString(req.Password)
	if err != nil {
		return &pb.SignUpReply{
			Code: 1,
			Msg:  "User registration failed",
		}, err
	}
	if !passwordValid {
		return &pb.SignUpReply{
			Code: 1,
			Msg:  "Password must contain letters, numbers, and special characters, and be at least 8 characters long",
		}, nil
	}
	// 尝试注册用户
	err = s.uc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		// 检查是否为重复邮箱错误
		if errors.Is(err, data.ErrDuplicateEmail) {
			return &pb.SignUpReply{
				Code: 1,
				Msg:  "Email already exists",
			}, nil
		}
		return &pb.SignUpReply{
			Code: 1,
			Msg:  "User registration failed",
		}, err
	}
	return &pb.SignUpReply{
		Code: 0,
		Msg:  "User registration succeeded",
	}, nil
}
func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	_, err := s.uc.Login(ctx, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, data.ErrInvalidUserOrPassword) {
			return &pb.LoginReply{
				Code: 1,
				Msg:  "username or password is incorrect",
			}, err
		}
		return &pb.LoginReply{
			Code: 1,
			Msg:  "User login failed",
		}, err
	}
	return &pb.LoginReply{
		Code: 0,
		Msg:  "User login successful",
	}, nil
}

func (s *UserService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	ginCtx := ctx.(*gin.Context)
	// 清除JWT令牌
	if err := s.ijwt.ClearToken(ginCtx); err != nil {
		return &pb.LogoutReply{
			Code: 1,
			Msg:  "User logout failed",
		}, err
	}
	return &pb.LogoutReply{
		Code: 0,
		Msg:  "User logout successful",
	}, nil
}
func (s *UserService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenReply, error) {
	var rc data.RefreshClaims
	ginCtx := ctx.(*gin.Context)
	// 从前端的Authorization中取出token
	tokenString := s.ijwt.ExtractToken(ginCtx)
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &rc, func(token *jwt.Token) (interface{}, error) {
		return data.Key2, nil
	})
	if err != nil {
		return &pb.RefreshTokenReply{
			Code: 1,
			Msg:  "User refresh token failed",
		}, err
	}
	if token == nil || !token.Valid {
		return &pb.RefreshTokenReply{
			Code: 1,
			Msg:  "User refresh token failed",
		}, err
	}
	// 检查会话状态是否异常
	if err = s.ijwt.CheckSession(ginCtx, rc.Ssid); err != nil {
		return &pb.RefreshTokenReply{
			Code: 1,
			Msg:  "User refresh token failed",
		}, err
	}
	// 刷新短token
	if err = s.ijwt.SetJWTToken(ginCtx, rc.Uid, rc.Ssid); err != nil {
		return &pb.RefreshTokenReply{
			Code: 1,
			Msg:  "User refresh token failed",
		}, err
	}
	return &pb.RefreshTokenReply{
		Code: 0,
		Msg:  "User refresh token successful",
	}, nil
}
func (s *UserService) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordReply, error) {
	// 检查新密码和确认密码是否匹配
	ginCtx := ctx.(*gin.Context)
	if req.Password != req.ConfirmPassword {
		return &pb.ChangePasswordReply{
			Code: 1,
			Msg:  "The two passwords entered are different, please re-enter",
		}, nil
	}
	err := s.uc.ChangePassword(ginCtx.Request.Context(), req.Email, req.Password, req.ConfirmPassword)
	if err != nil {
		if errors.Is(err, data.ErrInvalidUserOrPassword) {
			return &pb.ChangePasswordReply{
				Code: 1,
				Msg:  "username or password is incorrect",
			}, err
		}
		return &pb.ChangePasswordReply{
			Code: 1,
			Msg:  "User change password failed",
		}, err
	}
	return &pb.ChangePasswordReply{
		Code: 0,
		Msg:  "User change password successful",
	}, nil
}
func (s *UserService) WriteOff(ctx context.Context, req *pb.WriteOffRequest) (*pb.WriteOffReply, error) {
	ginCtx := ctx.(*gin.Context)
	uc := ginCtx.MustGet("user").(data.UserClaims)
	err := s.uc.DeleteUser(ctx, req.Email, req.Password, uc.Uid)
	if err != nil {
		return &pb.WriteOffReply{
			Code: 1,
			Msg:  "User write off failed",
		}, err
	}
	if err = s.ijwt.ClearToken(ginCtx); err != nil {
		return &pb.WriteOffReply{
			Code: 1,
			Msg:  "User write off failed",
		}, err
	}
	return &pb.WriteOffReply{
		Code: 0,
		Msg:  "User write off successful",
	}, nil
}

func (s *UserService) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileReply, error) {
	ginCtx := ctx.(*gin.Context)
	uc := ginCtx.MustGet("user").(data.UserClaims)
	profile, err := s.uc.GetProfileByUserID(ctx, uc.Uid)
	if err != nil {
		return &pb.GetProfileReply{
			Code: 1,
			Msg:  "User get profile failed",
		}, err
	}
	profileData := &pb.Profile{
		Nickname: profile.NickName,
		Avatar:   profile.Avatar,
		About:    profile.About,
		Birthday: profile.Birthday,
	}
	return &pb.GetProfileReply{
		Code: 0,
		Msg:  "User get profile successful",
		Data: profileData,
	}, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileReply, error) {
	ginCtx := ctx.(*gin.Context)
	uc := ginCtx.MustGet("user").(data.UserClaims)
	if err := s.uc.UpdateProfile(ctx, domain.Profile{
		NickName: req.Nickname,
		Avatar:   req.Avatar,
		About:    req.About,
		Birthday: req.Birthday,
		UserID:   uc.Uid,
	}); err != nil {
		return &pb.UpdateProfileReply{
			Code: 1,
			Msg:  "Profile update failed",
		}, err
	}
	return &pb.UpdateProfileReply{
		Code: 0,
		Msg:  "Profile update successful",
	}, nil
}
