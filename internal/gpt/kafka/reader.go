package kafka

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/archMqq/book-helper/internal/models"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func StartReading(ctx context.Context, broker, topic string, logger *logrus.Entry, out chan *models.KakfaGPTAskQuery) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
	defer reader.Close()

	for {
		select {
		case <-ctx.Done():
			return

		default:
			msg, err := reader.ReadMessage(ctx)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return
				}
				logger.Warnf("Got error while reading msg: %w", err)
				continue
			}

			query := models.KakfaGPTAskQuery{}
			if err = json.Unmarshal(msg.Value, &query); err != nil {
				logger.Warnf("Got error msg: %w", err)
				continue
			}

			out <- &query
		}
	}
}
