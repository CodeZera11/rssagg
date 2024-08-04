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

	type Response struct {
		Feed       Feed      `json:"feed"`
		FeedFollow UsersFeed `json:"feed_follow"`
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
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	followedFeed, err := cfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		ID:        uuid.New(),
		FeedID:    feed.ID,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, Response{
		Feed:       databaseFeedToFeed(feed),
		FeedFollow: databaseUsersFeedToUsersFeed(followedFeed),
	})
}

func (cfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, feeds)
}

func (cfg *apiConfig) handlerFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type Parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	defer r.Body.Close()

	params := &Parameters{}

	err := decoder.Decode(params)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := database.FollowFeedParams{
		ID:        uuid.New(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	userFeed, err := cfg.DB.FollowFeed(r.Context(), data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseUsersFeedToUsersFeed(userFeed))
}

func (cfg *apiConfig) handlerUnfollowFeed(w http.ResponseWriter, r *http.Request) {
	feedFollowId := r.PathValue("feedFollowID")

	if feedFollowId == "" {
		respondWithError(w, http.StatusInternalServerError, "Id not found!")
		return
	}

	id := uuid.MustParse(feedFollowId)

	userFeed, err := cfg.DB.DeleteFollowFeed(r.Context(), id)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUsersFeedToUsersFeed(userFeed))
}

func (cfg *apiConfig) handlerGetFollowedFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	followedFeeds, err := cfg.DB.GetFollowedFeeds(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, databaseUsersFeedsToUsersFeeds(followedFeeds))
}
