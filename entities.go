package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//GetDB is imported
func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "testdb"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}
