package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/archMqq/book-helper/internal/bot"
	"github.com/archMqq/book-helper/internal/config"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/book-helper.toml", "path to app config")
}

func main() {
	flag.Parse()

	config := config.NewConfig()
	_, err := toml.Decode(configPath, config)
	if err != nil {
		log.Fatalf("%s %s", "err decoding toml:", err)
	}

	bot.Start(config)
}
