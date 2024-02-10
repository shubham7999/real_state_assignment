package utils

import (
	"encoding/json"
	"net/http"
)

func HandleError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
