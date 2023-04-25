package server

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"

	v1 "github.com/fxkt-tech/picsal/api/v1"
	"github.com/fxkt-tech/picsal/internal/conf"
	"github.com/fxkt-tech/picsal/internal/data"
	"github.com/fxkt-tech/picsal/internal/service"
)

func NewHTTPServer(c *conf.Bootstrap, d *data.Data, imgSrv *service.ImageService, logger log.Logger) (*http.Server, error) {
	opts := []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
		http.Filter(
			handlers.CORS(
				handlers.AllowedOrigins([]string{"*"}),
				handlers.AllowedHeaders([]string{"Content-Type", "AccessToken", "X-CSRF-Token", "Authorization", "Token", "X-Token", "X-User-Id"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE"}),
				handlers.ExposedHeaders([]string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"}),
				handlers.AllowCredentials(),
			),
		),
	}
	if c.HttpServer.Network != "" {
		opts = append(opts, http.Network(c.HttpServer.Network))
	}
	if c.HttpServer.Addr != "" {
		opts = append(opts, http.Address(c.HttpServer.Addr))
	}
	if c.HttpServer.Timeout > 0 {
		opts = append(opts, http.Timeout(time.Duration(c.HttpServer.Timeout)*time.Second))
	}
	srv := http.NewServer(opts...)
	v1.RegisterImageServiceHTTPServer(srv, imgSrv)
	return srv, nil
}
