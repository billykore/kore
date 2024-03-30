package websocket

import (
	"encoding/json"

	"github.com/billykore/kore/backend/pkg/log"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	Id   string
	Conn *websocket.Conn
	Pool *Pool
}

func NewClient(conn *websocket.Conn, pool *Pool) *Client {
	client := &Client{
		Conn: conn,
		Pool: pool,
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return client
	}
	client.Id = id.String()
	return client
}

type Message struct {
	Type int  `json:"type"`
	Body Body `json:"body"`
}

type Body struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	logger := log.NewLogger()
	for {
		msgType, p, err := c.Conn.ReadMessage()
		if err != nil {
			logger.Error(err)
			return
		}

		msgBody := Body{}
		if err := json.Unmarshal(p, &msgBody); err != nil {
			logger.Error(err)
			return
		}

		message := Message{Type: msgType, Body: msgBody}
		c.Pool.Broadcast <- message

		logger.Infof("Message Received: %+v", message.Body)
	}
}
