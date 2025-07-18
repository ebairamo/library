package web

import "net/http"

func SetupRouters() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/users/getall", getAllUsers)
	return mux
}

func StartServer() {
	mux := SetupRouters()
	http.ListenAndServe(":8081", mux)
}
