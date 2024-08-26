package main

import (
	"log"

	"github.com/quest-be/http"
	"github.com/quest-be/internal/database"
	"github.com/quest-be/internal/service"
	"github.com/quest-be/util"
)

func main() {
	err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	conn, err := database.New(true)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	router := service.NewRouter(conn)
	server := http.NewHTTP(router)
	server.Run(util.Default.SERVER_ADDRESS)

}
