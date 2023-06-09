package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"github.com/neoetheilred/l0/service/common"
)

var orderJson = `{
	"order_uid": "b563feb7b2b84b6test",
	"track_number": "WBILMTESTTRACK",
	"entry": "WBIL",
	"delivery": {
	  "name": "Test Testov",
	  "phone": "+9720000000",
	  "zip": "2639809",
	  "city": "Kiryat Mozkin",
	  "address": "Ploshad Mira 15",
	  "region": "Kraiot",
	  "email": "test@gmail.com"
	},
	"payment": {
	  "transaction": "b563feb7b2b84b6test",
	  "request_id": "",
	  "currency": "USD",
	  "provider": "wbpay",
	  "amount": 1817,
	  "payment_dt": 1637907727,
	  "bank": "alpha",
	  "delivery_cost": 1500,
	  "goods_total": 317,
	  "custom_fee": 0
	},
	"items": [
	  {
		"chrt_id": 9934930,
		"track_number": "WBILMTESTTRACK",
		"price": 453,
		"rid": "ab4219087a764ae0btest",
		"name": "Mascaras",
		"sale": 30,
		"size": "0",
		"total_price": 317,
		"nm_id": 2389212,
		"brand": "Vivienne Sabo",
		"status": 202
	  }
	],
	"locale": "en",
	"internal_signature": "",
	"customer_id": "test",
	"delivery_service": "meest",
	"shardkey": "9",
	"sm_id": 99,
	"date_created": "2021-11-26T06:22:19Z",
	"oof_shard": "1"
  }`

type Foo struct {
	Name string `fake:"{name}"`
}

func main() {
	// natsUri := os.Getenv("NATS_URI")
	// natsPort := os.Getenv("NATS_PORT")
	sc, err := stan.Connect("test-cluster", "sender")
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()
	// var f Foo
	// gofakeit.Struct(&f)
	// fmt.Println(f)
	var order common.Order
	// json.Unmarshal([]byte(orderJson), &order)

	// sc.Publish("orders", []byte(order))
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < n; i++ {
		gofakeit.Struct(&order)
		order.OrderUID = uuid.New().String()
		fmt.Println(order.OrderUID)
		marshalled, err := json.Marshal(order)
		if err != nil {
			log.Println(err)
		}
		sc.Publish("orders", marshalled)
	}
}
