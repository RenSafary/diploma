package routes

import (
	"log"
	"net/http"
	"text/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	log.Println("Got / request")

	token, err := r.Cookie("user")
	if err != nil {
		log.Println("No JWT cookie:", err)
		http.Redirect(w, r, "/sign-in", http.StatusFound)
		return
	}

	log.Println("JWT token:", token.Value)

	tmpl, err := template.ParseFiles("./templates/main.html")
	if err != nil {
		log.Println("Couldn't parse HTML 'main':", err)
		return
	}

	tmpl.Execute(w, nil)
}
