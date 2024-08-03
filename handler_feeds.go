package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/codezera11/rssagg/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type Request struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	reqParams := Request{}
	err := decoder.Decode(&reqParams)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding params")
		return
	}

	data := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      reqParams.Name,
		Url:       reqParams.Url,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding params")
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedToFeed(feed))
}

func (cfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, feeds)
}
