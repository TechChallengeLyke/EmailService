package provider

import (
	"errors"
	"fmt"
	"github.com/TechChallengeLyke/EmailService/data"
	"github.com/sendgrid/sendgrid-go"
	mailhelper "github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
)

type SendGridImpl struct {
	ApiKey   string
	EndPoint string
	Host     string
}

func (sg SendGridImpl) SendMail(mail *data.Email) error {

	//needs to be overwritten, because sendgrid allows only senders from this domain
	fromAddress := mailhelper.NewEmail(mail.FromName, mail.FromAddress)
	toAddress := mailhelper.NewEmail(mail.ToName, mail.ToAddress)
	contentText := mailhelper.NewContent("text/plain", mail.BodyText)
	contentHtml := mailhelper.NewContent("text/html", mail.BodyText)
	sgMail := mailhelper.NewV3MailInit(fromAddress, mail.Subject, toAddress, contentText, contentHtml)

	request := sendgrid.GetRequest(sg.ApiKey, sg.EndPoint, sg.Host)
	request.Method = "POST"
	request.Body = mailhelper.GetRequestBody(sgMail)
	response, err := sendgrid.API(request)

	if err != nil {
		fmt.Printf("Error while sending mail (%v) to mailgun : %v", mail.MailID, err)
		return err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	return nil
}

func InitializeSendGrid() (*SendGridImpl, error) {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	endPoint := os.Getenv("SENDGRID_ENDPOINT")
	host := os.Getenv("SENDGRID_HOST")

	//only initialize if all needed environment variables are set
	if len(apiKey) > 0 && len(endPoint) > 0 && len(host) > 0 {
		return &SendGridImpl{ApiKey: apiKey, EndPoint: endPoint, Host: host}, nil
	} else {
		fmt.Printf("Skipping SendGrid Initialization: environment variable(s) missing: \n\tApiKey : %v\n\tEndPoint: %v\n\tHost :%v\n",
			apiKey, endPoint, host)
		return nil, errors.New("SendGrid environment variable(s) missing")
	}
}
