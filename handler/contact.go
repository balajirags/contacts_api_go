package handler

import (
	"net/http"
	"github.com/contacts_api_go/logger"
	"encoding/json"
	"github.com/contacts_api_go/models"
	"strconv"
	"github.com/contacts_api_go/appcontext"
	"github.com/gorilla/mux"
)

func CreateContact(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("Called create contact endpoint")
	var contact models.Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	contactId := appcontext.GetContactRepo().Create(contact)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.FormatInt(contactId,10)))
}


func GetContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	logger.Log.Info("Called Get contact endpoint", params["id"])
	i, _ := strconv.ParseInt(params["id"], 10, 64)
	contact, e := appcontext.GetContactRepo().Get(i)
	if e!=nil{
		logger.Log.Info("error", e)
		w.WriteHeader(http.StatusInternalServerError)
	}else{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(contact)
	}
}
