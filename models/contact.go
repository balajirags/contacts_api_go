package models

type Contact struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber int64  `json:"phone_number,omitempty"`
	Email       int64  `json:"email,omitempty"`
	Address     string `json:"address,omitempty"`
	Id          int64  `json:"id,omitempty"`
}
