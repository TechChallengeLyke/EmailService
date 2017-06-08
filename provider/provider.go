package provider

import "github.com/TechChallengeLyke/EmailService/data"

type EmailProvider interface {
	SendMail(mail *data.Email) error
}
