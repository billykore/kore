package websocket

import (
	"github.com/billykore/kore/backend/pkg/logger"
)

// Pool hold the websocket connection pool.
type Pool struct {
	log        *logger.Logger
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

// NewPool return instance of Pool.
func NewPool() *Pool {
	return &Pool{
		log:        logger.New(),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

// Start new connection to the Pool.
func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.log.Infof("new connection added (%s)", client.Id)
			pool.Clients[client] = true
			break
		case client := <-pool.Unregister:
			pool.log.Infof("new connection deleted (%s)", client.Id)
			delete(pool.Clients, client)
			break
		case message := <-pool.Broadcast:
			for c := range pool.Clients {
				pool.log.Infof("send message to client %s", c.Id)
				if err := c.Conn.WriteJSON(message); err != nil {
					pool.log.Errorf("failed to send message to client %s: %v", c.Id, err)
					return
				}
			}
		}
	}
}
