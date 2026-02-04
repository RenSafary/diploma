package main

import (
	"diploma/gateway/routes"
	"diploma/gateway/routes/admin"
	"diploma/gateway/routes/auth"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Static
	static := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static))

	// Main
	r.HandleFunc("/", routes.MainPage).Methods("GET")

	// Authorization
	r.HandleFunc("/sign-in", auth.SignInForm).Methods("GET")
	r.HandleFunc("/sign-in-ws", auth.SignInWS)
	r.HandleFunc("/sign-in/put-token", auth.PutToken).Methods("POST")

	// Registration
	r.HandleFunc("/sign-up", auth.SignUpForm).Methods("GET")
	r.HandleFunc("/sign-up", auth.SignUpPost).Methods("POST")

	// Admin panel
	r.HandleFunc("/adm", admin.AdminPanel).Methods("GET", "POST")
	r.HandleFunc("/adm/sign-in", admin.SignIn).Methods("GET", "POST")

	log.Println("Server is started on port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Couldn't start the server:", err)
	}
}
