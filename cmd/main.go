package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Kamden-Rasmussen/PickupandGO/pkg/cron"
	"github.com/Kamden-Rasmussen/PickupandGO/pkg/data"
	"github.com/Kamden-Rasmussen/PickupandGO/pkg/handlers"
	"github.com/Kamden-Rasmussen/PickupandGO/pkg/mydatabase"
	"github.com/joho/godotenv"
)

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
	err = mydatabase.Init()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database")
	}

	// create tables
	// SetUpTables(db) // only run once

	// setup airports
	mydatabase.SetUpAirportCodes()

	// start cron jobs
	startCron(mydatabase.MyDB)
	
	// start server
	log.Println("Starting Server on Port 8018")
	
	allHealthChecks()
	router := handlers.InitializeHandlers()
	log.Fatal(http.ListenAndServe(":8018", router))
}

func startCron(db *sql.DB) {
	cron := cron.NewCron()
	log.Println("Starting Cron Jobs")
	cron.AddFunc("@every 30m", allHealthChecks)
	// cron.AddFunc("@every 12h", data.GetData)
	// cron.AddFunc("0 0 1 * * *", data.PrintLetsGo)
	cron.AddFunc("0 0 08 * * *", data.PrintLetsGo)

	cron.Start()
}

func allHealthChecks() {
	// log.Println("Application is running and healthy")
	err := mydatabase.DbHealthCheck()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database and connection are healthy")
	}
	
}