package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Miguelburitica/goavaMarket/src"
	"github.com/Miguelburitica/goavaMarket/src/public_domain/repository"
	"github.com/Miguelburitica/goavaMarket/src/transversal_domain"
	"github.com/Miguelburitica/goavaMarket/src/transversal_domain/database"
)

func main() {
	viewRoutes := src.GetViewRoutes()
	// middlewares := []transversal_domain.Middleware{transversal_domain.LoggingMiddleware(), transversal_domain.AuthenticationMiddleware()}

	for _, route := range viewRoutes {
		// var handler http.Handler = http.HandlerFunc(finalHandler)
		// for _, middleware := range middlewares {
		// 	handler = middleware(handler)
		// }

		http.Handle(route.Pattern, transversal_domain.HandlerWrapper(route.MainHandler()))
	}

	port := os.Getenv("PORT")
	databasePrimaryUrl := os.Getenv("DATABASE_PRIMARY_URL")
	databaseAuthToken := os.Getenv("DATABASE_AUTH_TOKEN")

	if os.Getenv("PORT") == "" {
		port = "8080"
	}

	repo, err := database.NewPostgresRepository(databasePrimaryUrl + "?authToken=" + databaseAuthToken)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer repo.Disconnect()

	repository.SetRepository(repo)

	fmt.Printf("*---------------------*" + "\n")
	fmt.Printf("Server running in %s port\n", port)
	fmt.Printf("Now, you can open http://localhost:%s in your favorite browser\n", port)
	fmt.Printf("*---------------------*" + "\n")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
