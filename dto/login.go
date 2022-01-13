package dto

import (
	"github.com/joepaul-martin/banking-auth/errs"
)

type Login struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (l Login) Validate() *errs.AppError {
	if l.UserName == "" {
		return errs.NewValidationError("Username must not be empty")
	}
	return nil
}
