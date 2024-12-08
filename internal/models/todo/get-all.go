package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/todoapp/internal/initializers"
)

func GetAll() (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, fmt.Errorf("nil db instance")
	}
	query := "SELECT * from Todos"
	rows, err := db.Query(query)
	return rows, err
}
