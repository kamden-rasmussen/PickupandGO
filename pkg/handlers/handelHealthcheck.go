package handlers

import (
	"net/http"

	"github.com/Kamden-Rasmussen/PickupandGO/src/data"
)

func GetHealthCheck() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`Server is up and running`))
	}
}

func GetDataEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data.GetData()
	}
}