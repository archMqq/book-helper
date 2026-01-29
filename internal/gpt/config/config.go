package config

import (
	"fmt"
	"log"

	"github.com/archMqq/book-helper/internal/basecfg"
	"github.com/spf13/viper"
)

type Config struct {
	basecfg.BaseCfg
	GPTToken    string  `mapstructure:"GPT_TOKEN"`
	Model       string  `mapstructure:"model"`
	Temperature float32 `mapstructure:"temperature"`
	MaxTokens   int     `mapstructure:"max_tokens"`
	Prompt      string  `mapstructure:"prompt"`
}

func NewConfig() *Config {
	cfg, err := load()
	if err != nil {
		log.Fatalf("err cfg loading: %s", err)
	}

	return cfg
}

func load() (*Config, error) {
	c := &Config{
		BaseCfg: basecfg.New(),
	}

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("env readin: %w", err)
	}

	viper.SetConfigFile("configs/config.toml")
	viper.SetConfigType("toml")
	if err := viper.MergeInConfig(); err != nil {
		return nil, fmt.Errorf("toml mergein: %w", err)
	}

	if err := viper.Unmarshal(c); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	return c, nil
}
