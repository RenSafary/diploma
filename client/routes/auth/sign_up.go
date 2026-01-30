package auth

import (
	grpc_auth "diploma/client/grpc/auth"
	"encoding/json"
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

func SignUpPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	age := r.FormValue("age")
	sex := r.FormValue("sex")

	if sex == "Мужской" {
		sex = "М"
	} else {
		sex = "Ж"
	}

	status, token := grpc_auth.GRPC_SignUp(username, password, firstname, lastname, email, sex, age)
	if !status {
		http.Error(w, "Couldn't sign up", http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"status": status,
		"token":  token,
	}

	json.NewEncoder(w).Encode(response)
}
