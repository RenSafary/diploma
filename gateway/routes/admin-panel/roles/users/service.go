package roles

import (
	grpc_admin "diploma/gateway/grpc/admin"
	grpc_users "diploma/gateway/grpc/users"
	userspb "diploma/proto/users"
	"html/template"
	"net/http"
	"strconv"
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

		switch action {
		case "make_admin":
			u_id, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, "Couldn't convert data", http.StatusInternalServerError)
				return
			}
			status, response := grpc_admin.GRPC_MAKE_Admin(u_id)
			if !status {
				http.Error(w, response, http.StatusBadRequest)
				return
			}
			http.Redirect(w, r, "/adm/users", http.StatusSeeOther)
			return
		case "remove_admin":
			u_id, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, "Couldn't convert data", http.StatusInternalServerError)
				return
			}
			status, response := grpc_admin.GRPC_DELETE_Admin(u_id)
			if !status {
				http.Error(w, response, http.StatusBadRequest)
				return
			}
			http.Redirect(w, r, "/adm/users", http.StatusSeeOther)
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
