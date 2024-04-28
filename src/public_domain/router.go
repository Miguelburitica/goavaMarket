package public_domain

import (
	"net/http"

	"github.com/Miguelburitica/goavaMarket/src/public_domain/controllers"
	"github.com/Miguelburitica/goavaMarket/src/transversal_domain"
)

func indexRoute() transversal_domain.Route {
	handler := make(map[string]http.HandlerFunc)

	handler["GET"] = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Welcome to the Goava Market API", "code": 200}`))
	}

	return transversal_domain.Route{
		Name:             "index",
		AvailableMethods: []string{"GET"},
		Pattern:          "/",
		HandlerFunctions: handler,
	}
}

func GetPublicDomainRoutes() transversal_domain.Routes {
	return transversal_domain.Routes{
		indexRoute(),
		controllers.UserRoutes(),
	}
}
