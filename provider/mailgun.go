package provider

import (
	"errors"
	"fmt"
	"github.com/TechChallengeLyke/EmailService/data"
	mailgun "github.com/mailgun/mailgun-go"
	"os"
)

type MailGunImpl struct {
	Domain       string
	ApiKey       string
	PublicApiKey string
}

func (mg MailGunImpl) SendMail(mail *data.Email) error {

	mailgun := mailgun.NewMailgun(mg.Domain, mg.ApiKey, mg.PublicApiKey)

	//needs to be overwritten, because mailgun allows only this sender
	mail.FromAddress = "Postmaster <postmaster@sandbox9fc95b70809d455481010f214c6eeaaf.mailgun.org>"

	msg := mailgun.NewMessage(mail.FromAddress, mail.Subject, mail.BodyText, mail.ToAddress)
	msg.SetHtml(mail.BodyHtml)

	_, _, err := mailgun.Send(msg)

	if err != nil {
		fmt.Printf("Error while sending mail (%v) to mailgun : %v", mail.MailID, err)
		return err
	}

	return nil
}

func InitializeMailGun() (*MailGunImpl, error) {
	domain := os.Getenv("MAILGUN_DOMAIN")
	apiKey := os.Getenv("MAILGUN_API_KEY")
	publicApiKey := os.Getenv("MAILGUN_PUBLIC_API_KEY")

	//only initialize if all needed environment variables are set
	if len(domain) > 0 && len(apiKey) > 0 && len(publicApiKey) > 0 {
		return &MailGunImpl{Domain: domain, ApiKey: apiKey, PublicApiKey: publicApiKey}, nil
	} else {
		fmt.Printf("Skipping Mailgun Initialization: environment variable(s) missing: \n\tDomain : %v\n\tApiKey: %v\n\tPublicApiKey :%v\n",
			domain, apiKey, publicApiKey)
		return nil, errors.New("MailGun environment variable(s) missing")
	}
}
