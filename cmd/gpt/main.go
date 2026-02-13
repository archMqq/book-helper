package main

import (
	"github.com/archMqq/book-helper/internal/gpt/clients"
	"github.com/archMqq/book-helper/internal/gpt/config"
	"github.com/archMqq/book-helper/internal/gpt/server"
	"github.com/archMqq/book-helper/internal/gpt/services"
)

func main() {
	config := config.NewConfig()

	gptClient := clients.NewYandexGpt(*config)
	recService := services.New(gptClient)
	server := server.New(recService)

	server.Start(config)
}
