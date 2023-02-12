package data

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func PrintHealth() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " " + "Container is healthy")
}

func GetAuth() (string){

	body := []byte(`{
		"grant_type": "client_credentials",
		"client_id": ` + os.Getenv("AMADEUS_CLIENT_ID") + `,
		"client_secret": ` + os.Getenv("AMADEUS_CLIENT_SECRET") + `
	}`)
	ioreader := bytes.NewReader(body)
	resp, err := http.Post("https://test.api.amadeus.com/v1/security/oauth2/token", "application/json", ioreader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

	return ""
}

func GetData() {
	
	_, err := http.Get("https://test.api.amadeus.com/v2/shopping/flight-offers")
	if err != nil {
		log.Fatal(err)
	}

}