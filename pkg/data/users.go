package data

import (
	"database/sql"
	"log"

	"github.com/Kamden-Rasmussen/PickupandGO/pkg/mydatabase"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Home      int
	Currency  string
}

func GetAllUsers() []User {
	results, err := mydatabase.MyDB.Query("SELECT * FROM users")
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

func AddUser(user User) {
	_, err := mydatabase.MyDB.Exec("INSERT INTO users (firstname, lastname, email, home, currency) values( ?, ?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Home, user.Currency)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetOneUser(db *sql.DB, id int) User {
	results, err := mydatabase.MyDB.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		log.Fatal(err.Error())
	}
	var user User
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
		user = User{id, firstname, lastname, email, home, currency}
	}
	return user
}