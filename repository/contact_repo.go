package repository

import (
	"github.com/contacts_api_go/models"
	"github.com/jmoiron/sqlx"
	"github.com/contacts_api_go/appcontext"
	"github.com/contacts_api_go/logger"
)

type ContactRepository interface {
	Create(contact models.Contact) error
	Get(id int64) (*models.Contact, error)
}

type contactRepository struct {
	Db *sqlx.DB
}

const createContacts = "insert into contacts(first_name, last_name, email, address, phone_number) values ($1,$2,$3,$4,$5)"

func (self contactRepository) Create(contact models.Contact) error {
	tx := self.Db.MustBegin()
	tx.MustExec(createContacts, contact.FirstName, contact.LastName, contact.Email, contact.Address, contact.PhoneNumber)
	err := tx.Commit()
	if err != nil{
		logger.Log.Errorf("Errors inserting into contact DB - %s", err.Error())
	}
	return err;
}

func (self contactRepository) Get(id int64) (*models.Contact, error) {
	return nil, nil
}

func NewContactRepo() ContactRepository {
	return &contactRepository{
		appcontext.GetDB(),
	}
}
