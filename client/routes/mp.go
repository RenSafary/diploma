package routes

import (
	"log"
	"net/http"
	"text/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("token"); err != nil {
		http.Redirect(w, r, "/sign-in", http.StatusPermanentRedirect)
		return
	}

	tmpl, err := template.ParseFiles("./templates/main.html")
	if err != nil {
		log.Println("Couldn't parse HTML 'main'... ", err)
		return
	}

	tmpl.Execute(w, nil)
}
