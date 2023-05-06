package bootstrap

import (
	"os"
)

type Env struct {
	YilanLogLevel string
	YilanConfig   string
	YilanEnv      string
}

func NewEnv() *Env {
	return &Env{}
}

func (e *Env) Load() *Env {
	e.YilanLogLevel = os.Getenv("YILAN_LOG_LEVEL")
	e.YilanConfig = os.Getenv("YILAN_CONFIG")
	e.YilanEnv = os.Getenv("YILAN_ENV")

	return e
}
