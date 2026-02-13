package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/archMqq/book-helper/internal/gpt/config"
	"github.com/archMqq/book-helper/internal/gpt/kafka"
	"github.com/archMqq/book-helper/internal/logger"
	"github.com/archMqq/book-helper/internal/models"
	"github.com/sirupsen/logrus"
)

type RecService interface {
	GetBooks(context.Context, *models.Preferences) (string, error)
}

type server struct {
	logger     *logrus.Entry
	recService RecService
}

func New(service RecService) *server {
	return &server{
		logger:     logger.InitForService("gpt"),
		recService: service,
	}
}

func (s *server) Start(cfg *config.Config) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		s.logger.Info("server stop by syscall")
		cancel()
	}()

	in := make(chan *models.KakfaGPTAskQuery, 1)
	out := make(chan *models.GptServiceOut, 1)
	go kafka.StartReading(ctx, cfg.KafkaURL, cfg.TopicIn, s.logger, in)
	go kafka.StartWrtiting(ctx, cfg.KafkaURL, cfg.TopicOut, s.logger, out)

	for {
		select {
		case <-ctx.Done():
			return
		case <-in:
			res := <-in
			data, err := s.Handle(ctx, res)
			if err != nil {
				s.logger.Error(err)
				continue
			}

			out <- data
		}
	}
}

func (s server) Handle(ctx context.Context, query *models.KakfaGPTAskQuery) (out *models.GptServiceOut, err error) {
	var res string
	out = &models.GptServiceOut{}

	switch query.Cmd {
	case "recommend":
		res, err = s.recService.GetBooks(ctx, &query.Pref)
		if err != nil {
			return
		}
	default:
		return nil, fmt.Errorf("unknown command")
	}

	s.wrapRes(res, out)

	return out, nil
}

func (s server) wrapRes(input string, out *models.GptServiceOut) {
	lines := strings.Split(input, "\n")

	var books []models.BookData
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		author := strings.TrimSpace(parts[0])
		book := strings.TrimSpace(parts[1])

		books = append(books, models.BookData{
			AuthorName: author,
			BookName:   book,
		})
	}

	out.Books = books
}
