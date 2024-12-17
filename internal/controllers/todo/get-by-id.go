package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorvk/todoapp/internal/initializers"
	models "github.com/gorvk/todoapp/internal/models/todo"
	"github.com/gorvk/todoapp/internal/types"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	auth := initializers.GetAuthInstance()
	userClaim, err := auth.VerifyIDToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	id := r.PathValue("id")
	rows, err := models.GetById(id, userClaim)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todos := []types.Todo{}
	defer rows.Close()
	for rows.Next() {
		row := types.Todo{}
		rows.Scan(
			&row.Id,
			&row.Title,
			&row.IsCompleted,
			&row.UserId,
		)
		todos = append(todos, row)
	}

	response, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
