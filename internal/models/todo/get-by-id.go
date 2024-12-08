package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/todoapp/internal/initializers"
)

func GetById(id string) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, fmt.Errorf("nil db instance")
	}
	query := fmt.Sprintf("SELECT * from Todos WHERE id = '%v'", id)
	fmt.Println(query)
	rows, err := db.Query(query)
	return rows, err
}
