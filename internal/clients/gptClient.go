package clients

import (
	"context"
	"fmt"
	"log"

	"github.com/archMqq/book-helper/internal/config"
	"github.com/sheeiavellie/go-yandexgpt"
)

type YandexGPTClient struct {
	token       string
	model       string
	temperature float64
	maxTokens   int
	prompt      string

	client *yandexgpt.YandexGPTClient
}

func NewYandexGpt(cfg config.GPTData) *YandexGPTClient {
	return &YandexGPTClient{
		model:       cfg.Model,
		temperature: cfg.Temperature,
		maxTokens:   cfg.MaxTokens,
		prompt:      cfg.Prompt,
		client:      yandexgpt.New(yandexgpt.CfgApiKey(cfg.GPTToken)),
	}
}

func (g *YandexGPTClient) AskForNewBooks(ctx context.Context, pref string) (string, error) {
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

	if len(res.Result.Alternatives) == 0 {
		return "Пустой ответ", nil
	}

	return res.Result.Alternatives[0].Message.Text, nil
}
