package repository

import (
	"github.com/contacts_api_go/models"
)

type ContactRepository interface{
	Create(contact models.Contact) int64
	Get(id int64) (*models.Contact, error)
}


type contactRepository struct {
	contactDB *ContactDB
}


func(self contactRepository) Create(contact models.Contact) int64{
	return self.contactDB.Save(contact)
}

func(self contactRepository) Get(id int64) (*models.Contact, error){
	return self.contactDB.Get(id)
}

func NewContactRepo(db *ContactDB) ContactRepository{
	return &contactRepository{
		db,
	}
}