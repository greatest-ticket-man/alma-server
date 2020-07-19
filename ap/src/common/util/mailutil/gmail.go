package mailutil

import (
	"alma-server/ap/src/common/config"
	"fmt"
	"net/smtp"
)

const gmailSMTPServer = "smtp.gmail.com"

// SendGmail Gmailを送信
func SendGmail(sendMailAddress string, subject string, message string) error {

	gmailConfig := config.ConfigData.Mail.Gmail

	auth := smtp.PlainAuth(
		"",
		gmailConfig.Address,
		gmailConfig.Password,
		gmailSMTPServer,
	)

	return smtp.SendMail(
		gmailSMTPServer+":587",
		auth,
		gmailConfig.Address,
		[]string{sendMailAddress},
		[]byte(fmt.Sprintf("To: %s\r\nSubject:%s\r\n\r\n%s", sendMailAddress, subject, message)),
	)
}
