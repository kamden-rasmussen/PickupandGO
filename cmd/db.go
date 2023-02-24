package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func ConnectToDatabase() (*sql.DB, error) {
	// connect to a database
	config := mysql.Config{
        User:      os.Getenv("DBACCESSUSERNAME"),
		Passwd:    os.Getenv("DBACCESSPASSWORD"),
		Net:       "tcp",
		Addr:      os.Getenv("DBHOST"),
		DBName:    os.Getenv("DBDATABASENAME"),
    }
	db, err := sql.Open("mysql", config.FormatDSN())
	pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
	return db, err
}

func GetFunTable(db *sql.DB) {
	results, err := db.Query("SELECT * FROM fun")
	if err != nil {
		log.Fatal(err.Error())
	}
	for results.Next() {
		var id int
		var name string
		var color string
		err = results.Scan(&id, &name, &color)
		if err != nil {
			log.Fatal(err.Error())
		}
		// log.Println(id, name, color)
	}

}

// func dbHealthCheck(db *sql.DB) error {
// 	return db.Ping()
// }