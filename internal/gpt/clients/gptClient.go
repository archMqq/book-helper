package clients

import (
	"context"
	"fmt"
	"log"

	"github.com/archMqq/book-helper/internal/gpt/config"
	"github.com/sheeiavellie/go-yandexgpt"
)

type YandexGPTClient struct {
	token       string
	model       string
	temperature float32
	maxTokens   int

	client *yandexgpt.YandexGPTClient
}

func NewYandexGpt(cfg config.Config) *YandexGPTClient {
	return &YandexGPTClient{
		model:       cfg.Model,
		temperature: cfg.Temperature,
		maxTokens:   cfg.MaxTokens,
		client:      yandexgpt.New(yandexgpt.CfgApiKey(cfg.GPTToken)),
	}
}

func (g *YandexGPTClient) AskForNewBooks(ctx context.Context, pref string) (string, error) {
	prompt := `You are a knowledgeable librarian with access to a vast collection of books. 
	Your task is to recommend 10 books based on the user's interests. 
	For each recommendation, you will provide only the title of the book and the author's name,
	without any additional commentary or explanations. 
	Make sure your recommendations are diverse and reflect different aspects of the user's interests,
	and can include new authors and subgenres. 
	Format your response as a simple list with each entry on a new line. 
	You should reply as russian with a message structure as author_full_name:book_name new_line.`

	res, err := g.client.GetCompletion(ctx, yandexgpt.YandexGPTRequest{
		ModelURI: g.model,
		CompletionOptions: yandexgpt.YandexGPTCompletionOptions{
			Temperature: float32(g.temperature),
			MaxTokens:   g.maxTokens,
		},
		Messages: []yandexgpt.YandexGPTMessage{
			{
				Role: yandexgpt.YandexGPTMessageRoleSystem,
				Text: prompt,
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
