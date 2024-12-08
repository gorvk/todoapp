package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorvk/todoapp/internal/initializers"
	"github.com/gorvk/todoapp/internal/types"
)

func Callback(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	auth := initializers.GetAuthInstance()

	token, err := auth.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	idToken, err := auth.VerifyIDToken(r.Context(), token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var profile types.Profile
	if err := idToken.Claims(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    token.AccessToken,
		Expires:  time.Now().Add(time.Hour * 168),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, cookie)

	response, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
