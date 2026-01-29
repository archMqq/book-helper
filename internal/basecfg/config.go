package basecfg

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type BaseCfg struct {
	KafkaURL string `mapstructure:"KAFKA_URL"`
	RedisURL string `mapstructure:"REDIS_URL"`
}

func New() BaseCfg {
	cfg, err := load()
	if err != nil {
		log.Fatalf("err basecfg loading: %s", err)
	}

	return cfg
}

func load() (BaseCfg, error) {
	c := BaseCfg{}

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.ReadInConfig()

	if err := viper.Unmarshal(&c); err != nil {
		return c, fmt.Errorf("unmarshal: %w", err)
	}

	return c, nil
}
