package bootstrap

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	kzap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Bootstrap[C any] struct {
	logger log.Logger
	env    *Env
	flag   *Flag
	config C

	err    error
	defers []func()
}

func New[C any]() *Bootstrap[C] {
	return &Bootstrap[C]{}
}

func Default[C any]() *Bootstrap[C] {
	rand.Seed(time.Now().Unix())
	json.MarshalOptions.UseProtoNames = true     // 使用jsontag作为jsonkey
	json.MarshalOptions.EmitUnpopulated = true   // jsonkey为空时仍然渲染
	json.UnmarshalOptions.DiscardUnknown = false // 如果传入为定义字段

	return New[C]()
}

func (b *Bootstrap[C]) Logger() log.Logger {
	return b.logger
}

func (b *Bootstrap[C]) Env() *Env {
	return b.env
}

func (b *Bootstrap[C]) Flag() *Flag {
	return b.flag
}

func (b *Bootstrap[C]) Config() C {
	return b.config
}

// ---------

func (b *Bootstrap[C]) LoadEnv() *Bootstrap[C] {
	if b.err != nil {
		return b
	}

	b.env = NewEnv().Load()
	return b
}

func (b *Bootstrap[C]) LoadFlag() *Bootstrap[C] {
	if b.err != nil {
		return b
	}

	b.flag = NewFlag().Load()
	return b
}

func (b *Bootstrap[C]) LoadLogger() *Bootstrap[C] {
	if b.err != nil {
		return b
	}

	b.logger = b.useZapLogger()
	log.SetLogger(b.logger)
	return b
}

func logLevelTransform(ll string) zapcore.Level {
	level, err := zapcore.ParseLevel(ll)
	if err != nil {
		// default level: info
		return zapcore.InfoLevel
	}
	return level
}

func (b *Bootstrap[C]) useZapLogger() log.Logger {
	var logLevel string
	if b.env != nil {
		logLevel = b.env.YilanLogLevel
	}

	return log.With(kzap.NewLogger(
		zap.New(
			zapcore.NewCore(
				zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
					TimeKey:       "ts",
					LevelKey:      "level",
					StacktraceKey: "stack",
					EncodeTime:    zapcore.ISO8601TimeEncoder,
					LineEnding:    zapcore.DefaultLineEnding,
					EncodeLevel:   zapcore.CapitalColorLevelEncoder,
				}),
				zapcore.AddSync(os.Stdout),
				logLevelTransform(logLevel),
			),
			// zap.Development(),
		)),
	)
}

func (b *Bootstrap[C]) LoadConfig() *Bootstrap[C] {
	if b.err != nil {
		return b
	}

	b.config = b.useLocalFileConfig()
	return b
}

func configFilePathFromEnv(env string) string {
	return fmt.Sprintf("configs/%s.yaml", env)
}

func (b *Bootstrap[C]) useLocalFileConfig() (cfg C) {
	var configFilePath string

	// 优先使用flag中的
	if b.flag != nil {
		if b.flag.configFilePath == "" {
			if b.env != nil {
				configFilePath = configFilePathFromEnv(b.env.YilanEnv)
			} else {
				b.err = errors.New("配置文件路径未配置")
				return
			}
		} else {
			configFilePath = b.flag.configFilePath
		}
	}

	c := config.New(
		config.WithSource(
			file.NewSource(configFilePath),
		),
	)

	if err := c.Load(); err != nil {
		b.err = err
		return
	}

	if err := c.Scan(&cfg); err != nil {
		b.err = err
		return
	}

	b.defers = append(b.defers, func() { c.Close() })

	return
}

func (b *Bootstrap[C]) Result() (*Bootstrap[C], error) {
	if b.err != nil {
		return nil, b.err
	}

	return b, nil
}

func (b *Bootstrap[C]) Defer() {
	for _, df := range b.defers {
		df()
	}
}
