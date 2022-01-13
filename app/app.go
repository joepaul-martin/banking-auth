package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	lh := LoginHandler{}

	router.HandleFunc("/login", lh.login).Methods(http.MethodPost)

	http.ListenAndServe("localhost:8001", router)
}
