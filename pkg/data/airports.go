package data

import "github.com/Kamden-Rasmussen/PickupandGO/pkg/mydatabase"

// returns a list of all destinations
func GetDestinationsForUser(userID int) []int {
	rows, err := mydatabase.MyDB.Query("SELECT * FROM destinations WHERE user_id = ?", userID)
	if err != nil {
		panic(err.Error())
	}
	var destinations []int
	for rows.Next() {
		var userID int
		var airportID int
		err = rows.Scan(&userID, &airportID)
		if err != nil {
			panic(err.Error())
		}
		destinations = append(destinations, airportID)
	}
	return destinations
}

func GetAirportById(id int) string{
	rows, err := mydatabase.MyDB.Query("SELECT * FROM airports WHERE id = ? order by id desc", id)
	if err != nil {
		panic(err.Error())
	}
	var name string
	var code string
	for rows.Next() {
		err = rows.Scan(&id, &name, &code)
		if err != nil {
			panic(err.Error())
		}
	}
	return code 

}