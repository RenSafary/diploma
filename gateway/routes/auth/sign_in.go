package auth

import (
	"diploma/auth-service/utils"
	grpc_auth "diploma/gateway/grpc/auth"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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
	ws, err := upgrader.Upgrade(w, r, nil)
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

func PutToken(w http.ResponseWriter, r *http.Request) {
	// Checking jwt token
	if token, err := r.Cookie("user"); err == nil {
		_, _, _, err = utils.ParseToken(token.Value)
		if err == nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
	}

	var body struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "user",
		Value:    body.Token,
		Path:     "/",
		MaxAge:   86400, // 24 hours
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusOK)
}
