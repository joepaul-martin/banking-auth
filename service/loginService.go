package service

import (
	"github.com/joepaul-martin/banking-auth/domain"
	"github.com/joepaul-martin/banking-auth/dto"
	"github.com/joepaul-martin/banking-auth/errs"
)

type LoginService interface {
	Login(dto.Login) (*string, *errs.AppError)
}

type DefaultLoginService struct {
	repo domain.AuthRepositoryDb
}

var _ LoginService = (*DefaultLoginService)(nil)

func (ls DefaultLoginService) Login(loginRequest dto.Login) (*string, *errs.AppError) {
	err := loginRequest.Validate()
	if err != nil {
		return nil, err
	}
	login, err := ls.repo.FindBy(loginRequest.UserName, loginRequest.Password)
	if err != nil {
		return nil, err
	}
	signedToken, err := login.GenerateToken()
	if err != nil {
		return nil, err
	}
	return signedToken, nil
}

func NewDefaultLoginService(repo domain.AuthRepositoryDb) DefaultLoginService {
	return DefaultLoginService{
		repo: repo,
	}
}
