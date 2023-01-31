package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Kamden-Rasmussen/PickupandGO/pkg/handlers"
)

func main(){
	_ = context.Background()

	fmt.Println("Starting Server on Port 8018")
	
	router := handlers.InitializeHandlers()
	log.Fatal(http.ListenAndServe(":8018", router))
}