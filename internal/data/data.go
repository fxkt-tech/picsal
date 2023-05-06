package data

import (
	"github.com/fxkt-tech/picsal/internal/conf"
	"github.com/fxkt-tech/picsal/pkg/image"
	"github.com/go-kratos/kratos/v2/log"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewData,

	NewImageRepo,
)

type Data struct {
	// debug bool
	imgStore *image.Storage
}

func NewData(c *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	// debug
	// debug := os.Getenv("YILAN_LOG_LEVEL") == "DEBUG"

	return &Data{
			// debug: debug,
			imgStore: image.NewStorage(),
		}, func() {
			log.NewHelper(logger).Warn("closing the data resources")
		}, nil
}