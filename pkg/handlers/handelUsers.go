package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kamden-Rasmussen/PickupandGO/pkg/data"
)

func AddUserHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user data.User
		_ = json.NewDecoder(r.Body).Decode(&user)
		fmt.Println(user)
		data.AddUser(user)
	}
}
