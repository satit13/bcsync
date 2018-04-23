package auth

import (
	//"context"
	"net/http"
	//"fmt"
)

// MakeMiddleware creates new auth middleware
func MakeMiddleware(s Service) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// get token from header
			h.ServeHTTP(w, r)
		})
	}
}
