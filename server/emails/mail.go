package emails

import (
	"go-bookmark/util"
	"log"
)

type Mail struct {
	Domain     string   `json:"domain"`
	ApiKey     string   `json:"api_key"`
	Sender     string   `json:"sender"`
	Subject    string   `json:"subject"`
	Recipients []string `json:"recipients"`
	Code       string   `json:"code"`
}

func NewMail(sender, subject, code string, recipients []string) (*Mail, error) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &Mail{
		Domain:     config.MAILGUN_DOMAIN,
		ApiKey:     config.MailgunAPIKey,
		Sender:     sender,
		Subject:    subject,
		Recipients: recipients,
		Code:       code,
	}, nil
}
