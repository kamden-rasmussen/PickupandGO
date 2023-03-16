package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

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

	log.Println("Starting Server " + time.Now().Format("2006-01-02 15:04:05"))

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
	if os.Getenv("CREATE_TABLES") == "true" {
		mydatabase.SetUpTables(mydatabase.MyDB) // only run once
		os.Setenv("CREATE_TABLES", "false")
	}

	// setup airports
	if os.Getenv("SETUP_AIRPORTS") == "true" {
		mydatabase.SetUpAirportCodes()
		os.Setenv("SETUP_AIRPORTS", "false")
	}

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
	cron.AddFunc("0 0 * 30 * *", allHealthChecks)
	cron.AddFunc("0 0 * 00 * *", allHealthChecks)
	cron.AddFunc("0 0 08 * * *", data.PrintLetsGo)
	cron.AddFunc("0 0 08 * * *", data.ReoccuringTask)

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