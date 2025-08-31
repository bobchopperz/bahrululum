package config

import "time"

type ServerConfig struct {
	Port    string        `mapstructure:"port"`
	Host    string        `mapstructure:"host"`
	Mode    string        `mapstructure:"mode"`
	Timeout time.Duration `mapstructure:"timeout"`
}
