package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Miguelburitica/goavaMarket/src"
	"github.com/Miguelburitica/goavaMarket/src/transversal_domain"
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

	envPort := os.Getenv("PORT")

	if os.Getenv("PORT") == "" {
		envPort = "8080"
	}

	port := envPort

	fmt.Printf("*---------------------*" + "\n")
	fmt.Printf("Server running in %s port\n", port)
	fmt.Printf("Now, you can open http://localhost:%s in your favorite browser\n", port)
	fmt.Printf("*---------------------*" + "\n")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
