package controllers

import (
	"net/http"

	models "github.com/gorvk/todoapp/internal/models/todo"
	"github.com/gorvk/todoapp/internal/utils"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := models.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := utils.ConstructResponse(true, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
