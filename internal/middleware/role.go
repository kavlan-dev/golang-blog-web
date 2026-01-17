package middleware

import (
	"go-blog-web/internal/services"
	"net/http"
)

func AuthAdminMiddleware(service *services.Service, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name, pass, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		user, err := service.AuthenticateUser(name, pass)
		if err != nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if user.Role != "admin" {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Не достаточно прав", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}
}
