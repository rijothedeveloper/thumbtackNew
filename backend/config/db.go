package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const host     = "host.docker.internal"
const port     = 5432

var DB *sql.DB

func ConnectDB() {
	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		user = "rijo1"
	}
	password, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		password = "rijopass1"
	}
	dbname, ok := os.LookupEnv("DB_NAME")
	if !ok {
		dbname = "rijo1"
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		  "password=%s dbname=%s sslmode=disable",
		  host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
	// defer db.Close()
	DB = db
}