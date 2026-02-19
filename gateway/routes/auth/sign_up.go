package auth

import (
	"diploma/auth-service/utils"
	grpc_auth "diploma/gateway/grpc/auth"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Age       string `json:"age"`
	Sex       string `json:"sex"`
}

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

func SignUpWS(w http.ResponseWriter, r *http.Request) {
	// Checking jwt token
	if token, err := r.Cookie("user"); err == nil {
		_, _, _, err = utils.ParseToken(token.Value)
		if err == nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
	}

	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Err with sign up ws: ", err)
		return
	}
	defer ws.Close()

	for {
		var user User
		if err := ws.ReadJSON(&user); err != nil {
			log.Println("Error reading JSON: ", err)
			return
		}

		status, token := grpc_auth.GRPC_SignUp(user.Username, user.Password, user.Firstname, user.Lastname, user.Email, user.Sex, user.Age)
		resp := map[string]interface{}{
			"status": status,
		}
		if status {
			resp["token"] = token
		} else {
			resp["error"] = "Coulnd't sign up"
		}

		if err := ws.WriteJSON(resp); err != nil {
			log.Println("Error writing JSON to ws:", err)
			return
		}
	}
}
