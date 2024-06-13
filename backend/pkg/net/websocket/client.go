package websocket

import (
	"encoding/json"

	"github.com/billykore/kore/backend/pkg/log"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Client to communicate with the websocket.
type Client struct {
	Id   string
	Conn *websocket.Conn
	Pool *Pool
}

// NewClient create new client connection to websocket.
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
	Name    string `json:"name"`
	Message string `json:"message"`
}

// Read the message from websocket.
func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	logger := log.NewLogger()
	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			logger.Error(err)
			return
		}
		message := Message{}
		err = json.Unmarshal(p, &message)
		if err != nil {
			logger.Error(err)
		}
		c.Pool.Broadcast <- message
		logger.Infof("Message Received: %+v", message)
	}
}
