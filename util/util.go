package util

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

// RespondWithJSON Write responde HTTP Type application/json
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// GetOSEnvironment get a env var or set to default
func GetOSEnvironment(env string, defaultEnv string) string {
	if e := os.Getenv(env); e != "" {
		return strings.ToLower(os.Getenv(env))
	}
	return defaultEnv
}
