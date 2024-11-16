package models

import (
	"github.com/gorvk/todoapp/database"
	"github.com/gorvk/todoapp/internal/types"
)

func Create(todo types.Todo) error {
	db := database.GetDBInstance()

	stmt, err := db.Prepare("INSERT INTO Todos VALUES(gen_random_uuid(), $1, $2)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		todo.Title,
		todo.IsCompleted,
	)

	return err
}
