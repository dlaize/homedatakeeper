package util

import (
	"net/http"
	"log"
	"encoding/json"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	if code != http.StatusNotFound {
		log.Printf(message)
	}
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}