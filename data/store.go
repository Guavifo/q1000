package data

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// Store is the database
type Store struct {
	database *sql.DB
}

// OpenStore opens the database
func OpenStore(dataResourceName string) (*Store, error) {
	db, err := sql.Open("mysql", dataResourceName)
	if err != nil {
		return nil, err
	}

	return &Store{
			database: db,
		},
		nil
}

// Close the connection to store
func (s *Store) Close() {
	s.database.Close()
}

// Prepare a sql statement for execution
func (s *Store) Prepare(sql string) (*sql.Stmt, error) {
	stmt, err := s.database.Prepare(sql)
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

// Query the database
func (s *Store) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	rows, err := s.database.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
