package handlers

import (
	"github.com/gorilla/mux"
)

func InitializeHandlers() *mux.Router {



	router := mux.NewRouter()
	router.HandleFunc("/api/v1/healthcheck", GetHealthCheck()).Methods("GET")

	router.HandleFunc("/api/v1/getall", GetDataEndpoint()).Methods("GET")

	return router
}