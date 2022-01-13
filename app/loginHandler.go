package app

import (
	"encoding/json"
	"net/http"

	"github.com/joepaul-martin/banking-auth/dto"
	"github.com/joepaul-martin/banking-auth/service"
)

type LoginHandler struct {
	service service.LoginService
}

func (lh *LoginHandler) login(w http.ResponseWriter, req *http.Request) {
	var loginRequest dto.Login
	err := json.NewDecoder(req.Body).Decode(&loginRequest)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Error while reading content from login request")
	}
	appErr := lh.service.Login(loginRequest)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.Message)
	} else {
		writeResponse(w, http.StatusAccepted, "Login request validated")
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
