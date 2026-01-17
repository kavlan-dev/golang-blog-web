package middleware

import (
	"net/http"

	"golang-blog-web/internal/services"
)

func AuthMiddleware(service *services.Service, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name, pass, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		_, err := service.AuthenticateUser(name, pass)
		if err != nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
