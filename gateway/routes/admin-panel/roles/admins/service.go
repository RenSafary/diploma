package roles

import (
	grpc_admin "diploma/gateway/grpc/admin"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type User struct {
	Id string `json:"id"`
}

func MakeAdmin(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Err with sign in ws: ", err)
		return
	}
	defer ws.Close()

	for {
		var user User
		if err := ws.ReadJSON(&user); err != nil {
			log.Println("Error reading JSON:", err)
			break
		}

		user_id, err := strconv.Atoi(user.Id) // str to int
		if err != nil {
			log.Println("Error converting string to int", err)
			continue
		}

		status, msg := grpc_admin.GRPC_Make_Admin(user_id)
		resp := map[string]interface{}{
			"status":   status,
			"response": msg,
		}

		if err := ws.WriteJSON(resp); err != nil {
			log.Println("Error writing JSON to ws:", err)
			break
		}
	}
}

func RemoveAdmin() {

}
