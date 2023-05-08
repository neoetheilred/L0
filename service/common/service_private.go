package common

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm/clause"
)

func (s *Service) configureRoutes() {
	s.mx.HandleFunc("/order/{id}", s.handleFindOrderById)

	s.mx.HandleFunc("/cached", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(s.cache)
	})

	s.mx.HandleFunc("/available", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(keys(s.cache))
	})
}

// Gives back Order with corresponding order_uid (`id` in request url)
func (s *Service) handleFindOrderById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.Println(id)
	if order, ok := s.cache[id]; ok {
		log.Println("cache")
		log.Printf("Returing order with uid: %s\n", order.OrderUID)
		json.NewEncoder(w).Encode(order)
	} else {
		var order Order
		r := s.db.Preload(clause.Associations).First(&order, "order_uid = ?", id)
		if r.Error != nil {
			log.Printf("Not found order with id %s\n", id)
			// w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"Error": "Not found"})
			return
		} else {
			log.Printf("Returing order with uid: %s\n", order.OrderUID)
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

// Load all data from postgres
func (s *Service) restoreCache() {
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
