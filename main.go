package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"github.com/contacts_api_go/config"
	"strconv"
)

var Log *logrus.Logger

func main() {
	config.Load()
	initLogger(config.GetLogLevel())
	router := initRouter()
	port := strconv.Itoa(config.GetAppPort())
	log.Fatal(http.ListenAndServe(":"+port, router))
}
func initRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", health).Methods("GET")
	return router
}

func initLogger(logLevel string) {
	level, _ := logrus.ParseLevel(logLevel)
	Log = &logrus.Logger{
		Out:       os.Stdout,
		Level:     level,
		Formatter: &logrus.JSONFormatter{},
	}
}
func health(writer http.ResponseWriter, request *http.Request) {
	Log.Info("Called ping endpoint")
	fmt.Fprint(writer, "pong")
}
