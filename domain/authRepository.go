package domain

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/joepaul-martin/banking-auth/errs"
)

type AuthRepository interface {
	FindBy(string, string) (*Login, *errs.AppError)
}

type AuthRepositoryDb struct {
	client *sqlx.DB
}

var _ AuthRepository = (*AuthRepositoryDb)(nil)

func (d AuthRepositoryDb) FindBy(userName string, password string) (*Login, *errs.AppError) {
	var login Login
	sqlFindBy := ""
	err := d.client.Get(&login, sqlFindBy, userName, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NotFoundError("User not found. Please provide correct username and password")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error while trying to fetch user details")
		}
	}
	return &login, nil
}
