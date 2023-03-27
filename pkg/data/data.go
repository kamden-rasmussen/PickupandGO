package data

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Kamden-Rasmussen/PickupandGO/pkg/mydatabase"
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
	OneWay bool `json:"oneWay,omitempty"`
	NumberOfBookableSeats int `json:"numberOfBookableSeats,omitempty"`
	Itineraries []Itinerary `json:"itineraries,omitempty"`
	Pricing Pricing `json:"price,omitempty"`
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
	Id string `json:"id,omitempty"`
	NumberOfStops int `json:"numberOfStops,omitempty"`
}

type Departure struct {
	IATACode string `json:"iataCode,omitempty"`
	At string `json:"at,omitempty"`
}

type Arrival struct {
	IATACode string `json:"iataCode,omitempty"`
	At string `json:"at,omitempty"`
}
type Pricing struct {
	Currency string `json:"currency,omitempty"`
	Total string `json:"total,omitempty"`
	Base string `json:"base,omitempty"`
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
	body, err := io.ReadAll(resp.Body)
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

	return &OAuthToken
}

func GetExampleData() {
	// get auth token
	OAuth2 := GetAuth()

	// get flights
	flightOffers := GetFlights(OAuth2.Token, "SGU", "SLC", "2023-12-01", "2023-12-02", "1", "DL", "USD")

	// print flights
	log.Println(flightOffers)
	for _, flightOffer := range flightOffers.Data {
		printable := flightOffer.ID
		log.Println(printable)
		SaveData(flightOffer)
	}
}

func GetData(home string, destination string, departure_date string, return_date string, passengers string, airline_code string, currency string) (*FlightOffers) {
	// get auth token
	OAuth2 := GetAuth()

	// get flights
	flightOffers := GetFlights(OAuth2.Token, home, destination, departure_date, return_date, passengers, airline_code, currency)

	for _, flightOffer := range flightOffers.Data {
		SaveData(flightOffer)
	}

	return flightOffers
}

func SaveData(flightOffer FlightOffer){

	// convert itineraries to json
	itineraries, err := json.Marshal(flightOffer.Itineraries)
	if err != nil {
		log.Fatal("error converting itineraries to json: ", err)
	}

	// get airports 
	departureAirport := GetAirport(flightOffer.Itineraries[0].Segments[0].Departure.IATACode)
	arrivalAirport := GetAirport(flightOffer.Itineraries[0].Segments[len(flightOffer.Itineraries[0].Segments)-1].Arrival.IATACode)
	
	_, err = mydatabase.MyDB.Exec("INSERT INTO flights (price, stops, duration, seats, departure_location, departure_date, arrival_location, return_date, one_way, itineraries) values( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", 
		flightOffer.Pricing.Total, 
		flightOffer.Itineraries[0].Segments[0].NumberOfStops, 
		flightOffer.Itineraries[0].Duration, 
		flightOffer.NumberOfBookableSeats, 
		departureAirport, 
		GetDepartureDate(),
		arrivalAirport, 
		GetReturnDate(),
		flightOffer.OneWay, 
		itineraries)

	if err != nil {
		log.Fatal("error inserting into flights: ", err)
	}
}

func GetAirport(iataCode string) (int){
	var airportID int
	err := mydatabase.MyDB.QueryRow("SELECT id FROM airports WHERE code = ? order by id desc", iataCode).Scan(&airportID)
	if err != nil {
		log.Fatal("error getting airport id: ", err)
	}
	return airportID
}

func GetDepartureDate() (string){
	t := time.Now()
	departureDate := t.AddDate(0, 0, 7).Format("2006-01-02")
	return departureDate
}

func GetReturnDate() (string){
	t := time.Now()
	returnDate := t.AddDate(0, 0, 14).Format("2006-01-02")
	return returnDate
}

func ReoccuringTask() {
	log.Println("Starting Data Collection")
	departureDate := GetDepartureDate()
	returnDate := GetReturnDate()
	users := GetAllUsers()

	for _, user := range users {
		destinations := GetDestinationsForUser(user.Id)
		home := GetAirportById(user.Home)
		for _, destination := range destinations {
			dest := GetAirportById(destination)
			GetData(home, dest, departureDate, returnDate, "2", "DL", "USD")
		}
	}
	log.Println("Data Collection Complete")

}
