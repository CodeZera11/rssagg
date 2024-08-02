package main

import (
	"encoding/json"
	"net/http"
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

	respondWithJSON(w, http.StatusOK, user)
}
