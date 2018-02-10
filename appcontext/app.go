package appcontext

import "github.com/contacts_api_go/repository"

var contactDB *repository.ContactDB

var contactRepo repository.ContactRepository


func InitializeDB() {
	contactDB = repository.NewContactDB()
	contactRepo = repository.NewContactRepo(contactDB)
}

func GetContactRepo() repository.ContactRepository{
	return contactRepo
}