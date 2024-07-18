package middleware

import (
	"context"
	"errors"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-user/internal/data"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/wire"
	"log"
)

var ProviderSet = wire.NewSet(NewJWTMiddleware)

var (
	ErrMissingToken = errors.New("missing token")
)

type JWTMiddleware struct {
	data.Handler
}

func NewJWTMiddleware(hdl data.Handler) *JWTMiddleware {
	return &JWTMiddleware{
		Handler: hdl,
	}
}

// CheckLogin 校验JWT
func (m *JWTMiddleware) CheckLogin() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			// 从 Kratos 上下文中获取传输信息
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, ErrMissingToken
			}
			// 打印所有请求头
			headers := tr.RequestHeader()
			res := headers.Get("test")
			log.Println(res)
			// 获取请求路径
			path := tr.Operation()
			if path == "/api.user.v1.User/SignUp" ||
				path == "/api.user.v1.User/Login" ||
				path == "/api.user.v1.User/RefreshToken" ||
				path == "/api.user.v1.User/ChangePassword" {
				return handler(ctx, req)
			}
			// 从请求头中提取 JWT Token
			tokenStr := m.ExtractToken(ctx)
			if tokenStr == "" {
				return nil, ErrMissingToken
			}
			var uc data.UserClaims
			token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
				return data.Secret, nil
			})
			if err != nil || token == nil || !token.Valid {
				return nil, errors.New("invalid or expired token")
			}
			if uc.UserAgent == "" {
				return nil, errors.New("missing user agent")
			}
			// 检查会话
			err = m.CheckSession(ctx, uc.Ssid)
			if err != nil {
				return nil, err
			}
			// 将用户信息存储到上下文中
			ctx = context.WithValue(ctx, "user", uc)
			return handler(ctx, req)
		}
	}
}
