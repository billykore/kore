package messaging

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/billykore/kore/backend/infra/messaging/consumer"
	"github.com/billykore/kore/backend/pkg/config"
	"github.com/billykore/kore/backend/pkg/logger"
)

type Consumer struct {
	cfg           *config.Config
	log           *logger.Logger
	orderConsumer *consumer.OrderConsumer
}

func NewConsumer(cfg *config.Config, log *logger.Logger, orderConsumer *consumer.OrderConsumer) *Consumer {
	return &Consumer{
		cfg:           cfg,
		log:           log,
		orderConsumer: orderConsumer,
	}
}

func (c *Consumer) Consume() {
	c.log.Usecase("Consume").Info("Consumer start consume...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := c.orderConsumer.ListenOrderStatusChanges(ctx)
	if err != nil {
		c.log.Fatalf("Failed to start user consumer: %v", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop // Wait for a termination signal
	c.log.Info("Shutting down consumer...")

	time.Sleep(2 * time.Second)
	cancel()

	c.log.Info("Consumer stopped")
	os.Exit(0)
}
