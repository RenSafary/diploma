package main

import (
	"diploma/client/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", routes.MainPage).Methods("GET")

	log.Println("Server is started on port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Couldn't start the server: ", err)
	}
}
