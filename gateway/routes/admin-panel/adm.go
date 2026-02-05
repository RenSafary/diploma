package admin

import (
	"diploma/auth-service/utils"
	"html/template"
	"net/http"
)

func AdminPanel(w http.ResponseWriter, r *http.Request) {
	// Checking jwt token
	token, err := r.Cookie("user")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	_, _, admin, err := utils.ParseToken(token.Value)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if !admin {
		http.Error(w, "Access is denied", http.StatusForbidden)
		return
	}

	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")

		tmpl, err := template.ParseFiles("./templates/admin/panel.html")
		if err != nil {
			http.Error(w, "Unable to parse admin sign-in.html form", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	}
}
