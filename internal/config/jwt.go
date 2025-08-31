package config

import "time"

type JWTConfig struct {
	Secret     string        `mapstructure:"secret"`
	Expiry     time.Duration `mapstructure:"expiry"`
	RefreshExp time.Duration `mapstructure:"refresh_expiry"`
}
