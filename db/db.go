package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const DBNAME = "timb"
const DBUSER = "timb"
const DBPASS = "timb"
const DBHOST = "localhost"

func ConnectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		DBUSER,
		DBPASS,
		DBHOST,
		DBNAME)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = initTablesDB(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initTablesDB(db *sql.DB) error {
	res, err := db.Exec(`CREATE TABLE IF NOT EXISTS transaction (
		uuid VARCHAR(255) PRIMARY KEY,
		name VARCHAR(255),
		day SMALLINT NOT NULL,
		month SMALLINT NOT NULL,
		year SMALLINT NOT NULL,
		amount DECIMAL NOT NULL)`)
	if err != nil {
		return err
	}
	aff, _ := res.RowsAffected()
	fmt.Println("initTablesDB: Affected=", aff)
	return nil

}
