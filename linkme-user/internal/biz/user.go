package biz

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"linkme-user/domain"
)

var (
	ErrCodeDuplicateEmailNumber              uint16 = 1062
	ErrDuplicateEmail                               = errors.New("duplicate email")
	ErrUserNotFound                                 = errors.New("user not found")
	ErrInvalidUserOrPassword                        = errors.New("username or password is incorrect")
	ErrNewPasswordAndConfirmPasswordNotMatch        = errors.New("two passwords do not match")
)

type UserBiz struct {
	ui UserInteractive
	l  *zap.Logger
}

type UserInteractive interface {
	CreateUser(ctx context.Context, u domain.User) error
	FindByID(ctx context.Context, id int64) (domain.User, error)
	FindByPhone(ctx context.Context, phone string) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	ChangePassword(ctx context.Context, email string, newPassword string) error
	DeleteUser(ctx context.Context, email string, uid int64) error
	UpdateProfile(ctx context.Context, profile domain.Profile) error
	GetProfile(ctx context.Context, userId int64) (domain.Profile, error)
}

func NewUserBiz(ui UserInteractive, l *zap.Logger) *UserBiz {
	return &UserBiz{
		ui: ui,
		l:  l,
	}
}

func (uu *UserBiz) SignUp(ctx context.Context, u domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		uu.l.Error("password conversion failed")
		return err
	}
	u.Password = string(hash)
	return uu.ui.CreateUser(ctx, u)
}

func (uu *UserBiz) Login(ctx context.Context, email string, password string) (domain.User, error) {
	u, err := uu.ui.FindByEmail(ctx, email)
	// 如果用户没有找到(未注册)，则返回空对象
	if errors.Is(err, ErrUserNotFound) {
		return domain.User{}, err
	} else if err != nil {
		uu.l.Error("user not found", zap.Error(err))
		return domain.User{}, err
	}
	// 将密文密码转为明文
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		uu.l.Error("password conversion failed", zap.Error(err))
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, nil
}

func (uu *UserBiz) ChangePassword(ctx context.Context, email string, password string, newPassword string, confirmPassword string) error {
	u, err := uu.ui.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return err
		}
		uu.l.Error("failed to find user", zap.Error(err))
		return err
	}
	if newPassword != confirmPassword {
		uu.l.Error("new password and confirm password do not match")
		return ErrNewPasswordAndConfirmPasswordNotMatch
	}
	if er := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); er != nil {
		uu.l.Error("password verification failed", zap.Error(er))
		return ErrInvalidUserOrPassword
	}
	newHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		uu.l.Error("failed to hash new password", zap.Error(err))
		return err
	}
	if er := uu.ui.ChangePassword(ctx, email, string(newHash)); er != nil {
		uu.l.Error("failed to change password", zap.Error(er))
		return er
	}
	return nil
}

func (uu *UserBiz) DeleteUser(ctx context.Context, email string, password string, uid int64) error {
	u, err := uu.ui.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return err
		}
		uu.l.Error("failed to find user", zap.Error(err))
		return err
	}
	if er := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); er != nil {
		uu.l.Error("password verification failed", zap.Error(er))
		return ErrInvalidUserOrPassword
	}
	err = uu.ui.DeleteUser(ctx, email, uid)
	if err != nil {
		uu.l.Error("failed to delete user", zap.Error(err))
		return err
	}
	return nil
}

func (uu *UserBiz) UpdateProfile(ctx context.Context, profile domain.Profile) (err error) {
	return uu.ui.UpdateProfile(ctx, profile)
}

func (uu *UserBiz) GetProfileByUserID(ctx context.Context, UserID int64) (profile domain.Profile, err error) {
	return uu.ui.GetProfile(ctx, UserID)
}
