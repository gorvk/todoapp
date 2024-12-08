package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	models "github.com/gorvk/todoapp/internal/models/todo"
	"github.com/gorvk/todoapp/internal/types"
	"github.com/gorvk/todoapp/internal/utils"
)

func Update(w http.ResponseWriter, r *http.Request) {
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var input = types.Todo{}
	err = json.Unmarshal(d, &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = models.Update(input)
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
