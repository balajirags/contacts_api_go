package server

import (
	"strconv"
	"github.com/contacts_api_go/config"
	"net/http"
	"github.com/contacts_api_go/logger"
)

func StartServer() {
	router := InitRouter()
	port := strconv.Itoa(config.GetAppPort())
	logger.Log.Fatal(http.ListenAndServe(":"+port, router))
}
