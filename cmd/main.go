package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Kamden-Rasmussen/PickupandGO/pkg/handlers"
	"github.com/Kamden-Rasmussen/PickupandGO/src/cron"
	"github.com/Kamden-Rasmussen/PickupandGO/src/data"
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
	db, err := ConnectToDatabase()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database")
	}
	defer db.Close()
	GetFunTable(db)

	// start cron jobs
	startCron()

	// start server
	log.Println("Starting Server on Port 8018")

	router := handlers.InitializeHandlers()
	log.Fatal(http.ListenAndServe(":8018", router))
}

func startCron() {
	cron := cron.NewCron()
	cron.AddFunc("@every 5m", data.PrintHealth)
	// cron.AddFunc("@every 12h", data.GetData)
	// cron.AddFunc("0 0 1 * * *", data.PrintLetsGo)
	cron.AddFunc("0 0 13 * * *", data.PrintLetsGo)

	cron.Start()
}

// func allHealthChecks(db *sql.DB) {
// 	err := dbHealthCheck(db)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
	
// }