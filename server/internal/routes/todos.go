package routes

import (
	"net/http"

	controllers "github.com/gorvk/todoapp/internal/controllers/todo"
)

func init() {
	http.HandleFunc("GET /api/todo", controllers.GetAll)
	http.HandleFunc("GET /api/todo/{id}", controllers.GetById)
	http.HandleFunc("POST /api/todo", controllers.Create)
	http.HandleFunc("PUT /api/todo", controllers.Update)
	http.HandleFunc("DELETE /api/todo/{id}", controllers.Delete)
}
