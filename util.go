package main

import (
	"net/http"
	"encoding/json"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func InSlice(e int, s []int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}