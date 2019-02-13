package model

import(
	"github.com/criszelaya24/Go/db"
)

func NewUser(name string, u_name string, email string, password string) string {
	var id string
	q := `INSERT INTO users (name, u_name, email, password) 
		VALUES ($1, $2, $3, $4)`

	err := postgres.DbClient.MustExec(q, name, u_name, email, password).Scan(&id)
	if err != nil {
		return " "
	}
	return id
}