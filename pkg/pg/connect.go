package pg

import (
	"fmt"

	"github.com/NKTKLN/ShortURLGenerator/pkg/others"
	"github.com/jmoiron/sqlx"
)

// Connecting to the db
func DbConnect() *sqlx.DB {
    user, password, dbname :=
		others.GetInfromationFromEnv("env/pg.env", "POSTGRES_USER"),
        others.GetInfromationFromEnv("env/pg.env", "POSTGRES_PASSWORD"),
        others.GetInfromationFromEnv("env/pg.env", "POSTGRES_DB")
	psqlconn := fmt.Sprintf("host=pg port=5432 user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	db, err := sqlx.Connect("postgres", psqlconn) 
	others.ErrorChecking(err)
    return db
}
