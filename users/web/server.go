package web

import "net/http"

func SetupRouters() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/users/getall", getAllUsers)
	mux.HandleFunc("/users/youngest", findYoungestUser)
	mux.HandleFunc("/users/save", saveUsers)
	mux.HandleFunc("/users.load", saveUser)
	return mux
}

func StartServer() {
	mux := SetupRouters()
	http.ListenAndServe(":8081", mux)
}

// handleUsers обрабатывает разные HTTP-методы для /users
func saveUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		saveUser(w, r)
	}
	if r.Method == http.MethodPatch {
		loadUser(w, r)
	}
}
