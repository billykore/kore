package handler

import (
	"github.com/billykore/kore/backend/pkg/net/websocket"
	"github.com/billykore/kore/backend/services/chat/internal/usecase"
	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	uc   *usecase.ChatUsecase
	pool *websocket.Pool
}

func NewChatHandler(uc *usecase.ChatUsecase, pool *websocket.Pool) *ChatHandler {
	return &ChatHandler{
		uc:   uc,
		pool: pool,
	}
}

func (s *ChatHandler) Chat(ctx echo.Context) error {
	conn, err := websocket.Upgrade(ctx.Response(), ctx.Request())
	if err != nil {
		return err
	}

	client := websocket.NewClient(conn, s.pool)
	s.pool.Register <- client
	client.Read()

	return nil
}
