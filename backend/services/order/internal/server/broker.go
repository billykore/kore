package server

import (
	"github.com/billykore/kore/backend/pkg/log"
	"github.com/billykore/kore/backend/services/order/internal/handler"
)

type BrokerServer struct {
	log          *log.Logger
	orderHandler *handler.OrderHandler
}

func NewBrokerServer(log *log.Logger, orderHandler *handler.OrderHandler) *BrokerServer {
	return &BrokerServer{
		log:          log,
		orderHandler: orderHandler,
	}
}

func (s *BrokerServer) Serve() {
	s.orderHandler.ListenOrderStatusChanges()
}
