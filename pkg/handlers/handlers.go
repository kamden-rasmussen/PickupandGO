package handlers

import (
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func InitializeHandlers() *mux.Router {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load env. Err: %s", err)
	}

	pass := os.Getenv("Endpoint_Password")


	router := mux.NewRouter()
	router.HandleFunc("/api/v1/healthcheck", GetHealthCheck()).Methods("GET")

	router.HandleFunc("/api/v1/initexampledata", GetDataEndpoint()).Methods("GET")
	router.HandleFunc("/api/v1/" + pass, ForceSync()).Methods("GET") // destinations

	// users
	router.HandleFunc("/api/v1/users", AddUserHandler()).Methods("POST")


	return router
}