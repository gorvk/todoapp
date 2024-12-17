package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/todoapp/internal/initializers"
	"github.com/gorvk/todoapp/internal/types"
)

func GetAll(userClaim types.UserClaim) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, fmt.Errorf("nil db instance")
	}
	query := fmt.Sprintf("SELECT * from Todos WHERE user_id = '%v'", userClaim.UserId)
	fmt.Println(query)
	rows, err := db.Query(query)
	return rows, err
}
