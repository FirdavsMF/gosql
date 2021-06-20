package middleware

import (
	"log"
	"net/http"
)

//Basic - check login and password by BasicAuth method.
func Basic(auth func(login string, password string) bool) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			name, pass, ok := r.BasicAuth()
			if !ok {
				log.Println("Has no such password and login:", ok)
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			if !auth(name, pass) {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			handler.ServeHTTP(w, r)
		})
	}
}
