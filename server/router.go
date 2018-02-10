package server

import (
	"github.com/gorilla/mux"
	"github.com/contacts_api_go/handler"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")
	return router
}
