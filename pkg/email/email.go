package email

import (
	"log"
	"os"
	"strconv"

	"github.com/Kamden-Rasmussen/PickupandGO/pkg/data"
	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(user data.User, prices []float64, destination int) int{
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load env. Err: %s", err)
	}

	senderEmail := os.Getenv("SENDING_EMAIL")

	destString := data.GetAirportById(destination)

	// who
	from := mail.NewEmail("Pick Up and GO", senderEmail)
	subject := "Pickup and Go Alert - Price Drop for destination: " + "!" // TODO: add destination
	to := mail.NewEmail(user.FirstName + " " + user.LastName, user.Email)

	// content
	plainTextContent, htmlContent := SetupEmail(user, destString, data.GetDepartureDate(), data.GetReturnDate(), prices)

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
	return response.StatusCode
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

	testUser := data.User{
		FirstName: "Test",
		LastName: "User",
		Email: email,
	}

	prices := []float64{100.00, 200.00, 300.00, 50.00, 400.00, 500.00}

	// content
	plainTextContent, htmlContent := SetupEmail(testUser, "Party Place", "2020-01-01", "2020-01-01", prices)

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

func SetupEmail(user data.User, location string, departureDate string, returnDate string, prices []float64) (body string, html string) {

	currentPrice := prices[0]
	sevenDay := prices[1]
	fourteenDay := prices[2]
	thirtyDay := prices[3]
	ninetyDay := prices[4]
	oneYear := prices[5]

	// body
	body = "Hello " + user.FirstName + ",\n\n"
	body += "A Flight you are tracking to " + location + " has had a signifigant price drop!\n\n"
	body += "The flight details are as follows \n"
	body += "Departure Date: " + departureDate + "\n"
	body += "Return Date: " + returnDate + "\n"
	body += "Current Price: " + strconv.FormatFloat(float64(currentPrice), 'f', 2, 64) + "\n\n"
	body += "Price History \n"
	body += "Todays Price: " + strconv.FormatFloat(float64(currentPrice), 'f', 2, 64) + "\n"
	body += "7 Days: " + strconv.FormatFloat(float64(sevenDay), 'f', 2, 64) + "\n"
	body += "14 Days: " + strconv.FormatFloat(float64(fourteenDay), 'f', 2, 64) + "\n"
	body += "30 Days: " + strconv.FormatFloat(float64(thirtyDay), 'f', 2, 64) + "\n"
	body += "90 Days: " + strconv.FormatFloat(float64(ninetyDay), 'f', 2, 64) + "\n"
	body += "1 Year: " + strconv.FormatFloat(float64(oneYear), 'f', 2, 64) + "\n"


	color := "green"
	// html
	html = "<h1>Hello " + user.FirstName + ",</h1>"
	html += "<p>A Flight you are tracking to " + location + " has had a signifigant price drop!</p>"
	html += "<p> Details of the flight are as follows </p>"
	html += "<p>Departure Date: " + departureDate + "</p>"
	html += "<p>Return Date: " + returnDate + "</p>"
	html += "<p>Current Price: " + "$" + strconv.FormatFloat(float64(currentPrice), 'f', 2, 64) + "</p>"
	html += "<p>The flight History is as follows :</p>"
	// table of price checks
	html += "<table style=\"width:100%\">"
	html += "<tr>"
	html += "<th>Price Check</th>"
	html += "<th>Price</th>"
	html += "</tr>"
	html += "<tr>"
	html += "<td style=\"text-align: center\">Todays Price</td>"
	html += "<td style=\"text-align: center\" style=\"color: " + color + "\">" + "$" + strconv.FormatFloat(float64(currentPrice), 'f', 2, 64) + "</td>"
	html += "</tr>"
	html += "<tr>"
	html += "<td style=\"text-align: center\">7 Days</td>"
	if currentPrice < sevenDay {
		color = "red"
	} else {
		color = "green"
	}
	html += "<td style=\"text-align: center\" style=\"color: " + color + "\">" + "$" + strconv.FormatFloat(float64(sevenDay), 'f', 2, 64) + "</td>"
	html += "</tr>"
	html += "<tr>"
	html += "<td style=\"text-align: center\">14 Days</td>"
	if currentPrice < fourteenDay {
		color = "red"
	} else {
		color = "green"
	}
	html += "<td style=\"text-align: center\" style=\"color: " + color + "\">" + "$" + strconv.FormatFloat(float64(fourteenDay), 'f', 2, 64) + "</td>"
	html += "</tr>"
	html += "<tr>"
	html += "<td style=\"text-align: center\">30 Days</td>"
	if currentPrice < thirtyDay {
		color = "red"
	} else {
		color = "green"
	}
	html += "<td style=\"text-align: center\" style=\"color: " + color + "\">" + "$" + strconv.FormatFloat(float64(thirtyDay), 'f', 2, 64) + "</td>"
	html += "</tr>"
	html += "<tr>"
	html += "<td style=\"text-align: center\">90 Days</td>"
	if currentPrice < ninetyDay {
		color = "red"
	} else {
		color = "green"
	}
	html += "<td style=\"text-align: center\" style=\"color: " + color + "\">" + "$" + strconv.FormatFloat(float64(ninetyDay), 'f', 2, 64) + "</td>"
	html += "</tr>"
	html += "<tr>"
	html += "<td style=\"text-align: center\">1 Year</td>"
	if currentPrice < oneYear {
		color = "red"
	} else {
		color = "green"
	}
	html += "<td style=\"text-align: center\" style=\"color: " + color + "\">" + "$" + strconv.FormatFloat(float64(oneYear), 'f', 2, 64) + "</td>"
	html += "</tr>"
	html += "</table>"

	// html += "<p>Click <a href=\"https://www.google.com/flights?hl=en#flt=/m/0h8j./m/0h8j." + departureDate + "*/m/0h8j." + returnDate + ";c:USD;e:1;sd:1;t:f\">here</a> to view the flight on Google Flights</p>"



	return body, html
}
