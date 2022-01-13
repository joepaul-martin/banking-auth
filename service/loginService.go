package service

import (
	"github.com/joepaul-martin/banking-auth/dto"
	"github.com/joepaul-martin/banking-auth/errs"
)

type LoginService interface {
	Login(dto.Login) *errs.AppError
}

type DefaultLoginService struct {
}

func (ls DefaultLoginService) Login(loginRequest dto.Login) *errs.AppError {
	err := loginRequest.Validate()
	if err != nil {
		return err
	}
	return nil
}

func NewDefaultLoginService() DefaultLoginService {
	return DefaultLoginService{}
}
