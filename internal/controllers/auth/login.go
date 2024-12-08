package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorvk/todoapp/internal/initializers"
	"github.com/gorvk/todoapp/internal/utils"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	auth := initializers.GetAuthInstance()
	state, err := utils.GenerateRandomState()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	url := auth.AuthCodeURL(state)
	body := map[string]string{
		"url": url,
	}
	response, err := json.Marshal(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
