package clients

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/archMqq/book-helper/internal/config"
	"github.com/sheeiavellie/go-yandexgpt"
)

type GPTClient struct {
	token       string
	model       string
	temperature float64
	maxTokens   int
	prompt      string

	client *yandexgpt.YandexGPTClient
}

func NewGpt(cfg config.GPTData) *GPTClient {
	return &GPTClient{
		model:       cfg.Model,
		temperature: cfg.Temperature,
		maxTokens:   cfg.MaxTokens,
		prompt:      cfg.Prompt,
		client:      yandexgpt.New(yandexgpt.CfgApiKey(cfg.GPTToken)),
	}
}

func (g *GPTClient) AskForNewBooks(pref string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := g.client.GetCompletion(ctx, yandexgpt.YandexGPTRequest{
		ModelURI: g.model,
		CompletionOptions: yandexgpt.YandexGPTCompletionOptions{
			Temperature: float32(g.temperature),
			MaxTokens:   g.maxTokens,
		},
		Messages: []yandexgpt.YandexGPTMessage{
			{
				Role: yandexgpt.YandexGPTMessageRoleSystem,
				Text: g.prompt,
			},
			{
				Role: yandexgpt.YandexGPTMessageRoleUser,
				Text: fmt.Sprintf("My preferences: %s", pref),
			},
		},
	})

	if err != nil {
		log.Fatal(err)
		return "", fmt.Errorf("YandexGPT: %w", err)
	}

	// ← res это *YandexGPTResponse, не строка!
	if len(res.Result.Alternatives) == 0 {
		return "Пустой ответ", nil
	}

	return res.Result.Alternatives[0].Message.Text, nil
}
