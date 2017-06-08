package action

import (
	"fmt"
	"github.com/TechChallengeLyke/EmailService/data"
	"github.com/TechChallengeLyke/EmailService/provider"
	"github.com/pkg/errors"
	"net/mail"
)

var (
	providerPriorityList = []string{"SEND_GRID", "MAIL_GUN"}
)

func InitializeImplementations() (map[string]provider.EmailProvider, error) {

	providerList := make(map[string]provider.EmailProvider)

	sendGridImpl, err := provider.InitializeSendGrid()
	if err == nil {
		providerList["SEND_GRID"] = *sendGridImpl
	}
	mailGunImpl, err := provider.InitializeMailGun()
	if err == nil {
		providerList["MAIL_GUN"] = *mailGunImpl
	}

	if len(providerList) == 0 {
		return nil, errors.New("Could not initialize any email providers")
	}

	return providerList, nil
}

func processEmail(providerList map[string]provider.EmailProvider, email data.Email) {
	fmt.Printf("Processing email: %v \n", email.MailID)

	for _, providerName := range providerPriorityList {
		if provider, ok := providerList[providerName]; ok {
			err := provider.SendMail(&email)
			if err == nil {
				email.Status = data.SUCCESS
				email.Update()
				return
			} else {
				fmt.Printf("Sending email(%v) via %v failed. Trying next provider \n", email.MailID, providerName)
			}
		}
	}

	//this will only be executed, if no provider was successful
	email.Status = data.FAILURE
	email.Update()
}

func CheckEmailData(email *data.Email) error {

	if email.FromName == "" {
		return errors.New("email From Name missing")
	}
	if email.FromAddress == "" {
		return errors.New("email From Address missing")
	}
	if email.ToName == "" {
		return errors.New("email To Name missing")
	}
	if email.ToAddress == "" {
		return errors.New("email To Address missing")
	}
	if email.Subject == "" {
		return errors.New("email subject missing")
	}
	if email.BodyText == "" && email.BodyHtml == "" {
		return errors.New("email body (html and text) missing")
	}

	_, err := mail.ParseAddress(email.FromAddress)
	if err != nil {
		return errors.New("invalid From Address used")
	}

	_, err = mail.ParseAddress(email.ToAddress)
	if err != nil {
		return errors.New("invalid To Address used")
	}

	return nil
}
