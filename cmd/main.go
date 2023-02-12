package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"

	"github.com/Kamden-Rasmussen/PickupandGO/pkg/handlers"
	"github.com/Kamden-Rasmussen/PickupandGO/src/cron"
	"github.com/Kamden-Rasmussen/PickupandGO/src/data"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main(){
	_ = context.Background()

	// load env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load env. Err: %s", err)
	}

	// connect to a database
	db, err := connectToDatabase()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to Database")
	}
	defer db.Close()
	getFunTable(db)

	// start cron jobs
	startCron()

	fmt.Println("Starting Server on Port 8018")

	router := handlers.InitializeHandlers()
	log.Fatal(http.ListenAndServe(":8018", router))
}

func connectToDatabase() (*sql.DB, error) {
	// connect to a database
	fmt.Println(os.Getenv("DBACCESSUSERNAME"))
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

func startCron() {
	cron := cron.NewCron()
	cron.AddFunc("@every 60s", data.PrintHealth)
	cron.AddFunc("@every 12h", data.GetData)
	cron.Start()
}

func getFunTable(db *sql.DB) {
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
		fmt.Println(id, name, color)
	}

}