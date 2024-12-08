package controllers

import (
	"net/http"
	"time"

	"github.com/gorvk/todoapp/internal/initializers"
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

	var profile map[string]interface{}
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

	_, err = w.Write([]byte("success"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
