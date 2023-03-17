package sendEmail

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail() {
	from := "vinhkiet2001@gmail.com"
	password := "nguyenthimaichi"
	to := []string{"ndakiet1001@gmail.com"}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	msg := []byte("From: john.doe@example.com\r\n" +
		"To: roger.roe@example.com\r\n" +
		"Subject: Test mail\r\n\r\n" +
		"Email body\r\n")

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Email sent successfully")
}
