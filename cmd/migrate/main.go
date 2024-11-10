package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gorvk/todoapp/database"
	"github.com/gorvk/todoapp/internal/initializers"
)

func main() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	db := database.GetDBInstance()
	dbMigration(db)
}

func dbMigration(db *sql.DB) {
	fmt.Println("DB Migration Started...")

	file, err := os.ReadFile("./database/schema.sql")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(string(file))
	if err != nil {
		panic(err)
	}

	fmt.Println("DB Migration Completed!")
}
