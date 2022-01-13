package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joepaul-martin/banking-auth/service"
)

func Start() {
	router := mux.NewRouter()

	lh := LoginHandler{
		service: service.NewDefaultLoginService(),
	}

	router.HandleFunc("/login", lh.login).Methods(http.MethodPost)

	http.ListenAndServe("localhost:8001", router)
}
