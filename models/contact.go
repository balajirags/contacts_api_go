package models

type Contact struct {
	FirstName   string `json:"first_name,omitempty" db:"first_name" `
	LastName    string `json:"last_name,omitempty" db:"last_name"`
	PhoneNumber int64  `json:"phone_number,omitempty" db:"phone_number"`
	Email       string  `json:"email,omitempty" db:"email""`
	Address     string `json:"address,omitempty" db:"address"`
	Id          int64  `json:"id" db:"id"`
}
