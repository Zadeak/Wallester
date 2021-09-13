package main

import (
	"database/sql"
	"fmt"
)

func dbConn() (db *sql.DB) {
	dbHost := "localhost"
	dbPort := 5432
	dbDriver := "postgres"
	dbUser := "postgres"
	dbPass := "postgres"
	dbName := "postgres"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := sql.Open(dbDriver, psqlInfo)
	if err != nil {
		panic(err.Error())
	}
	return db
}
