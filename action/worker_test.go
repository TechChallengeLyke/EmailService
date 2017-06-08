package action

import (
	"github.com/TechChallengeLyke/EmailService/data"
	"testing"
)

func TestCheckEmailData(t *testing.T) {

	email := data.Email{
		FromAddress: "sender@test.com",
		FromName:    "TestSender",
		Subject:     "TestSubject",
		ToAddress:   "receiver@test.com",
		ToName:      "TestReceiver",
		BodyText:    "simple text content",
		BodyHtml: `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
			<html>
			<head>
			<meta http-equiv="content-type" content="text/html; charset=ISO-8859-1">
			</head>
			<body bgcolor="#ffffff" text="#000000">
        			Simple html content<br>
			<div class="moz-signature"><i><br>
			<br>
			Regards<br>
			Tester<br>
			</i></div>
			</body>
		</html>`,
	}

	err := CheckEmailData(&email)
	if err != nil {
		t.Errorf("valid email produced data validation error")
	}

	email2 := email
	email2.BodyHtml = ""
	email2.BodyText = ""

	err = CheckEmailData(&email2)
	if err == nil {
		t.Errorf("missing email body not recognized")
	}

	email2 = email
	email2.ToName = ""
	err = CheckEmailData(&email2)
	if err == nil {
		t.Errorf("missing ToName not recognized")
	}

	email2 = email
	email2.ToAddress = ""
	err = CheckEmailData(&email2)
	if err == nil {
		t.Errorf("missing ToAddress not recognized")
	}

	//.....

	email2 = email
	email2.ToAddress = "test@test."
	err = CheckEmailData(&email2)
	if err == nil {
		t.Errorf("invalid receiver address not recognized")
	}

	email2 = email
	email2.FromAddress = "test@@test.com"
	err = CheckEmailData(&email2)
	if err == nil {
		t.Errorf("invalid sender address not recognized")
	}

}
