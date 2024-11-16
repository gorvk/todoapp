package routes

import (
	"net/http"

	controllers "github.com/gorvk/todoapp/internal/controllers/todo"
)

func init() {
	http.HandleFunc("GET /api/todo", controllers.GetAll)
	http.HandleFunc("POST /api/todo", controllers.Create)
	http.HandleFunc("DELETE /api/todo", controllers.Delete)
}
