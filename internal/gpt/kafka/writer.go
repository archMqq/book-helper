package kafka

import (
	"context"
	"encoding/json"

	"github.com/archMqq/book-helper/internal/models"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func StartWrtiting(ctx context.Context, broker, topic string,
	logger *logrus.Entry, in chan *models.GptServiceOut) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{broker},
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    10,
		RequiredAcks: 1,
	})
	defer writer.Close()

	for {
		select {
		case <-ctx.Done():
			return
		case <-in:
			data := <-in
			payload, err := json.Marshal(data)
			if err != nil {
				logger.Error("kafka payload marshalling error", "error", err)
			}

			writer.WriteMessages(ctx, kafka.Message{
				Value: payload,
			})
		}
	}
}
