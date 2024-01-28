package mail

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendNotificationMail(domain string) error {
	fmt.Println("SendNotificationMail execute")
	// Choose auth method and set it up
	username := os.Getenv("GMAIL_USERNAME")
	password := os.Getenv("GMAIL_PASSWORD")
	receiver := os.Getenv("GMAIL_RECEIVER")

	auth := smtp.PlainAuth("", username, password, "smtp.gmail.com")

	// Here we do it all: connect to our server, set up a message and send it

	to := []string{receiver}

	msg := []byte("To:" + receiver + "\r\n" +

		"Subject: Notification from the system Check Domain Status \r\n" +

		"\r\n" +

		"Domains not available: \r" + domain)

	err := smtp.SendMail("smtp.gmail.com:587", auth, receiver, to, msg)

	if err != nil {

		fmt.Println(err)

	}
	fmt.Println("Send mail successfully")
	return err
}
