package handler

import (
	"net/http"
	"github.com/contacts_api_go/logger"
)

func 	PingHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("Called ping endpoint")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\": \"pong10\"}"))
}
