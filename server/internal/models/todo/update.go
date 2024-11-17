package models

import (
	"github.com/gorvk/todoapp/database"
	"github.com/gorvk/todoapp/internal/types"
)

func Update(todo types.Todo) error {
	db := database.GetDBInstance()

	stmt, err := db.Prepare("UPDATE Todos SET title = $1, isCompleted = $2 WHERE id = $3")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		todo.Title,
		todo.IsCompleted,
		todo.Id,
	)

	return err
}
