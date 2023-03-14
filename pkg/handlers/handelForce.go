package handlers

import (
	"net/http"

	"github.com/Kamden-Rasmussen/PickupandGO/pkg/data"
)

func GetDataEndpoint() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data.GetExampleData()
	}
}

func ForceSync() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data.ReoccuringTask()
	}
}