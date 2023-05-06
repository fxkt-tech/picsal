package bootstrap

import "flag"

type Flag struct {
	configFilePath string
}

func NewFlag() *Flag {
	return &Flag{}
}

func (f *Flag) Load() *Flag {
	flag.StringVar(&f.configFilePath, "conf", "", "config file path, eg: -conf configs/config.yaml")
	flag.Parse()
	return f
}
