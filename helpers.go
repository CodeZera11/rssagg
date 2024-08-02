package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.WriteHeader(code)

	data, err := json.Marshal(payload)

	if err != nil {
		fmt.Println("Error marshalling data: ", err)
		return
	}

	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)

	type ErrorType struct {
		Error string `json:"error"`
	}

	payload := ErrorType{
		Error: msg,
	}

	data, err := json.Marshal(payload)

	if err != nil {
		fmt.Println("Error marshalling data: ", err)
		return
	}

	w.Write(data)
}
