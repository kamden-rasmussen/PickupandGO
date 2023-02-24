package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Kamden-Rasmussen/PickupandGO/pkg/handlers"
	"github.com/Kamden-Rasmussen/PickupandGO/src/cron"
	"github.com/Kamden-Rasmussen/PickupandGO/src/data"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main(){
	_ = context.Background()

	// create log file
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)


	// load env variables
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load env. Err: %s", err)
	}

	// connect to a database
	db, err := Init()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database")
	}
	defer db.Close()
	GetFunTable(db)

	// start cron jobs
	startCron(db)

	// start server
	log.Println("Starting Server on Port 8018")

	router := handlers.InitializeHandlers()
	log.Fatal(http.ListenAndServe(":8018", router))
}

func startCron(db *sql.DB) {
	cron := cron.NewCron()
	log.Println("Starting Cron Jobs")
	cron.AddFunc("@every 5m", data.PrintHealth)
	// cron.AddFunc("@every 5m", allHealthChecks)
	// cron.AddFunc("@every 12h", data.GetData)
	// cron.AddFunc("0 0 1 * * *", data.PrintLetsGo)
	cron.AddFunc("0 0 13 * * *", data.PrintLetsGo)

	cron.Start()
}

func allHealthChecks() {
	err := dbHealthCheck(db)
	if err != nil {
		log.Fatal(err)
	}
	
}

func Init() (*sql.DB, error) {
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

func dbHealthCheck(db *sql.DB) error {
	return db.Ping()
}