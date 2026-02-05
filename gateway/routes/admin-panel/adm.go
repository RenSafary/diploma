package admin

import (
	"diploma/auth-service/utils"
	"html/template"
	"net/http"
)

func AdminPanel(w http.ResponseWriter, r *http.Request) {
	// Checking jwt token
	if token, err := r.Cookie("user"); err == nil {
		_, _, admin, err := utils.ParseToken(token.Value)
		if err == nil {
			if !admin {
				http.Error(w, "Access forbidden. You're not an admin", http.StatusForbidden)
				return
			}
			http.Redirect(w, r, "/adm", http.StatusTemporaryRedirect)
			return
		}
	} else {
		http.Error(w, "Access forbidden. You're not an admin", http.StatusForbidden)
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
