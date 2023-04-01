package handlers

import (
	"net/http"

	"github.com/Kamden-Rasmussen/PickupandGO/pkg/calcs"
	"github.com/Kamden-Rasmussen/PickupandGO/pkg/data"
	"github.com/Kamden-Rasmussen/PickupandGO/pkg/email"
	"github.com/gorilla/mux"
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

func ForceCalcs() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		calcs.BeginCalc()
	}
}

func ForceEmail() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// get email from url
		emailAddr := mux.Vars(r)["email"]
		email.SendTestEmail(emailAddr, "Hello World")
	}
}

func ForceCalcsWithEmail() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// get email from url
		calcs.ForceBool = true
		calcs.BeginCalc()
		calcs.ForceBool = false
	}
}