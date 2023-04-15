package data

import (
	"log"

	"github.com/Kamden-Rasmussen/PickupandGO/pkg/mydatabase"
)

// returns a list of all destinations
func GetDestinationsForUser(userID int) []int {
	rows, err := mydatabase.MyDB.Query("SELECT * FROM destinations WHERE user_id = ?", userID)
	if err != nil {
		log.Fatalf("error getting destinations for user: %v", err)
	}
	var destinations []int
	for rows.Next() {
		var userID int
		var airportID int
		err = rows.Scan(&userID, &airportID)
		if err != nil {
			log.Fatalf("error scanning destinations: %v", err)
		}
		destinations = append(destinations, airportID)
	}
	return destinations
}

func GetAirportById(id int) string{
	rows, err := mydatabase.MyDB.Query("SELECT * FROM airports WHERE id = ? order by id desc", id)
	if err != nil {
		log.Fatalf("error getting airport: %v", err)
	}
	var name string
	var code string
	for rows.Next() {
		err = rows.Scan(&id, &name, &code)
		if err != nil {
			log.Fatalf("error scanning airport: %v", err)
		}
	}
	return code 

}