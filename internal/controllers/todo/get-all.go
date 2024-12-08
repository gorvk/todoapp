package controllers

import (
	"encoding/json"
	"net/http"

	models "github.com/gorvk/todoapp/internal/models/todo"
	"github.com/gorvk/todoapp/internal/types"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	rows, err := models.GetAll()
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
