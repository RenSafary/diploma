package auth

import (
	grpc_auth "diploma/client/grpc/auth"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	username := vars["username"]
	password := vars["password"]

	if username == "" || password == "" {
		log.Println("User's data is empty")
		http.Error(w, "username or password is empty", http.StatusBadRequest)
		return
	}

	status, token := grpc_auth.GRPC_SignIn(username, password)
	if status == false {
		http.Error(w, "Couldn't sign in", http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"status": status,
		"token":  token,
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("JSON encode error:", err)
	}
}
