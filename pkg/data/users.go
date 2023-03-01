package data

import (
	"database/sql"
	"log"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Home      int
	Currency  string
}

func GetAllUsers(db *sql.DB) []User {
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err.Error())
	}
	var users []User
	for results.Next() {
		var id int
		var firstname string
		var lastname string
		var email string
		var home int
		var currency string
		err = results.Scan(&id, &firstname, &lastname, &email, &home, &currency)
		if err != nil {
			log.Fatal(err.Error())
		}
		users = append(users, User{id, firstname, lastname, email, home, currency})
	}
	return users
}