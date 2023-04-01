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

	// data
	router.HandleFunc("/api/v1/initexampledata", GetDataEndpoint()).Methods("GET")
	router.HandleFunc("/api/v1/" + pass, ForceSync()).Methods("GET") // destinations

	// users
	router.HandleFunc("/api/v1/users", AddUserHandler()).Methods("POST")

	// calcs
	router.HandleFunc("/api/v1/shallwecalc", ForceCalcs()).Methods("GET")

	// email
	router.HandleFunc("/api/v1/shallWeEmail/{email}", ForceEmail()).Methods("POST")

	// calc and email
	router.HandleFunc("/api/v1/shallWeCalcWithEmail", ForceCalcsWithEmail()).Methods("POST")

	return router
}