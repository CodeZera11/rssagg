package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	data, err := json.Marshal(payload)

	if err != nil {
		fmt.Println("Error marshalling data: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
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
