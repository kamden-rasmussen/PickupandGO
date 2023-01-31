package handlers

import (
	"github.com/gorilla/mux"
)

func InitializeHandlers(){
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/healthcheck", GetHealthCheck()).Methods("GET")

	return router
}