package main

import (
	"fmt"

	"github.com/archMqq/book-helper/internal/bot/config"
	bot "github.com/archMqq/book-helper/internal/bot/server"
)

func main() {
	config := config.NewConfig()
	fmt.Println()
	bot.Start(&config)
}
