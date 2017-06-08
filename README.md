# EmailService

## Setup
    ```
    $ go get github.com/sendgrid/sendgrid-go
    $ go get github.com/mailgun/mailgun-go
    $ go get github.com/satori/go.uuid
    $ go get goji.io
    $ make build
    ```
## Supported Email Providers

SendGrid<br/>
MailGun

## Installation Notes

> source config.sh<br/>

## Architecture

The Email Service has a layered architecture.
It contains a few static web pages, but mostly it is centered around the API.
The API has a frontend component, which contains function calls to send emails, get a list of sent emails, a simple status function (for monitoring) and
a simple call to retrieve some metrics.
The most important/complex call is the sendMail call.
After passing through the routing, which is done with the help of [goji.io](https://goji.io/), all calls are handled in the handler package,
which is responsible for unmarshalling and marshalling of input and output values.

All the actual logic resides in the action package. Upon receiving a sendMail request, those requests will be checked for validity and if it proves to be
valid, it will be persisted (that part is only hinted at and not really implemented).
The usage of the external email provider is decoupled from the sendMail call through the use of a channel.
A number of goroutines (which can be set as a commandline option) reads all requests from that channel and tries to send those emails via one of the available
email providers. If it succeeds, the mail will be marked as _SUCCESS_ and the goroutine waits for the next mail. If it fails, it will go through all
available email providers. If all of them fail, the mail will be marked as _FAILURE_.

## Security Considerations

The service on it's own does not contain any authentication/authorization. For the test deployment it is only secured by http basic authentication.
Since the test deployment does not have domain attached, it was not possible to use https, which would be strict requirement for a live deployment of this
service to ensure confidentiality of the credentials.

## API Documentation
----
### Send Email
----
  Send email via the supported email providers (html and/or text mail supported)

* **URL**

  /sendmail

* **Method:**

  POST

* **Data Params**

  **Required**
    ```
  	FromName       = [alphanumeric]
  	FromAddress    = [alphanumeric]
  	Subject        = [alphanumeric]
  	ToName         = [alphanumeric]
  	ToAddress      = [alphanumeric]
  	```
<br/>
  	One or both of
<br/>
    ```
  	BodyText       = [alphanumeric]
    BodyHtml       = [alphanumeric]
    ```
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

