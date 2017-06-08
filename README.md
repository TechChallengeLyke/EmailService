# EmailService

## Setup

    $ go get github.com/sendgrid/sendgrid-go<br />
    $ go get github.com/mailgun/mailgun-go<br />
    $ go get github.com/satori/go.uuid<br />
    $ go get goji.io<br />
    $ make build<br />

## Supported Email Providers

SendGrid
MailGun

## Installation Notes

> source config.sh<br />

## Architecture

## Security Considerations

## API Documentation
----
###Send Email
----
  Send email via the supported email providers (html and/or text mail supported)

* **URL**

  /sendmail

* **Method:**

  POST

* **Data Params**

  **Required**

  	`FromName       = [alphanumeric]`<br />
  	`FromAddress    = [alphanumeric]`<br />
  	`Subject        = [alphanumeric]`<br />
  	`ToName         = [alphanumeric]`<br />
  	`ToAddress      = [alphanumeric]`<br />
<br />
  	One or both of
<br />
  	`BodyText       = [alphanumeric]`<br />
    `BodyHtml       = [alphanumeric]`<br />

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `OK`<br />

* **Error Responses:**

  * **Code:** 401 UNAUTHORIZED <br />

  * **Code:** 400 BAD REQUEST <br />
    **Content:** `{ error : "error message goes here" }`

* **Sample Call:**

    on the local machine:
    $ curl -H "Content-Type: application/json" -X POST -d '{"FromName":"Tester", "FromAddress":"test@testing.com", "ToName":"John", "ToAddress":"johnf43@gmx.net", "Subject":"Testmail", "BodyText":"This is a test mail from curl"}' http://localhost:8000/sendmail

