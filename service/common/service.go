package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nats-io/stan.go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	natsUri := os.Getenv("NATS_URI")
	natsPort := os.Getenv("NATS_PORT")
	sc, err := stan.Connect("test-cluster", "accepter", stan.NatsURL(fmt.Sprintf("nats://%s:%s", natsUri, natsPort)))
	if err != nil {
		return nil, err
	}
	mx := mux.NewRouter()
	service := &Service{db: db, sc: &sc, mx: mx, cache: map[string]*Order{}}
	service.configureRoutes()
	return service, nil
}

func (s *Service) configureRoutes() {
	s.mx.HandleFunc("/order/{id}", s.HandleFindOrderById)

	s.mx.HandleFunc("/cached", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(s.cache)
	})

	s.mx.HandleFunc("/available", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(keys(s.cache))
	})
}

// Gives back Order with corresponding order_uid (`id` in request url)
func (s *Service) HandleFindOrderById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.Println(id)
	if order, ok := s.cache[id]; ok {
		log.Println("cache")
		log.Println(order)
		json.NewEncoder(w).Encode(order)
	} else {
		var order Order
		r := s.db.Preload(clause.Associations).First(&order, "order_uid = ?", id)
		if r.Error != nil {
			log.Printf("Not found order with id %s\n", id)
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			log.Println(order)
			json.NewEncoder(w).Encode(order)
		}
	}
	w.Header().Add("Content-Type", "application/json")
}

// Get keys of a map
func keys[T comparable, U any](m map[T]U) []T {
	k := make([]T, len(m))
	i := 0
	for key := range m {
		k[i] = key
		i++
	}
	return k
}

// Run service
func (s *Service) Run() {
	log.Println("Starting up...")
	s.RestoreCache()
	(*s.sc).Subscribe("orders", s.acceptOrders)

	log.Println("Service started")
	http.ListenAndServe(":8000", handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "OPTIONS"}),
		handlers.AllowCredentials(),
		handlers.AllowedHeaders([]string{"Content-Type"}))(s.mx))
}

// Load all data from postgres
func (s *Service) RestoreCache() {
	log.Println("Restoring cache from database")
	var orders []Order
	s.db.Preload(clause.Associations).Find(&orders)
	for _, o := range orders {
		order := o
		log.Println(order.OrderUID)
		s.cache[order.OrderUID] = &order
	}
	log.Printf("Cache size: %d\n", len(s.cache))
}

func (s *Service) Stop() {
	log.Println("Shutting down...")
	(*s.sc).Close()
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
	s.db.Preload(clause.Associations).Last(&order)
	fmt.Println(order)
	s.cache[order.OrderUID] = &order
}
