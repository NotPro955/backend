package httpx

import (
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}
