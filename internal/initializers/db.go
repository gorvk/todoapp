package initializers

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gorvk/todoapp/database"
	_ "github.com/lib/pq"
)

func ConnectDB() {
	fmt.Println("Connecting to DB...")

	var err error

	DB_CONNECTION_STRING := os.Getenv("DB_CONNECTION_STRING")

	// opening the connection with Postgres DB
	db, err := sql.Open("postgres", DB_CONNECTION_STRING)
	if err != nil {
		panic(err)
	}

	// checking the ping to DBinitializers
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	database.SetDBInstance(db)

	fmt.Println("DB Connected Successfully!")
}
