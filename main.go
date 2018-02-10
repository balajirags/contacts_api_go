package main

import (
	"github.com/contacts_api_go/config"
	"github.com/contacts_api_go/logger"
	"github.com/contacts_api_go/server"
)

func main() {
	config.Load()
	logger.InitLogger(config.GetLogLevel())
	server.StartServer()
}



