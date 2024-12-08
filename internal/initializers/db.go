package initializers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func GetDBInstance() *sql.DB {
	return db
}

func ConnectDB() {
	fmt.Println("Connecting to DB...")

	var err error

	DB_CONNECTION_STRING := os.Getenv("DB_CONNECTION_STRING")

	// opening the connection with Postgres DB
	db, err = sql.Open("postgres", DB_CONNECTION_STRING)
	if err != nil {
		panic(err)
	}

	// checking the ping to DBinitializers
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("DB Connected Successfully!")
}
