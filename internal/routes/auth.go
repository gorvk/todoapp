package routes

import (
	"net/http"

	controllers "github.com/gorvk/todoapp/internal/controllers/auth"
)

func init() {
	http.HandleFunc("GET /auth", controllers.Authenticate)
	http.HandleFunc("GET /auth/callback/{code}", controllers.Callback)
}
