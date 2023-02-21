package database

import (
	"database/sql"
	"fmt"
)
/*
This implementation uses the database/sql package to connect to a MySQL database, 
and provides the required methods to implement the database.DB interface, 
such as Query, QueryRow, Exec, and Close. These methods can then be used by the db.
Repository implementation in internal/data/repository/db.go to interact with the database.
*/
type MySQLDB struct {
	db *sql.DB
}

func NewMySQLDB(dataSourceName string) (*MySQLDB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return &MySQLDB{
		db: db,
	}, nil
}

func (db *MySQLDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	return rows, nil
}

func (db *MySQLDB) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.db.QueryRow(query, args...)
}

func (db *MySQLDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.db.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing statement: %v", err)
	}
	return result, nil
}

func (db *MySQLDB) Close() error {
	return db.db.Close()
}
