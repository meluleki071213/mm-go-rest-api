package http

import (
	"net/http"
	"strings"
)

func JWTAUth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]

		if authHeader == nil {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		//Bearer: token-string
		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		// Validate Token
		if validateToken(authHeaderParts[1]) {
			original(w,r)
		} else {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
	}
}

func validateToken(accessToken string) bool  {
	// Not implemented
	 return true
}
