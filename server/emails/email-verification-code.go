package emails

import (
	"context"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

func (e *Mail) SendEmailVerificationCode() (string, string, error) {
	mg := mailgun.NewMailgun(e.Domain, e.ApiKey)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	t, err := mg.GetTemplate(context.Background(), "email-verification-code")
	if err != nil {
		return "", "", err
	}
	time.Sleep(time.Second * 1)
	message := mg.NewMessage(e.Sender, e.Subject, "", e.Recipients...)
	message.SetTemplate(t.Name)
	if errx := message.AddVariable("code", e.Code); errx != nil {
		return "", "", errx
	}
	msg, id, err := mg.Send(ctx, message)
	if err != nil {
		return "", "", err
	}
	return msg, id, err

}
