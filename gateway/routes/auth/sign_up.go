package auth

import (
	"diploma/auth-service/utils"
	grpc_auth "diploma/gateway/grpc/users"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func SignUpForm(w http.ResponseWriter, r *http.Request) {
	// Checking jwt token
	if token, err := r.Cookie("user"); err == nil {
		_, _, _, err = utils.ParseToken(token.Value)
		if err == nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
	}

	// Parse template if there is one
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
	// Checking jwt token
	if token, err := r.Cookie("user"); err == nil {
		_, _, _, err = utils.ParseToken(token.Value)
		if err == nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
	}

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

	status, userId := grpc_auth.GRPC_SignUp(username, password, firstname, lastname, email, sex, age)
	if !status {
		http.Error(w, "Couldn't sign up", http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"status": status,
		"userId": userId,
	}

	json.NewEncoder(w).Encode(response)
}
