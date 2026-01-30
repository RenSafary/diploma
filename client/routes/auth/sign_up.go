package auth

import (
	"html/template"
	"log"
	"net/http"
)

func SignUpForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/auth/sign-up.html")
	if err != nil {
		log.Println("Couldn't parse HTML 'sign-up': ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, nil)
}
