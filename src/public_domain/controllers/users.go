package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Miguelburitica/goavaMarket/src/transversal_domain"
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	reader := strings.NewReader(r.URL.Query().Encode())

	var request transversal_domain.GetRequestResource
	decoder := json.NewDecoder(reader)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Here is your user..."))
}

func postUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("New user created..."))
}

func putUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User updated..."))
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User deleted..."))
}

func UserRoutes() transversal_domain.Route {
	handlers := make(map[string]http.HandlerFunc)

	handlers["GET"] = getUserHandler
	handlers["POST"] = postUserHandler
	handlers["PUT"] = putUserHandler
	handlers["DELETE"] = deleteUserHandler

	return transversal_domain.Route{
		Name:             "users",
		AvailableMethods: []string{"GET", "POST", "PUT", "DELETE"},
		Pattern:          "/user",
		HandlerFunctions: handlers,
	}
}
