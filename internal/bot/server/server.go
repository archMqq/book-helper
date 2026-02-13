package bot

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/archMqq/book-helper/internal/logger"
	"github.com/archMqq/book-helper/internal/models"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
	"gopkg.in/telebot.v4"
	tele "gopkg.in/telebot.v4"
)

var (
	ErrInternalServer = errors.New("Наблюдается ошибка на стороне сервера. Повторите попытку позже.")
	ErrServerIsBusy   = errors.New("В настоящий момент сервер слишком занят. Повторите попытку позже.")
)

type State int

const (
	StateNone State = iota
	StateRegistered
	StateWaitAuthors
	StateWaitGenres
	StateSearching
)

type UserService interface {
	// Возвращает ErrUserExists, если пользователь существует
	CreateUser(context.Context, int64, string) error

	GetPreferences(context.Context, int64) (*models.Preferences, error)

	SaveAuthors(context.Context, int64, []string) error
	SaveGenres(context.Context, int64, []string) error
}

type server struct {
	bot         *tele.Bot
	logger      *logrus.Entry
	userService UserService
	states      *userStates
	msgQueue    *msgQueue
}

func newServer(b *tele.Bot, userService UserService) *server {
	return &server{
		logger:      logger.InitForService("bot"),
		bot:         b,
		userService: userService,
		states:      newUserStates(),
		msgQueue:    newMsgQueue(),
	}
}

type userStates struct {
	states map[int64]State
	mu     *sync.RWMutex
}

func newUserStates() *userStates {
	return &userStates{
		states: make(map[int64]State),
		mu:     &sync.RWMutex{},
	}
}

func (us *userStates) Read(userID int64) (State, bool) {
	us.mu.RLock()
	defer us.mu.RUnlock()

	res, ok := us.states[userID]
	return res, ok
}

func (us *userStates) Save(userID int64, state State) {
	us.mu.Lock()
	defer us.mu.Unlock()

	us.states[userID] = state
}

type msgQueue struct {
	globalLimiter *rate.Limiter
	userLimiter   map[int64]*rate.Limiter
	mu            *sync.RWMutex
}

func newMsgQueue() *msgQueue {
	queue := &msgQueue{
		globalLimiter: rate.NewLimiter(20, 20),
		userLimiter:   make(map[int64]*rate.Limiter),
		mu:            &sync.RWMutex{},
	}

	go queue.cleanLoop()
	return queue
}

func (mq *msgQueue) getUserLimiter(userID int64) *rate.Limiter {
	mq.mu.RLock()
	if lim, ok := mq.userLimiter[userID]; ok {
		mq.mu.RUnlock()
		return lim
	}
	mq.mu.RUnlock()

	mq.mu.Lock()
	defer mq.mu.Unlock()

	if lim, ok := mq.userLimiter[userID]; ok {
		return lim
	}

	lim := rate.NewLimiter(1, 1)
	mq.userLimiter[userID] = lim
	return lim
}

func (mq *msgQueue) Middleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(ctx tele.Context) error {
		userID := ctx.Sender().ID
		botCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		userLim := mq.getUserLimiter(userID)
		userLim.Wait(botCtx)

		mq.globalLimiter.Wait(botCtx)

		return next(ctx)
	}
}

func (mq *msgQueue) cleanLoop() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		mq.mu.Lock()

		for userID, lim := range mq.userLimiter {
			if lim.Allow() {
				delete(mq.userLimiter, userID)
			}
		}
		mq.mu.Unlock()
	}
}
