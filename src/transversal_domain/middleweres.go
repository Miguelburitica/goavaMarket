package transversal_domain

import "net/http"

func HandlerWrapper(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// set headers for view json content type
		w.Header().Set("Content-Type", "application/json")

		handler(w, r)
	}
}
