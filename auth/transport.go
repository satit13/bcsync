package auth

import (
	//"context"
	"net/http"
	//"fmt"
	"fmt"
	"context"
)

// MakeMiddleware creates new auth middleware
func MakeMiddleware(s Service) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// get token from header
			tokenID := r.Header.Get("X-Access-Token")
			//check with auth
			fmt.Println("auth.transport token : ", tokenID)
			if len(tokenID) == 0 {
				h.ServeHTTP(w, r)
				return
			}

			tk, err := s.GetToken(tokenID)
			if err != nil {
				h.ServeHTTP(w, r)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, keyToken{}, tk)
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}
