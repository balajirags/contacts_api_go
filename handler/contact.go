package handler

import (
	"net/http"
	"github.com/contacts_api_go/logger"
	"encoding/json"
	"github.com/contacts_api_go/models"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/contacts_api_go/repository"
	"time"
	"math/rand"
)

func CreateContact(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("Called create contact endpoint")
	var contact models.Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	time.Sleep(time.Duration(random(500, 4000)) * time.Millisecond)
	contactRepository := repository.NewContactRepo()
	contactRepository.Create(contact)
	w.WriteHeader(http.StatusOK)
}


func GetContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	logger.Log.Info("Called Get contact endpoint", params["id"])
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	contactRepository := repository.NewContactRepo()
	contact, e := contactRepository.Get(id)
	if e!=nil{
		logger.Log.Info("error", e)
		w.WriteHeader(http.StatusInternalServerError)
	}else{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(contact)
	}
}


func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}
