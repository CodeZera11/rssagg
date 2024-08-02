package main

import (
	"fmt"
	"log"
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
		log.Fatal("Port not mentioned")
	}

	fmt.Println("Hello world")
}
