package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func getDatabaseConnection() *sql.DB {
	conn := "user=admin password=admin dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func RandomItem(db *sql.DB) (*sql.Rows, error) {

	rows, err := db.Query(`SELECT * FROM Items
		WHERE ID >= CEIL(RANDOM()* (SELECT MAX(ID) FROM Items))
		LIMIT 1`)

	return rows, err
}
