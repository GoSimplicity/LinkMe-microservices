package data

import (
	"context"
	"errors"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-user/domain"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-user/internal/biz"
	sf "github.com/bwmarrin/snowflake"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

// User 用户信息结构体
type User struct {
	ID           int64   `gorm:"primarykey"`                          // 用户ID，主键
	CreateTime   int64   `gorm:"column:created_at;type:bigint"`       // 创建时间，Unix时间戳
	UpdatedTime  int64   `gorm:"column:updated_at;type:bigint"`       // 更新时间，Unix时间戳
	DeletedTime  int64   `gorm:"column:deleted_at;type:bigint;index"` // 删除时间，Unix时间戳，用于软删除
	PasswordHash string  `gorm:"not null"`                            // 密码哈希值，不能为空
	Deleted      bool    `gorm:"column:deleted;default:false"`        // 删除标志，表示该用户是否被删除
	Email        string  `gorm:"type:varchar(100);uniqueIndex"`       // 邮箱地址，唯一
	Phone        *string `gorm:"type:varchar(15);uniqueIndex"`        // 手机号码，唯一
	Profile      Profile `gorm:"foreignKey:UserID;references:ID"`     // 关联的用户资料
}

// Profile 用户资料信息结构体
type Profile struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`         // 用户资料ID，主键
	UserID   int64  `gorm:"not null;index"`                   // 用户ID，外键
	NickName string `gorm:"size:50"`                          // 昵称，最大长度50
	Avatar   string `gorm:"type:text"`                        // 头像URL
	About    string `gorm:"type:text"`                        // 个人简介
	Birthday string `gorm:"column:birthday;type:varchar(10)"` // 生日
}

type UserData struct {
	data *Data
	l    *zap.Logger
	node *sf.Node
}

func NewUserData(data *Data, l *zap.Logger, node *sf.Node) biz.UserInteractive {
	return &UserData{
		data: data,
		l:    l,
		node: node,
	}
}

func (ud *UserData) currentTime() int64 {
	return time.Now().UnixMilli()
}

func (ud *UserData) CreateUser(ctx context.Context, u domain.User) error {
	user := toDataUser(u)
	user.CreateTime = ud.currentTime()
	user.UpdatedTime = user.CreateTime
	// 使用雪花算法生成id
	user.ID = ud.node.Generate().Int64()
	// 初始化用户资料
	profile := Profile{
		UserID:   user.ID,
		NickName: "",
		Avatar:   "",
		About:    "",
		Birthday: "",
	}
	// 开始事务
	tx := ud.data.db.WithContext(ctx).Begin()
	// 创建用户
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == biz.ErrCodeDuplicateEmailNumber {
			ud.l.Error("duplicate email error", zap.String("email", user.Email), zap.Error(err))
			return biz.ErrDuplicateEmail
		}
		ud.l.Error("failed to create user", zap.Error(err))
		return err
	}
	// 创建用户资料
	if err := tx.Create(&profile).Error; err != nil {
		tx.Rollback()
		ud.l.Error("failed to create profile", zap.Error(err))
		return err
	}
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		ud.l.Error("transaction commit failed", zap.Error(err))
		return err
	}
	return nil
}

func (ud *UserData) FindByID(ctx context.Context, id int64) (domain.User, error) {
	var user User
	err := ud.data.db.WithContext(ctx).Where("id = ? AND deleted = ?", id, false).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, biz.ErrUserNotFound
		}
		return domain.User{}, err
	}
	return toDomainUser(user), nil
}

func (ud *UserData) FindByPhone(ctx context.Context, phone string) (domain.User, error) {
	var user User
	err := ud.data.db.WithContext(ctx).Where("phone = ? AND deleted = ?", phone).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, biz.ErrUserNotFound
		}
		return domain.User{}, err
	}
	return toDomainUser(user), nil
}

func (ud *UserData) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	var user User
	err := ud.data.db.WithContext(ctx).Where("email = ? AND deleted = ?", email, false).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, biz.ErrUserNotFound
		}
		return domain.User{}, err
	}
	return toDomainUser(user), nil
}

func (ud *UserData) ChangePassword(ctx context.Context, email string, newPassword string) error {
	tx := ud.data.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		ud.l.Error("failed to begin transaction", zap.Error(tx.Error))
		return tx.Error
	}
	// 更新密码
	if err := tx.Model(&User{}).Where("email = ? AND deleted = ?", email, false).Update("password_hash", newPassword).Error; err != nil {
		ud.l.Error("update password failed", zap.String("email", email), zap.Error(err))
		if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
			ud.l.Error("failed to rollback transaction", zap.Error(rollbackErr))
		}
		return err
	}
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		ud.l.Error("failed to commit transaction", zap.Error(err))
		return err
	}
	ud.l.Info("password updated successfully", zap.String("email", email))
	return nil
}

func (ud *UserData) DeleteUser(ctx context.Context, email string, uid int64) error {
	tx := ud.data.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()
	if err := tx.Model(&User{}).Where("email = ? AND deleted = ? AND id = ?", email, false, uid).Update("deleted", true).Error; err != nil {
		tx.Rollback()
		ud.l.Error("failed to mark user as deleted", zap.String("email", email), zap.Error(err))
		return err
	}
	if err := tx.Commit().Error; err != nil {
		ud.l.Error("failed to commit transaction", zap.String("email", email), zap.Error(err))
		return err
	}
	ud.l.Info("user marked as deleted", zap.String("email", email))
	return nil
}

func (ud *UserData) UpdateProfile(ctx context.Context, p domain.Profile) error {
	profile := toDataProfile(p)
	updates := map[string]interface{}{
		"nick_name": profile.NickName,
		"avatar":    profile.Avatar,
		"about":     profile.About,
		"birthday":  profile.Birthday,
	}
	// 更新操作
	err := ud.data.db.WithContext(ctx).Model(&Profile{}).Where("user_id = ?", profile.UserID).Updates(updates).Error
	if err != nil {
		ud.l.Error("failed to update profile", zap.Error(err))
		return err
	}
	return nil
}

func (ud *UserData) GetProfile(ctx context.Context, userId int64) (domain.Profile, error) {
	var profile domain.Profile
	if err := ud.data.db.WithContext(ctx).Where("user_id = ?", userId).First(&profile).Error; err != nil {
		ud.l.Error("failed to get profile by user id", zap.Error(err))
		return domain.Profile{}, err
	}
	return profile, nil
}

func toDataUser(u domain.User) User {
	return User{
		ID:           u.ID,
		PasswordHash: u.Password,
		Email:        u.Email,
		Phone:        u.Phone,
		CreateTime:   u.CreateTime,
	}
}

func toDomainUser(u User) domain.User {
	return domain.User{
		ID:         u.ID,
		Password:   u.PasswordHash,
		Email:      u.Email,
		Phone:      u.Phone,
		CreateTime: u.CreateTime,
	}
}

func toDataProfile(profile domain.Profile) Profile {
	return Profile{
		ID:       profile.ID,
		UserID:   profile.UserID,
		NickName: profile.NickName,
		Avatar:   profile.Avatar,
		About:    profile.About,
		Birthday: profile.Birthday,
	}
}
