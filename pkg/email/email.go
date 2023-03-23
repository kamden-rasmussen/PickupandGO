package email

import (
	"log"
	"net/smtp"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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

	auth := smtp.PlainAuth("", from, pass, "smtp.gmail.com")

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + //SUBJECT HERE +
		body

	err = smtp.SendMail("smtp.gmail.com:587",
		auth,
		// auth := smtp.PlainAuth("", "john.doe@gmail.com", "extremely_secret_pass", "smtp.gmail.com")
		from, 
		[]string{to}, 
		[]byte(msg))

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

	// who
	from := mail.NewEmail("Pick Up and GO", senderEmail)
	subject := "Test Email from Pickup and Go"
	to := mail.NewEmail("TestUser", email)

	// content
	plainTextContent := "and easy to do anywhere, even with Go" + body
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"

	// setup
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	// SHIPIT
	response, err := client.Send(message)
	log.Printf("Sent Email with response code:" + strconv.Itoa(response.StatusCode))
	if err != nil {
		log.Println(err)
	} else {
		log.Println(response.StatusCode)
		log.Println(response.Body)
		log.Println(response.Headers)
	}
	return "Email Sent Successfully to " + email + ""
}