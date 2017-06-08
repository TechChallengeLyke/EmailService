package data

import "time"

type Status string

const (
	CREATED Status = "CREATED"
	SUCCESS Status = "SUCCESS"
	FAILURE Status = "FAILURE"
)

type Email struct {
	MailID      string    `json:"MailId"`
	CreatedAt   time.Time `json:"CreatedAt"`
	FromName    string    `json:"FromName"`
	FromAddress string    `json:"FromAddress"`
	Subject     string    `json:"Subject"`
	ToName      string    `json:"ToName"`
	ToAddress   string    `json:"ToAddress"`
	BodyText    string    `json:"BodyText"`
	BodyHtml    string    `json:"BodyHtml"`
	Status      Status    `json:"Status"`
}

func (email *Email) Create() error {

	email.Status = CREATED
	email.MailID = "blabla"
	email.CreatedAt = time.Now().UTC()

	return nil
}

func (email *Email) Update() error {
	return nil
}
