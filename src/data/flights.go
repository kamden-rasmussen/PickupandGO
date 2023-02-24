package data

import (
	"encoding/json"
	"io/ioutil"
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
	currencyCode string) *FlightOffers{
		
	params := ("?originLocationCode=" + originLocationCode +
		"&destinationLocationCode=" + destinationLocationCode +
		"&departureDate=" + departureDate +
		"&returnDate=" + returnDate +
		"&adults=" + adults +
		"&includedAirlineCodes=" + airlineCodes +
		"&currencyCode=" + currencyCode)

	log.Println("Starting to get flights")

	// get data
	var bearer = "Bearer " + bearerToken
	log.Println(bearer)

	url := "https://test.api.amadeus.com/v2/shopping/flight-offers" + params

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("error creating request: ", err)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)


	log.Println(req.Header)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	// ready body for unmarshalling
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body.\n[ERROR] -", err)
	}
	
	// fail if status code is not 200
	if resp.StatusCode != 200 {
		log.Println("Error getting flights: ", resp.StatusCode, resp.Body)
		log.Println("response body: ", string(bytes))
	}

	flightOffers := FlightOffers{}
	err = json.Unmarshal(bytes, &flightOffers)
	if err != nil {
		log.Println("error unmarshalling response body: ", err)
	}


	// log.Println(flightOffer)
	return &flightOffers
}