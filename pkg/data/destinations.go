package data

import "database/sql"


func GetDestinationsForUsers(db *sql.DB, users []User) []string {
	// need to do a stored procedure to get all the destinations for a user
	return []string{"Denver", "New York", "Los Angeles"}
}

func GetAirportCodes(db *sql.DB, ids []int) []string {
	return []string{"DEN", "JFK", "LAX"}
}