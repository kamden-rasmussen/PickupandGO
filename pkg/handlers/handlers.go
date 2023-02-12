package handlers

import (
	"github.com/gorilla/mux"
	"gopkg.in/robfig/cron.v2"
)

func InitializeHandlers() *mux.Router {

	cronService := cron.New()

	
	cronService.AddFunc("@every 5s", func() {
		data.printHealth()
	})

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/healthcheck", GetHealthCheck()).Methods("GET")

	return router
}