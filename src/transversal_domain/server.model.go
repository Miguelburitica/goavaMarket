package transversal_domain

import (
	"fmt"
	"net/http"
)

// Route struct
type Route struct {
	Name             string
	AvailableMethods []string
	Pattern          string
	HandlerFunctions map[string]http.HandlerFunc
}

func (route *Route) MainHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received")

		handler := route.HandlerFunctions[r.Method]
		if (handler == nil) {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		handler(w, r)
	}
}

// Routes type
type Routes []Route

// Request Resource struct
type GetRequestResource struct {
	Id      string `json:"id"`
	Page    int    `json:"page"`
	Offset  int    `json:"offset"`
	Pattern string `json:"pattern"`
}

type PostRequestSeller struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	SignInMethod string `json:"sign_in_method"`
	Password     string `json:"password"`
}

// Response Resource struct
