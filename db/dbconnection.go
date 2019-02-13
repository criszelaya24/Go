package postgres

import(
	"github.com/jmoiron/sqlx"
	"fmt"
	_ "github.com/lib/pq"
)

//here initialise postgres db connection

var DbClient *sqlx.DB

func ConnectDB(dbuser string, dbpass string, dbhost string, dbport string, dbname string) error {
	var err error
	DbClient, err = sqlx.Connect(
		"postgres",
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbuser, dbpass, dbhost, dbport, dbname))
	if err != nil {
		return err
	}
	return nil
}