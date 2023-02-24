package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
   "type": "amadeusOAuth2Token",
   "username": "d00306195@utahtech.edu",
   "application_name": "Pick up and GO",
   "client_id": "GwYYrehNTx3xk7ACpDN4bRMKrctdhsrA",
   "token_type": "Bearer",
   "access_token": "Cg19kB9fyPeJGvv0aQB7Uu5QdEMj",
   "expires_in": 1799,
   "state": "approved",
   "scope":
*/

type AmadeusToken struct{
	Username string `json:"username,omitempty"`
	ApplicationName string `json:"application_name,omitempty"`
	ClientID string `json:"client_id,omitempty"`
	TokenType string `json:"token_type,omitempty"`
	Token string `json:"access_token,omitempty"`
	Expires int `json:"expires_in,omitempty"`
	State string `json:"state,omitempty"`
}

type FlightOffers struct {
	Data []FlightOffer `json:"data,omitempty"`
}

type FlightOffer struct {
	ID string `json:"id,omitempty"`
	Source string `json:"source,omitempty"`
	OneWay bool `json:"oneWay,omitempty"`
	LastTicketingDate string `json:"lastTicketingDate,omitempty"`
	NumberOfBookableSeats int `json:"numberOfBookableSeats,omitempty"`
	Itineraries []Itinerary `json:"itineraries,omitempty"`
	Pricing []Pricing `json:"pricing[totalPrice],omitempty"`
	ValidatingAirlineCodes []string `json:"validatingAirlineCodes,omitempty"`
}

type Itinerary struct {
	Duration string `json:"duration,omitempty"`
	Segments []Segment `json:"segments,omitempty"`
}

type Segment struct {
	Departure Departure `json:"departure,omitempty"`
	Arrival Arrival `json:"arrival,omitempty"`
	CarrierCode string `json:"carrierCode,omitempty"`
	Number string `json:"number,omitempty"`
	Operating Operating `json:"operating,omitempty"`
	Duration string `json:"duration,omitempty"`
	Id string `json:"id,omitempty"`
	NumberOfStops int `json:"numberOfStops,omitempty"`
	BlacklistedInEU bool `json:"blacklistedInEU,omitempty"`
}

type Departure struct {
	IATACode string `json:"iataCode,omitempty"`
	Terminal string `json:"terminal,omitempty"`
	At string `json:"at,omitempty"`
}

type Arrival struct {
	IATACode string `json:"iataCode,omitempty"`
	Terminal string `json:"terminal,omitempty"`
	At string `json:"at,omitempty"`
}
type Operating struct {
	CarrierCode string `json:"carrierCode,omitempty"`
}

type Pricing struct {
	Price string `json:"price,omitempty"`
	TotalPrice string `json:"totalPrice,omitempty"`
	BasePrice string `json:"basePrice,omitempty"`
	Taxes []Tax `json:"taxes,omitempty"`
	Fees []Fee `json:"fees,omitempty"`
}

type Tax struct {
	Amount string `json:"amount,omitempty"`
	Type string `json:"type,omitempty"`
}

type Fee struct {
	Amount string `json:"amount,omitempty"`
	Type string `json:"type,omitempty"`
}


func PrintHealth() {
	log.Println("Container is healthy")
}

func PrintLetsGo() {
	log.Println("--------------------- Let's go! ---------------------")
}

func GetAuth() (*AmadeusToken){

	payload := strings.NewReader("client_id=" + os.Getenv("AMADEUS_CLIENT_ID") + "&client_secret=" + os.Getenv("AMADEUS_CLIENT_SECRET") + "&grant_type=client_credentials")

	resp, err := http.Post("https://test.api.amadeus.com/v1/security/oauth2/token", "application/x-www-form-urlencoded", payload)
	if err != nil {
		log.Fatal("error getting oauth token: ", err)
	}
	defer resp.Body.Close()

	// fail if status code is not 200
	if resp.StatusCode != 200 {
		log.Fatal("error getting oauth token: ", resp.StatusCode)
	}
	
	// read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error reading response body: ", err)
	}
	log.Println(string(body))

	// unmarshal response body into struct
	var OAuthToken AmadeusToken
	err = json.Unmarshal(body, &OAuthToken)
	if err != nil {
		log.Fatal("error unmarshalling response body: ", err)
	}
	log.Println(OAuthToken)

	return &OAuthToken
}

func GetData() {
	// get auth token
	OAuth2 := GetAuth()
	log.Println(OAuth2.Token)

	// get flights
	GetFlights(OAuth2.Token, "SGU", "SLC", "2023-12-01", "2023-12-05", "1", "DL", "USD")
	
}

