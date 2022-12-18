package controllers

import (
	"encoding/json"
	"net/http"
)

// encodeResponseAsJson encodes the User object as JSON
func encodeResponseAsJson(data interface{}, w http.ResponseWriter) {
	enc := json.NewEncoder(w)
	enc.Encode(data)

}
