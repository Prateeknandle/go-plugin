package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Prateeknandle/go-plugin/apis"
	"github.com/Prateeknandle/go-plugin/database"
	"github.com/joho/godotenv"
)

func main() {
	//loading environmental variables
	err := godotenv.Load("example.env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	//connecting to database
	go func() {
		database.DBinstance(ctx)
	}()

	defer func() {
		cancel()
	}()
	//registering routes
	router := apis.NewRouter()

	//connecting to the port
	log.Println("server is listening on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
