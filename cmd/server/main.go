package main

import (
	"flag"
	"math/rand"
	"os"
	"time"

	"github.com/fxkt-tech/picsal/internal/conf"
	"github.com/fxkt-tech/picsal/pkg/config"

	"github.com/go-kratos/kratos/v2"
	kconfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	rand.Seed(time.Now().Unix())
	json.MarshalOptions.UseProtoNames = true
	json.MarshalOptions.EmitUnpopulated = true
	json.UnmarshalOptions.DiscardUnknown = false
	flag.StringVar(&flagconf, "conf", "", "config path, eg: -conf conf/dev.yaml")
}

func newApp(logger log.Logger, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(hs),
	)
}

func main() {
	var err error
	flag.Parse()

	if flagconf == "" {
		flagconf, err = config.LoadConfigFileFromEnv()
		if err != nil {
			panic(err)
		}
	}
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
	)

	c := kconfig.New(
		kconfig.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	// bc, err := config.Load[conf.Bootstrap](flagconf)
	// if err != nil {
	// 	panic(err)
	// }

	app, cleanup, err := initApp(&bc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
