package domain

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joepaul-martin/banking-auth/errs"
)

const signingKey = "secretKeysdfsdfsdfsdfsdf"

type Login struct {
	UserName   string         `db:"username"`
	CustomerId sql.NullString `db:"customer_id"`
	Accounts   sql.NullString `db:"accounts"`
	Role       string         `db:"role"`
}

func (l Login) GenerateToken() (*string, *errs.AppError) {
	var claims jwt.Claims
	if l.Accounts.Valid && l.CustomerId.Valid {
		claims = l.claimsForUser()
	} else {
		claims = l.claimsForAdmin()
	}
	generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := generatedToken.SignedString([]byte(signingKey))
	if err != nil {
		return nil, errs.NewUnexpectedError(fmt.Sprintf("Error while generating token : %s", err.Error()))
	}
	return &signedToken, nil
}

func (l Login) claimsForUser() jwt.MapClaims {
	return jwt.MapClaims{
		"customer_id": l.CustomerId,
		"username":    l.UserName,
		"accounts":    strings.Split(l.Accounts.String, ","),
		"role":        l.Role,
		"exp":         time.Now().Add(time.Minute * 15).Unix(),
	}
}

func (l Login) claimsForAdmin() jwt.MapClaims {
	return jwt.MapClaims{
		"username": l.UserName,
		"role":     l.Role,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	}
}
