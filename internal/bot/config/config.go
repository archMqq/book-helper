package config

import (
	"fmt"
	"log"

	"github.com/archMqq/book-helper/internal/basecfg"
	"github.com/spf13/viper"
)

type Config struct {
	basecfg.BaseCfg
	Database string `mapstructure:"DB_URL"`
	TGToken  string `mapstructure:"TG_TOKEN"`
}

func NewConfig() Config {
	cfg, err := load()
	if err != nil {
		log.Fatalf("err cfg loading: %s", err)
	}

	return *cfg
}

func load() (*Config, error) {
	c := &Config{
		BaseCfg: basecfg.New(),
	}

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.ReadInConfig()

	if err := viper.Unmarshal(c); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	return c, nil
}
