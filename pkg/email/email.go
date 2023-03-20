package email

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendEmail(email string, body string) {
	
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load env. Err: %s", err)
	}

	senderEmail := os.Getenv("SENDING_EMAIL")
	senderPassword := os.Getenv("SENDING_EMAIL_PASSWORD")
	
	from := senderEmail
	pass := senderPassword
	to := email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + //SUBJECT HERE +
		body

	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Println("Email Sent Successfully to " + email + "")
}

func SendTestEmail(email string, body string) string{
	
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load env. Err: %s", err)
	}

	senderEmail := os.Getenv("SENDING_EMAIL")
	senderPassword := os.Getenv("SENDING_EMAIL_PASSWORD")
	
	from := senderEmail
	pass := senderPassword
	to := email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: TESTEMAIL FOR PICKUP AND GO" +
		body

	log.Println("Sending Test Email " + msg)

	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return "Error Sending Email"
	}

	log.Println("TEST Email Sent Successfully to " + email + "")
	return "Email Sent Successfully"
}