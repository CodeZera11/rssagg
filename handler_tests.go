package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {

	type repsonse struct {
		Status string `json:"status"`
	}

	respondWithJSON(w, http.StatusOK, repsonse{Status: "ok"})
}

func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal server error")
}
