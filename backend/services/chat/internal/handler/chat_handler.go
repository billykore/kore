package handler

import (
	"github.com/billykore/kore/backend/pkg/net/websocket"
	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	pool *websocket.Pool
}

func NewChatHandler(pool *websocket.Pool) *ChatHandler {
	return &ChatHandler{
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
