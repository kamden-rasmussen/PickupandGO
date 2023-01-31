package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	fmt.Println("Starting Server on Port 8018")
	
	router := handlers.InitializeHandlers()
	log.Fatal(http.ListenAndServe(":8080", router))
}