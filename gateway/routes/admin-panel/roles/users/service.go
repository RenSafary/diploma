package roles

import (
	grpc_users "diploma/gateway/grpc/users"
	userspb "diploma/proto/users"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Users []*userspb.User
}

func MakeEmployee() {

}

func RemoveEmployee() {

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// get users' list
		users, err := grpc_users.GRPC_Get_All_Users()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// parsing
		tmpl, err := template.ParseFiles("./templates/admin/roles/users/GetAllUsers.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := PageData{
			Users: users,
		}

		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPost:
		id := r.FormValue("id")
		action := r.FormValue("action")
		log.Println(id)

		switch action {
		case "make_admin":

			return
		case "remove_admin":
			return
		case "delete":
			return
		case "history":
			return
		}
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}
