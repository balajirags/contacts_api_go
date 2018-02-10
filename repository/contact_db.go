package repository

import (
	"github.com/contacts_api_go/models"
	"github.com/pkg/errors"
	"github.com/contacts_api_go/logger"
)

type ContactDB struct {
	contacts []models.Contact
	seq int64
}

func(self *ContactDB) Save(c models.Contact) int64{
	self.seq = self.seq + 1;
	c.Id = self.seq
	self.contacts = append(self.contacts, c)
	return self.seq
}

func(self *ContactDB) Get(id int64) (*models.Contact, error){
	logger.Log.Info(self.contacts)
	for _, c := range self.contacts{
		if(c.Id == id){
			return &c, nil
		}
	}
	return nil, errors.New("contact not found")
}


func NewContactDB() *ContactDB{
	db := make([]models.Contact, 200)
	return &ContactDB{
		db,
		700,
	}
}