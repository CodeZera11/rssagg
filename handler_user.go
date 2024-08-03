package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/codezera11/rssagg/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	reqParams := Request{}

	err := decoder.Decode(&reqParams)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding params!")
		return
	}

	id := uuid.New()

	data := database.CreateUserParams{
		ID:        id,
		Name:      reqParams.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	user, err := cfg.DB.CreateUser(r.Context(), data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (cfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {

	key, err := extractKeyFromHeader(r.Header, "ApiKey")

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("key:", key)

	user, err := cfg.DB.GetUserByApiKey(r.Context(), key)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func extractKeyFromHeader(h http.Header, key string) (string, error) {
	val := h.Get("authorization")

	if val == "" || len(val) < 2 {
		return "", errors.New("key not found")
	}

	result := strings.Split(val, " ")

	if result[0] != key {
		return "", errors.New("incorrect key name")
	}

	return result[1], nil
}
