package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Miguelburitica/goavaMarket/src/public_domain/models"
	"github.com/Miguelburitica/goavaMarket/src/public_domain/services"
	"github.com/Miguelburitica/goavaMarket/src/transversal_domain"
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// get from de path the id of the user as required and as optional page, offset and patern (search)

	// get the id of the user
	id := r.URL.Query().Get("id")

	// validate if the id is empty
	if id != "" {

		// get the user from the repository
		user, err := services.GetUser(r.Context(), id)

		// validate if the user exists
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error getting the user"))
			return
		}

		// validate if the user exists
		if user.ID == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("User not found"))
			return
		}

		response, err := json.Marshal(user)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error parsing the user"))
			return
		}

		w.Write(response)
		w.WriteHeader(http.StatusOK)
	}
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

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// get the page, offset and pattern
	page := transversal_domain.GetQueryParamWithDefault(r, "page", "1")
	offset := transversal_domain.GetQueryParamWithDefault(r, "offset", "10")
	pattern := transversal_domain.GetQueryParamWithDefault(r, "pattern", "*")

	props := models.GetUsersRequest{
		Page:    page,
		Offset:  offset,
		Pattern: pattern,
	}

	// get the users from the service
	users, err := services.GetUsers(r.Context(), props)

	if err != nil {
		log.Printf("Error getting the users: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(users)

	if err != nil {
		log.Printf("Error parsing the users: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
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

func UsersRoutes() transversal_domain.Route {
	handlers := make(map[string]http.HandlerFunc)

	handlers["GET"] = getUsersHandler

	return transversal_domain.Route{
		Name:             "users",
		AvailableMethods: []string{"GET"},
		Pattern:          "/users",
		HandlerFunctions: handlers,
	}
}
