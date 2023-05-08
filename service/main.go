package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	_ "github.com/lib/pq"
	"github.com/neoetheilred/l0/service/common"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_DBNAME")
)

func main() {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	s, err := common.NewService(connString)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Stop()
	s.Run()
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)
	<-stopChan
}
