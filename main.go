package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codezera11/rssagg/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env variables: ", err)
	}

	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DATABASE_URL")

	if port == "" {
		log.Fatal("Port not found:", err)
	}

	if dbUrl == "" {
		log.Fatal("Database url not found:", err)
	}

	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal("Error opening connection to db: ", err)
	}

	defer db.Close()

	dbQueries := database.New(db)

	cfg := apiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// SERVER TESTS ENDPOINTS
	mux.HandleFunc("GET /v1/healthz", handlerReadiness)
	mux.HandleFunc("GET /v1/error", handlerError)

	// USERS ENDPOINTS
	mux.HandleFunc("POST /v1/users", cfg.handlerCreateUser)
	mux.HandleFunc("GET /v1/users", cfg.authMiddleware(cfg.handlerGetUser))

	// FEEDS ENDPOINTS
	mux.HandleFunc("POST /v1/feeds", cfg.authMiddleware(cfg.handlerCreateFeed))
	mux.HandleFunc("GET /v1/feeds", cfg.handlerGetFeeds)
	mux.HandleFunc("POST /v1/feed_follows", cfg.authMiddleware(cfg.handlerFeedFollow))

	fmt.Println("Server listening on port:", port)
	log.Fatal(server.ListenAndServe())
}
