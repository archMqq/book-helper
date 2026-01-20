package recommend

import (
	"github.com/archMqq/book-helper/internal/clients"
	"github.com/archMqq/book-helper/internal/config"
)

type RecService struct {
	gptClient *clients.GPTClient
}

func New(cfg *config.RecData) *RecService {
	return &RecService{
		gptClient: clients.NewGpt(cfg.GPTData),
	}
}

func (rs RecService) Request() {}
