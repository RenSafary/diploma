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

	r.HandleFunc("/", routes.MainPage).Methods("GET")

	// authentication
	r.HandleFunc("/sign-in", auth.SignInForm).Methods("GET")
	r.HandleFunc("/sign-in", auth.SignInPost).Methods("POST")

	log.Println("Server is started on port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Couldn't start the server: ", err)
	}
}
