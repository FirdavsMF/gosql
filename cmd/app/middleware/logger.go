package middleware

import (
	"log"
	"net/http"
)

//Logger - logging requests.
func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("START Logger: %s %s", r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
		log.Printf("FINISH Logger: %s %s", r.Method, r.URL.Path)
	})
}
