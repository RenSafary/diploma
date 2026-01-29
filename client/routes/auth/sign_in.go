package auth

import (
	grpc_auth "diploma/client/grpc/auth"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func SignInForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/auth/sign-in.html")
	if err != nil {
		log.Println("Couldn't parse HTML 'main'... ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, nil)
}

func SignInPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	status, token := grpc_auth.GRPC_SignIn(username, password)
	if !status {
		http.Error(w, "Couldn't sign in", http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"status": status,
		"token":  token,
	}

	json.NewEncoder(w).Encode(response)
}
