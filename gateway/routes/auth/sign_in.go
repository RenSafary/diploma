package auth

import (
	"diploma/auth-service/utils"
	grpc_auth "diploma/gateway/grpc/auth"
	"html/template"
	"log"
	"net/http"
)

type Client struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignInForm(w http.ResponseWriter, r *http.Request) {
	// Checking jwt token
	if token, err := r.Cookie("user"); err == nil {
		_, _, _, err = utils.ParseToken(token.Value)
		if err == nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
	}

	// Parse if there is one
	tmpl, err := template.ParseFiles("./templates/auth/sign-in.html")
	if err != nil {
		log.Println("Couldn't parse HTML 'sign-in': ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, nil)
}

func SignInWS(w http.ResponseWriter, r *http.Request) {
	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Err with sign in ws: ", err)
		return
	}
	defer ws.Close()

	for {
		var client Client
		if err := ws.ReadJSON(&client); err != nil {
			log.Println("Error reading JSON:", err)
			return
		}

		status, token := grpc_auth.GRPC_SignIn(client.Username, client.Password)
		resp := map[string]interface{}{
			"status": status,
			"token":  token,
		}

		if err := ws.WriteJSON(resp); err != nil {
			log.Println("Error writing JSON to ws:", err)
			return
		}
	}
}
