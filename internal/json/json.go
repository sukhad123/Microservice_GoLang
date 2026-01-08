package json

import (
	"encoding/json"
	"net/http"
)

// Similar to next request and response
// Strucuted json response
func Write(w http.ResponseWriter, status int, data any) {
	w.WriteHeader(status)
	//set header type
	w.Header().Set("Content-Type:", "application/json")
	json.NewEncoder(w).Encode(data)
}
