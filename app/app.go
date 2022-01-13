package app

import (
	"fmt"
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

	fmt.Printf("Started the application")

	http.ListenAndServe("localhost:8001", router)

}
