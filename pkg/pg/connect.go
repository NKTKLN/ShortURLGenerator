package pg

import (
	"fmt"
	"os"

	"github.com/NKTKLN/ShortURLGenerator/pkg/others"
	"github.com/joho/godotenv"
	"github.com/jmoiron/sqlx"
)

// Connecting to the db
func DbConnect() *sqlx.DB {
    others.ErrorChecking(godotenv.Load("env/pg.env"))
	user, password, dbname :=
        os.Getenv("POSTGRES_USER"),
        os.Getenv("POSTGRES_PASSWORD"),
        os.Getenv("POSTGRES_DB")
	psqlconn := fmt.Sprintf("host=pg port=5432 user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	db, err := sqlx.Connect("postgres", psqlconn) 
	others.ErrorChecking(err)
    return db
}
