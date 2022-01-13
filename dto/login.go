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
	if len(l.UserName) < 2 || len(l.UserName) > 40 {
		return errs.NewValidationError("Username should contains between 2 and 40 number of characters")
	}
	return nil
}
