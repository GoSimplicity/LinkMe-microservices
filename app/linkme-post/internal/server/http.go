package server

import (
	v1 "github.com/GoSimplicity/LinkMe-microservices/api/post/v1"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/internal/conf"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/internal/service"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, sp *service.PostService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
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
	v1.RegisterPostHTTPServer(srv, sp)
	return srv
}
