package repository

import (
	"github.com/contacts_api_go/models"
	"github.com/jmoiron/sqlx"
	"github.com/contacts_api_go/appcontext"
	"github.com/contacts_api_go/logger"
	statsdv2 "gopkg.in/alexcesaro/statsd.v2"
)

type ContactRepository interface {
	Create(contact models.Contact) error
	Get(id int64) (*models.Contact, error)
}

type contactRepository struct {
	Db     *sqlx.DB
	statsd *statsdv2.Client
}

const createContacts = "insert into contacts(first_name, last_name, email, address, phone_number) values ($1,$2,$3,$4,$5)"
const GetContacts = "select * from contacts where id = $1"

func (self contactRepository) Create(contact models.Contact) error {
	tx := self.Db.MustBegin()
	tx.MustExec(createContacts, contact.FirstName, contact.LastName, contact.Email, contact.Address, contact.PhoneNumber)
	err := tx.Commit()
	self.statsd.Increment("contacts.count")
	if err != nil{
		logger.Log.Errorf("Errors inserting into contact DB - %s", err.Error())
	}
	return err;
}

func (self contactRepository) Get(id int64) (*models.Contact, error) {
	contact := &models.Contact{}
	err := self.Db.Get(contact, GetContacts, id)
	if err != nil {
		logger.Log.Errorf("Error retrieving contacts for id - %d. Error is %s ", id, err.Error())
	}
	return contact, err
}

func NewContactRepo() ContactRepository {
	return &contactRepository{
		appcontext.GetDB(),
		appcontext.GetStatsDClient(),
	}
}
