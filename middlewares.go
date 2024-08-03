package main

import (
	"net/http"

	"github.com/codezera11/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) authMiddleware(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := extractKeyFromHeader(r.Header, "ApiKey")

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		user, err := cfg.DB.GetUserByApiKey(r.Context(), key)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		handler(w, r, user)
	}
}
