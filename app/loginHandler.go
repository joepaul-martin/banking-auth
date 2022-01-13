package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joepaul-martin/banking-auth/dto"
)

type LoginHandler struct {
}

func (ah *LoginHandler) login(w http.ResponseWriter, req *http.Request) {
	var loginRequest dto.Login
	json.NewDecoder(req.Body).Decode(&loginRequest)
	fmt.Printf("UserName:%s , Password:%s", loginRequest.UserName, loginRequest.Password)
}
