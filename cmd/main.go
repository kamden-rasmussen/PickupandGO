package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Kamden-Rasmussen/PickupandGO/pkg/handlers"
	"github.com/Kamden-Rasmussen/PickupandGO/src/cron"
	"github.com/Kamden-Rasmussen/PickupandGO/src/data"
)

func main(){
	_ = context.Background()

	startCron()


	fmt.Println("Starting Server on Port 8018")

	router := handlers.InitializeHandlers()
	log.Fatal(http.ListenAndServe(":8018", router))
}

func startCron() {
	cron := cron.NewCron()
	cron.AddFunc("@every 60s", data.PrintHealth)
	cron.AddFunc("@every 12h", data.GetData)
	cron.Start()
}