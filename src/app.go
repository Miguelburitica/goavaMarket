package src

import (
	"github.com/Miguelburitica/goavaMarket/src/public_domain"
	"github.com/Miguelburitica/goavaMarket/src/transversal_domain"
)

func GetViewRoutes() transversal_domain.Routes {
	routes := transversal_domain.GetGroupedRoutes(
		public_domain.GetPublicDomainRoutes()...,
	)

	// Add the public domain routes
	return routes
}
