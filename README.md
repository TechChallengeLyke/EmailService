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

*SendGrid*<br/>
*MailGun*

## Architecture

The Email Service has a layered architecture.
It contains a couple of web pages, but mostly it is centered around the API.
The API has a frontend component, which contains function calls to
- send emails
- get a list of sent emails
- a simple status function (for monitoring)
- a simple call to retrieve some metrics

The most important/complex call is the sendMail call.
After passing through the routing, which is done with the help of [goji.io](https://goji.io/), all calls are handled in the handler package,
which is responsible for unmarshalling and marshalling of input and output values.

All the actual logic resides in the action package. Upon receiving a sendMail request, those requests will be checked for validity and if it proves to be
valid, it will be persisted (that part is only hinted at and not really implemented).
The usage of the external email provider is decoupled from the sendMail call through the use of a channel.
A number of goroutines (the number can be set as a commandline option) read all requests from that channel and try to send those emails via one of the available
email providers. If it succeeds, the mail will be marked as _SUCCESS_ and the goroutine waits for the next mail. If it fails, it will go through all
available email providers. If all of them fail, the mail will be marked as _FAILURE_.

## Security Considerations

The service on it's own does not contain any authentication/authorization. For the test deployment it is only secured by http basic authentication.
Since the test deployment does not have domain attached, it was not possible to use https, which would be a strict requirement for a live deployment of this
service to ensure confidentiality of the credentials.

## Tests

The amount of actual logic in this micro service is rather limited. Therefore I also kept the amount of unit tests rather low. I think it usually does
not make a lot of sense to put more effort into the tests than into the actual code or to waste time writing tests for trivial code.

To run the tests:

`   go test ./...

## Considerations/Outlook

**Scalability**: The service is scalable in the sense, that the number of parallel goroutines can be adjusted very easily. The database is of course
only hinted at right now and would also need to be built in a way, that it could scale with the rest of the service. It might be possible, that some of those
email providers have limitation on the number of concurrent connections. If that is the case, it needs to be taken care of in the code and would require changes in how
the work is distributed across the goroutines.
One improvement for the service could be to adjust the number of goroutines dynamically with the load. Since goroutines are quite cheap that would only
make sense for some quite extreme cases.

**Features**:
- If the persistence is fully implemented, it would also make sense to check the db for in progress mails during startup.
- Most email providers provide feedback, whether a mail was delivered or not and why. This could be retrieved and be used to maintain a blacklist and
provide useful for other services.
- Another useful feature for users of this service might be support for attachments.
- If this service is used to send mails to mailing groups or send out mass mails, a templating feature could be quite useful.
- ...


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


    One or both of


    BodyText       = [alphanumeric]
    BodyHtml       = [alphanumeric]
```

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `OK`<br />

* **Error Responses:**

  * **Code:** 401 UNAUTHORIZED <br/>

  * **Code:** 400 BAD REQUEST <br/>
    **Content:** `{ error : "error message goes here" }`

  * **Code:** 500 INTERNAL SERVER ERROR <br/>

* **Sample Call:**

    on the local machine:
```
    $ curl -H "Content-Type: application/json" -X POST -d '{"FromName":"Tester", "FromAddress":"test@testing.com", "ToName":"John", "ToAddress":"johnf43@gmx.net", "Subject":"Testmail", "BodyText":"This is a test mail from curl"}' http://localhost:8000/sendmail
```

----
### Get Emails
----
  Get a list of all sent emails and their status

* **URL**

  /getmails/:number/:from

* **Method:**

  GET

* **URL Params**

  **Required**

    number         = [int]

**Optional**

    from           = [int]


* **Success Response:**

  * **Code:** 200 <br/>
    **Content:** array of {"MailId":"XXX","CreatedAt":"YYY","FromName":"ZZZ","FromAddress":"AAA","Subject":"BBB","ToName":"CCC","ToAddress":"DDD","BodyText":"EEE","BodyHtml":"FFFF","Status":"GGG"}}<br/>

* **Error Responses:**

  * **Code:** 401 UNAUTHORIZED <br/>

  * **Code:** 400 BAD REQUEST <br/>
    **Content:** `{ error : "error message goes here" }`

  * **Code:** 500 INTERNAL SERVER ERROR <br/>

* **Sample Call:**

    on the local machine:
```
    $ curl -X GET http://localhost:8000/getmails/1/5
```

----
### Status
----
  Returns OK for monitoring solutions

* **URL**

  /status

* **Method:**

  GET

* **Success Response:**

  * **Code:** 200 <br/>
    **Content:** OK<br/>

* **Error Responses:**

* **Sample Call:**

    on the local machine:
```
    $ curl -X GET http://localhost:8000/status
```

----
### Metrics
----
  Returns number of in progress mails, failed mails and successful mails

* **URL**

  /metrics

* **Method:**

  GET

* **Success Response:**

  * **Code:** 200 <br/>
    **Content:** {"InProgressMails":XXX,"Failures":YYY,"Success":ZZZ}<br/>

* **Error Responses:**

  * **Code:** 500 INTERNAL SERVER ERROR <br/>

* **Sample Call:**

    on the local machine:
```
    $ curl -X GET http://localhost:8000/metrics
```
