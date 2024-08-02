package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env variables: ", err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port not found: ", err)
	}

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.HandleFunc("GET /v1/healthz", handlerReadiness)
	mux.HandleFunc("GET /v1/error", handlerError)

	fmt.Println("Server listening on port: ", port)
	log.Fatal(server.ListenAndServe())
}
