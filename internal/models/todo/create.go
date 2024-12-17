package models

import (
	"github.com/gorvk/todoapp/internal/initializers"
	"github.com/gorvk/todoapp/internal/types"
)

func Create(todo types.Todo, userClaim types.UserClaim) error {
	db := initializers.GetDBInstance()

	stmt, err := db.Prepare("INSERT INTO Todos VALUES(gen_random_uuid(), $1, $2, $3)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		todo.Title,
		todo.IsCompleted,
		userClaim.UserId,
	)

	return err
}
