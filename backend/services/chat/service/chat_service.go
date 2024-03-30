package service

import (
	"github.com/billykore/kore/backend/pkg/websocket"
	"github.com/billykore/kore/backend/services/chat/usecase"
	"github.com/labstack/echo/v4"
)

type ChatService struct {
	uc   *usecase.ChatUsecase
	pool *websocket.Pool
}

func NewChatService(uc *usecase.ChatUsecase, pool *websocket.Pool) *ChatService {
	return &ChatService{
		uc:   uc,
		pool: pool,
	}
}

func (s *ChatService) Chat(ctx echo.Context) error {
	conn, err := websocket.Upgrade(ctx.Response(), ctx.Request())
	if err != nil {
		return err
	}

	client := websocket.NewClient(conn, s.pool)
	s.pool.Register <- client
	client.Read()

	return nil
}
