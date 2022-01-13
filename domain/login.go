package domain

import "database/sql"

type Login struct {
	UserName   string         `db:"user_name"`
	CustomerId sql.NullString `db:"customer_id"`
	Accounts   sql.NullString `db:"accounts"`
	Role       string         `db:"role"`
}
