package main

import (
	"diploma/client/routes"
	"diploma/client/routes/auth"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	static := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static))

	// main page
	r.HandleFunc("/", routes.MainPage).Methods("GET")

	// authentication
	// sign in
	r.HandleFunc("/sign-in", auth.SignInForm).Methods("GET")
	r.HandleFunc("/sign-in", auth.SignInPost).Methods("POST")
	//sign up
	r.HandleFunc("sign-up", auth.SignUpForm).Methods("GET")
	r.HandleFunc("/sign-up", auth.SignUpPost).Methods("POST")

	log.Println("Server is started on port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Couldn't start the server: ", err)
	}
}
