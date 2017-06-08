package data

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"time"
)

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

var dataStore = []Email{}

//THIS IS JUST TEST CODE
//DON'T USE IN PRODUCTION
//NOT THREAD-SAFE

func (email *Email) Create() error {

	email.Status = CREATED
	email.MailID = uuid.NewV4().String()
	email.CreatedAt = time.Now().UTC()

	dataStore = append(dataStore, *email)
	return nil
}

func (email *Email) Update() error {
	for i, currentMail := range dataStore {
		if currentMail.MailID == email.MailID {
			dataStore[i] = *email
			return nil
		}
	}
	return errors.New("email " + email.MailID + " not found")
}

func GetMails(from int, number int) *[]Email {

	mails := []Email{}

	if len(dataStore) <= from {
		return &mails
	}

	for i := from; i < len(dataStore) && i < from+number; i++ {
		mails = append(mails, dataStore[i])
	}
	return &mails
}
