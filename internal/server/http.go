package server

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"

	v1 "github.com/fxkt-tech/picsal/api/image/v1"
	"github.com/fxkt-tech/picsal/internal/conf"
	"github.com/fxkt-tech/picsal/internal/data"
	"github.com/fxkt-tech/picsal/internal/service"
)

func NewHTTPServer(c *conf.Bootstrap, d *data.Data, imgSrv *service.ImageService, logger log.Logger) (*http.Server, error) {
	opts := []http.ServerOption{
		http.ResponseEncoder(ResponseEncoder),
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
	v1.RegisterProcessingServiceHTTPServer(srv, imgSrv)
	return srv, nil
}

func ContentType(subtype string) string {
	return strings.Join([]string{"application", subtype}, "/")
}

var _ GetImageStreamer = &v1.ImageStreamResponse{}

type GetImageStreamer interface {
	GetImageStream() []byte
	GetContentType() string
}

func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if v == nil {
		return nil
	}

	var (
		data        []byte
		contentType string
	)

	// 支持直接返回图片和返回接口数据
	if is, ok := v.(GetImageStreamer); ok {
		fmt.Println("is steam")
		data = is.GetImageStream()
		contentType = is.GetContentType()
	} else {
		fmt.Println("is jsonres")
		codec, _ := http.CodecForRequest(r, "Accept")
		data, _ = codec.Marshal(v)
		contentType = ContentType(codec.Name())
	}
	w.Header().Set("Content-Type", contentType)

	_, err := w.Write(data)
	if err != nil {
		return err
	}
	return nil
}
