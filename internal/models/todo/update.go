package models

import (
	"github.com/gorvk/todoapp/internal/initializers"
	"github.com/gorvk/todoapp/internal/types"
)

func Update(todo types.Todo, userClaim types.UserClaim) error {
	db := initializers.GetDBInstance()

	stmt, err := db.Prepare("UPDATE Todos SET title = $1, is_completed = $2 WHERE id = $3 AND user_id = $4")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		todo.Title,
		todo.IsCompleted,
		todo.Id,
		userClaim.UserId,
	)

	return err
}
