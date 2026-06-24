package db

import (
	"database/sql"

	_"github.com/mattn/go-sqlite3"
)

func InitDB(path string) (*sql.DB, error){
	db ,err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	createTableQuery := `CREATE TABLE IF NOT EXISTS tasks(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL UNIQUE,
			description TEXT NOT NULL,
			status TEXT NOT NULL DEFAULT('pending'),
			created_at DATETIME NOT NULL
		);`
	
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}

	return db, nil
}