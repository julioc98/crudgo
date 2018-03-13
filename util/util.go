package util

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON Write responde HTTP Type application/json
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
