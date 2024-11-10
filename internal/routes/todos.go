package routes

import (
	"net/http"

	"github.com/gorvk/todoapp/internal/controllers"
)

func init() {
	http.HandleFunc("GET /api/todo", controllers.GetAllTodo)
}
