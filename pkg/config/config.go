package config

import (
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

func Load[T any](filePath string) (T, error) {
	c := config.New(
		config.WithSource(
			file.NewSource(filePath),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc T
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	return bc, nil
}

func LoadConfigFileFromEnv() (string, error) {
	env := os.Getenv("YILAN_ENV")
	switch env {
	case "prod", "test", "dev", "local":
		return fmt.Sprintf("configs/%s.yaml", env), nil
	default:
		return "", fmt.Errorf("unkown env: %s", env)
	}
}
