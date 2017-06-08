# EmailService

## Setup

> go get github.com/sendgrid/sendgrid-go
> go get github.com/mailgun/mailgun-go
> source config.sh
> make build

## Supported Email Providers

SendGrid
MailGun

## Installation Notes

## Architecture

## Security Considerations

## API Documentation
[API](#api)
----
**Send Email**
----
  Send email via the supported email providers (html and/or text mail supported)

* **URL**

  /sendmail

* **Method:**

  POST

*  **URL Params**

   <_If URL params exist, specify them in accordance with name mentioned in URL section. Separate into optional and required. Document data constraints._>

   **Required:**

   `id=[integer]`

   **Optional:**

   `photo_id=[alphanumeric]`

* **Data Params**

  **Required**

  	`FromName       = [alphanumeric]`
  	`FromAddress    = [alphanumeric]`
  	`Subject        = [alphanumeric]`
  	`ToName         = [alphanumeric]`
  	`ToAddress      = [alphanumeric]`

  	One or both of
  	`BodyText       = [alphanumeric]`
    `BodyHtml       = [alphanumeric]`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `OK`

* **Error Responses:**

  * **Code:** 401 UNAUTHORIZED <br />

  * **Code:** 400 BAD REQUEST <br />
    **Content:** `{ error : "error message goes here" }`

* **Sample Call:**

  <_Just a sample call to your endpoint in a runnable format ($.ajax call or a curl request) - this makes life easier and more predictable._>
