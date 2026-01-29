package main

import (
	"fmt"

	"github.com/archMqq/book-helper/internal/gpt/config"
)

func main() {
	config := config.NewConfig()
	fmt.Println(config)
}
