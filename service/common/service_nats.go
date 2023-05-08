package common

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
)

func natsUri() string {
	natsUri := os.Getenv("NATS_URI")
	natsPort := os.Getenv("NATS_PORT")
	if natsUri == "" || natsPort == "" {
		return "nats://localhost:4222"
	}
	return fmt.Sprintf("nats://%s:%s", natsUri, natsPort)
}

// Receive messages from nats-streaming;
// Validate data and save valid orders
func (s *Service) acceptOrders(msg *stan.Msg) {
	var order Order
	err := json.Unmarshal(msg.Data, &order)
	if err != nil {
		log.Printf("Could not parse data: %v", err)
		return
	}
	err = validator.New().Struct(&order)
	if err != nil {
		log.Printf("Invalid data: %v", err)
		return
	}
	s.db.Create(&order)
	log.Printf("Received order with id: %s\n", order.OrderUID)
	s.cache[order.OrderUID] = &order
}
