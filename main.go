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

	for _, route := range viewRoutes {
		http.Handle(route.Pattern, transversal_domain.HandlerWrapper(route.MainHandler()))
	}

	envPort := os.Getenv("PORT")

	if os.Getenv("PORT") == "" {
		envPort = "8080"
	}

	port := envPort

	fmt.Printf("*---------------------*" + "\n")
	fmt.Printf("Server running in %s port\n", port)
	// console.log(`Now, you can open http://localhost:${port} in your favorite browser `);
	fmt.Printf("Now, you can open http://localhost:%s in your favorite browser\n", port)
	fmt.Printf("*---------------------*" + "\n")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
