package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

var (
	Key1 = []byte("sadfkhjlkkljKFJDSLAFUDASLFJKLjfj113d1")
	Key2 = []byte("sadfkhjlkkljKFJDSLAFUDASLFJKLjfj113d2")
)

type Handler interface {
	SetLoginToken(ctx context.Context, uid int64) (string, string, error)
	SetJWTToken(ctx context.Context, uid int64, ssid string) (string, error)
	ExtractToken(ctx context.Context) string
	CheckSession(ctx context.Context, ssid string) error
	ClearToken(ctx context.Context) error
	setRefreshToken(ctx context.Context, uid int64, ssid string) (string, error)
}

type UserClaims struct {
	jwt.RegisteredClaims
	Uid         int64
	Ssid        string
	UserAgent   string
	ContentType string
}

type RefreshClaims struct {
	jwt.RegisteredClaims
	Uid  int64
	Ssid string
}

type handler struct {
	client        redis.Cmdable
	signingMethod jwt.SigningMethod
	rcExpiration  time.Duration
}

func NewJWT(c redis.Cmdable) Handler {
	return &handler{
		client:        c,
		signingMethod: jwt.SigningMethodHS512,
		rcExpiration:  time.Hour * 24 * 7,
	}
}

// SetLoginToken 设置长短Token
func (h *handler) SetLoginToken(ctx context.Context, uid int64) (string, string, error) {
	ssid := uuid.New().String()
	refreshToken, err := h.setRefreshToken(ctx, uid, ssid)
	if err != nil {
		return "", "", err
	}
	jwtToken, err := h.SetJWTToken(ctx, uid, ssid)
	if err != nil {
		return "", "", err
	}
	return jwtToken, refreshToken, nil
}

// SetJWTToken 设置短Token
func (h *handler) SetJWTToken(ctx context.Context, uid int64, ssid string) (string, error) {
	// 从 Kratos 上下文中获取传输信息
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return "", errors.New("failed to get transport from context")
	}
	uc := UserClaims{
		Uid:         uid,
		Ssid:        ssid,
		UserAgent:   tr.RequestHeader().Get("User-Agent"),
		ContentType: tr.RequestHeader().Get("Content-Type"),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
	}
	token := jwt.NewWithClaims(h.signingMethod, uc)
	// 进行签名
	signedString, err := token.SignedString(Key1)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

// setRefreshToken 设置长Token
func (h *handler) setRefreshToken(ctx context.Context, uid int64, ssid string) (string, error) {
	rc := RefreshClaims{
		Uid:  uid,
		Ssid: ssid,
		RegisteredClaims: jwt.RegisteredClaims{
			// 设置刷新时间为一周
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(h.rcExpiration)),
		},
	}
	t := jwt.NewWithClaims(h.signingMethod, rc)
	signedString, err := t.SignedString(Key2)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

// ExtractToken 提取 Authorization 头部中的 Token
func (h *handler) ExtractToken(ctx context.Context) string {
	// 从 Kratos 上下文中获取传输信息
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return ""
	}
	authCode := tr.RequestHeader().Get("Authorization")
	if authCode == "" {
		return ""
	}
	// Authorization 头部格式需为 Bearer string
	s := strings.Split(authCode, " ")
	if len(s) != 2 {
		return ""
	}
	return s[1]
}

// CheckSession 检查会话状态
func (h *handler) CheckSession(ctx context.Context, ssid string) error {
	// 判断缓存中是否存在指定键
	c, err := h.client.Exists(ctx, fmt.Sprintf("linkme:user:ssid:%s", ssid)).Result()
	if err != nil {
		return err
	}
	if c != 0 {
		return errors.New("token失效")
	}
	return nil
}

// ClearToken 清空token
func (h *handler) ClearToken(ctx context.Context) error {
	uc, ok := ctx.Value("user").(UserClaims)
	if !ok {
		return errors.New("failed to get user claims from context")
	}
	// 获取 refresh token
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return errors.New("failed to get transport from context")
	}
	refreshTokenString := tr.RequestHeader().Get("X-Refresh-Token")
	if refreshTokenString == "" {
		return errors.New("missing refresh token")
	}
	// 解析 refresh token
	refreshClaims := &RefreshClaims{}
	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, refreshClaims, func(token *jwt.Token) (interface{}, error) {
		return Key2, nil
	})
	if err != nil || !refreshToken.Valid {
		return errors.New("invalid refresh token")
	}
	// 设置redis中的会话ID键的过期时间为refresh token的剩余过期时间
	remainingTime := refreshClaims.ExpiresAt.Time.Sub(time.Now())
	if er := h.client.Set(ctx, fmt.Sprintf("linkme:user:ssid:%s", uc.Ssid), "", remainingTime).Err(); er != nil {
		return er
	}
	return nil
}
