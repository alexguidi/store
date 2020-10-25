package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectWithPostgres() *sql.DB {
	conn := "user=postgres dbname=alura_store password=example host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err.Error())
	}

	return db
}
