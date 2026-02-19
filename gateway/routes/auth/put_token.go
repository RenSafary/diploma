package auth

import (
	"diploma/auth-service/utils"
	"encoding/json"
	"net/http"
)

func PutToken(w http.ResponseWriter, r *http.Request) {
	// Checking jwt token
	if token, err := r.Cookie("user"); err == nil {
		_, _, _, err = utils.ParseToken(token.Value)
		if err == nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
	}

	var body struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "user",
		Value:    body.Token,
		Path:     "/",
		MaxAge:   86400, // 24 hours
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusOK)
}
