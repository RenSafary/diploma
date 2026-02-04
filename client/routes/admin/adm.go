package admin

import (
	"html/template"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")

		tmpl, err := template.ParseFiles("./templates/admin/auth/sign-in.html")
		if err != nil {
			http.Error(w, "Unable to parse admin sign-in.html form", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	} else {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		if username != "admin" || password != "admin" {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		http.Redirect(w, r, "/adm", http.StatusSeeOther)
	}
}
