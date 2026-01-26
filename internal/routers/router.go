package routers

import (
	"go-blog-web/internal/middleware"
	"go-blog-web/internal/models"
	"net/http"
)

type handlerInterface interface {
	postHandler
	userHandler
	HealthCheck(w http.ResponseWriter, r *http.Request)
}

type userService interface {
	AuthenticateUser(username, password string) (*models.User, error)
}

type postHandler interface {
	CreatePost(w http.ResponseWriter, r *http.Request)
	Posts(w http.ResponseWriter, r *http.Request)
	PostById(w http.ResponseWriter, r *http.Request)
	PostByTitle(w http.ResponseWriter, r *http.Request)
	UpdatePost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
}

type userHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
}

func SetupRoutes(handler handlerInterface, service userService) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.HealthCheck)
	mux.HandleFunc("GET /api/posts", handler.Posts)
	mux.HandleFunc("GET /api/posts/{id}", handler.PostById)
	mux.HandleFunc("GET /api/posts/title/{title}", handler.PostByTitle)
	mux.HandleFunc("POST /api/auth/register", handler.CreateUser)

	mux.HandleFunc("POST /api/posts", middleware.AuthMiddleware(service, handler.CreatePost))

	mux.HandleFunc("PUT /api/posts/{id}", middleware.AuthAdminMiddleware(service, handler.UpdatePost))
	mux.HandleFunc("DELETE /api/posts/{id}", middleware.AuthAdminMiddleware(service, handler.DeletePost))
	mux.HandleFunc("PUT /api/users/{id}", middleware.AuthAdminMiddleware(service, handler.UpdateUser))

	return mux
}
