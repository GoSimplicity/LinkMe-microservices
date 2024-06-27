package biz

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"linkme-user/domain"
	"linkme-user/internal/data"
)

type UserUsecase struct {
	repo data.UserRepo
	l    *zap.Logger
}

func NewUserUsecase(repo data.UserRepo, l *zap.Logger) *UserUsecase {
	return &UserUsecase{
		l:    l,
		repo: repo,
	}
}

func (uu *UserUsecase) SignUp(ctx context.Context, u domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		uu.l.Error("password conversion failed")
		return err
	}
	u.Password = string(hash)
	return uu.repo.CreateUser(ctx, u)
}

func (uu *UserUsecase) Login(ctx context.Context, email string, password string) (domain.User, error) {
	u, err := uu.repo.FindByEmail(ctx, email)
	// 如果用户没有找到(未注册)，则返回空对象
	if errors.Is(err, data.ErrUserNotFound) {
		return domain.User{}, err
	} else if err != nil {
		uu.l.Error("user not found", zap.Error(err))
		return domain.User{}, err
	}
	// 将密文密码转为明文
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		uu.l.Error("password conversion failed", zap.Error(err))
		return domain.User{}, data.ErrInvalidUserOrPassword
	}
	return u, nil
}

func (uu *UserUsecase) ChangePassword(ctx context.Context, email string, password string, confirmPassword string) error {
	u, err := uu.repo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, data.ErrUserNotFound) {
			return err
		}
		uu.l.Error("failed to find user", zap.Error(err))
		return err
	}
	if er := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); er != nil {
		uu.l.Error("password verification failed", zap.Error(er))
		return data.ErrInvalidUserOrPassword
	}
	newHash, err := bcrypt.GenerateFromPassword([]byte(confirmPassword), bcrypt.DefaultCost)
	if err != nil {
		uu.l.Error("failed to hash new password", zap.Error(err))
		return err
	}
	if er := uu.repo.ChangePassword(ctx, email, string(newHash)); er != nil {
		uu.l.Error("failed to change password", zap.Error(er))
		return er
	}
	return nil
}

func (uu *UserUsecase) DeleteUser(ctx context.Context, email string, password string, uid int64) error {
	u, err := uu.repo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, data.ErrUserNotFound) {
			return err
		}
		uu.l.Error("failed to find user", zap.Error(err))
		return err
	}
	if er := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); er != nil {
		uu.l.Error("password verification failed", zap.Error(er))
		return data.ErrInvalidUserOrPassword
	}
	err = uu.repo.DeleteUser(ctx, email, uid)
	if err != nil {
		uu.l.Error("failed to delete user", zap.Error(err))
		return err
	}
	return nil
}

func (uu *UserUsecase) UpdateProfile(ctx context.Context, profile domain.Profile) (err error) {
	return uu.repo.UpdateProfile(ctx, profile)
}

func (uu *UserUsecase) GetProfileByUserID(ctx context.Context, UserID int64) (profile domain.Profile, err error) {
	return uu.repo.GetProfile(ctx, UserID)
}
