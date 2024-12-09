package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	log.Printf("Writing response with status code: %d", statusCode)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
