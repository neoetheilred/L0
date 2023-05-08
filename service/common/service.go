package common

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nats-io/stan.go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	db    *gorm.DB
	sc    *stan.Conn
	mx    *mux.Router
	cache map[string]*Order
}

func NewService(cdn string) (*Service, error) {
	db, err := gorm.Open(postgres.Open(cdn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Order{}, &Delivery{}, &Payment{}, &Item{})
	sc, err := stan.Connect("test-cluster", "accepter", stan.NatsURL(natsUri()))
	if err != nil {
		return nil, err
	}
	mx := mux.NewRouter()
	service := &Service{db: db, sc: &sc, mx: mx, cache: map[string]*Order{}}
	service.configureRoutes()
	return service, nil
}

// Run service
func (s *Service) Run() {
	log.Println("Starting up...")
	s.restoreCache()
	(*s.sc).Subscribe("orders", s.acceptOrders)

	log.Println("Service started")
	http.ListenAndServe(":8000", handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "OPTIONS"}),
		handlers.AllowCredentials(),
		handlers.AllowedHeaders([]string{"Content-Type"}))(s.mx))
}

func (s *Service) Stop() {
	log.Println("Shutting down...")
	(*s.sc).Close()
}
