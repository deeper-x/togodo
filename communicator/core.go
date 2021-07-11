package communicator

import (
	"log"
	"net/smtp"

	"github.com/deeper-x/togodo/settings"
)

// SendEmail email to given recipient
func SendEmail(rcpt, body string) (bool, error) {
	msg := []byte(body)

	auth := smtp.PlainAuth("", settings.SMTPUser, settings.EmailPasswd, settings.SMTPHOst)

	err := smtp.SendMail(settings.SMTPURI, auth, settings.SMTPUser, []string{rcpt}, msg)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil
}
