package main

import (
	"log"

	"github.com/quest-be/http"
	"github.com/quest-be/internal/repository/postgres"
	"github.com/quest-be/internal/service/router"
	"github.com/quest-be/util"
)

func main() {
	err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	conn, err := postgres.New(true)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	err = postgres.Setup(conn)
	if err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}
	router := router.NewRouter(conn)
	server := http.NewHTTP(router)
	server.Run(util.Default.SERVER_ADDRESS)

}
