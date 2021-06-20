package middleware

import (
	"net/http"
)

//CheckHeader - check headers in request.
func CheckHeader(header, value string) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if value != r.Header.Get(header) {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			handler.ServeHTTP(w, r)
		})
	}
}