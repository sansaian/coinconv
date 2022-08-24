package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Loglevel   string
	Interest   float64
	CoinMarket *CoinMarket
}

type CoinMarket struct {
	Url     string
	Token   string
	Timeout time.Duration
}

func NewFromENV() *Config {
	viper.AutomaticEnv()
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("TIMEOUT", "5s")
	return &Config{
		Loglevel: viper.GetString("LOG_LEVEL"),
		Interest: viper.GetFloat64("INTEREST"),
		CoinMarket: &CoinMarket{
			Url:     viper.GetString("URL"),
			Token:   viper.GetString("TOKEN"),
			Timeout: viper.GetDuration("TIMEOUT"),
		},
	}
}

func (cfg *Config) IsValid() bool {
	if cfg.CoinMarket == nil || cfg.CoinMarket.Url == "" ||
		cfg.CoinMarket.Token == "" {
		return false
	}
	return true
}
