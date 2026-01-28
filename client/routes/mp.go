package routes

import (
	"log"
	"net/http"
	"text/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/main.html")
	if err != nil {
		log.Println("Couldn't parse HTML 'main'... ", err)
		return
	}

	tmpl.Execute(w, nil)
}
