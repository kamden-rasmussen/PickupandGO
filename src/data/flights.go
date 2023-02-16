package data

import (
	"log"
	"net/http"
)


func GetFlights(
	bearerToken string,
	originLocationCode string, 
	destinationLocationCode string, 
	departureDate string, 
	returnDate string, 
	adults string, 
	airlineCodes string, 
	currencyCode string) {
		
	params := ("?originLocationCode=" + originLocationCode +
		"&destinationLocationCode=" + destinationLocationCode +
		"&departureDate=" + departureDate +
		"&returnDate=" + returnDate +
		"&adults=" + adults +
		"&airlineCodes=" + airlineCodes +
		"&currencyCode=" + currencyCode)

	// get data
	var bearer = "Bearer " + bearerToken
	url := "https://test.api.amadeus.com/v2/shopping/flight-offers" + params

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("error creating request: ", err)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()
}