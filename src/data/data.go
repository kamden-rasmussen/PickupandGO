package data

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func PrintHealth() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " " + "Container is healthy")
}

func GetData() {
	_, err := http.Get("https://test.api.amadeus.com/v2/shopping/flight-offers")
	if err != nil {
		log.Fatal(err)
	}

}