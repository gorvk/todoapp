package database

import "database/sql"

var db *sql.DB

func GetDBInstance() *sql.DB {
	return db
}

func SetDBInstance(instance *sql.DB) {
	db = instance
}
