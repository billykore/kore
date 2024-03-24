package service

import (
	"fmt"
	"time"

	"github.com/billykore/kore/backend/services/chat/usecase"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type ChatService struct {
	uc *usecase.ChatUsecase
}

func NewChatService(uc *usecase.ChatUsecase) *ChatService {
	return &ChatService{uc: uc}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 10 * time.Second,
}

func (s *ChatService) Greet(ctx echo.Context) error {
	ws, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	for {
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			ctx.Logger().Error(err)
		}

		_, msg, err := ws.ReadMessage()
		if err != nil {
			ctx.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}
