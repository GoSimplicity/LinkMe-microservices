package server

import (
	v1 "github.com/GoSimplicity/LinkMe-microservices/api/user/v1"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-user/internal/conf"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-user/internal/middleware"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-user/internal/service"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, userService *service.UserService, jwtMiddleware *middleware.JWTMiddleware) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			jwtMiddleware.CheckLogin(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterUserHTTPServer(srv, userService)
	return srv
}
