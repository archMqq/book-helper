package clients

import "github.com/archMqq/book-helper/internal/config"

type GPTClient struct {
}

func NewGpt(cfg config.GPTData) *GPTClient {
	return &GPTClient{}
}
