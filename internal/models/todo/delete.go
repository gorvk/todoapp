package models

import (
	"fmt"

	"github.com/gorvk/todoapp/database"
)

func Delete(id string) error {
	db := database.GetDBInstance()
	if db == nil {
		return fmt.Errorf("nil db instance")
	}

	stmt, err := db.Prepare("DELETE FROM Todos WHERE id = $1")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}
