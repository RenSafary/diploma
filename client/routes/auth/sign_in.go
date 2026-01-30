package auth

import (
	grpc_auth "diploma/client/grpc/auth"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignInForm(w http.ResponseWriter, r *http.Request) {
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
		// get data from front-end
		var client Client
		err := ws.ReadJSON(&client)
		if err != nil {
			log.Println("Error reading JSON:", err)
			continue
		}

		status, token := grpc_auth.GRPC_SignIn(client.Username, client.Password)
		response := map[string]interface{}{
			"status": status,
			"token":  token,
		}

		// send back to client
		err = ws.WriteJSON(response)
		if err != nil {
			log.Println("Error writing JSON to ws:", err)
			continue
		}
	}
}
