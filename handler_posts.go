package main

import (
	"net/http"
	"strconv"

	"github.com/codezera11/rssagg/internal/database"
)

func (cfg *apiConfig) handlerGetPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) {

	limitStr := r.URL.Query().Get("limit")
	limit := 10

	if specifiedLimit, err := strconv.Atoi(limitStr); err == nil {
		limit = specifiedLimit
	}

	posts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error fetching posts")
		return
	}

	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
