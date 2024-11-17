package notifications

import (
	"fmt"
	"net/smtp"
	cfg "smartsme-notificationservice/pkg/settings"

	"strconv"
)

// SendEmail sends an email notification using SMTP or a third-party service.
func SendEmail(to string, subject string, body string) error {
	from := cfg.Config.EMail.Id
	password := cfg.Config.EMail.Pwd
	smtpHost := cfg.Config.EMail.SmtpHost
	smtpPort := strconv.Itoa(cfg.Config.EMail.SmtpPort)

	// Setup the authentication information.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Create the email message.
	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", from, to, subject, body)

	// Send the email.
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
}
