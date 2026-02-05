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

func MakeAdminWS(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Err with MakeAdminWS: ", err)
		return
	}
	defer ws.Close()

	for {
		var user User
		err := ws.ReadJSON(&user)
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Println("WebSocket closed:", err)
			} else {
				log.Println("Error reading JSON:", err)
			}
			break
		}

		userID, err := strconv.Atoi(user.Id)
		if err != nil {
			log.Println("Error converting string to int:", err)
			continue
		}

		status, msg := grpc_admin.GRPC_Make_Admin(userID)
		resp := map[string]interface{}{
			"status":   status,
			"response": msg,
		}

		if err := ws.WriteJSON(resp); err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Println("WebSocket closed while writing:", err)
			} else {
				log.Println("Error writing JSON to ws:", err)
			}
			break
		}
	}
}

func RemoveAdmin() {

}
