package database

import "database/sql"

type Database struct {
	Db *sql.DB
}

func NewDatabase(url string) (*Database, error) {
	db, err := sql.Open("libsql", url)
	if err != nil {
		return nil, err
	}

	db.Exec(`CREATE TABLE [IF NOT EXISTS] users (
				username TEXT NOT NULL PRIMARY KEY,
				password TEXT NOT NULL,
				name TEXT,
				surname TEXT,
				datecreated TEXT);
			CREATE TABLE [IF NOT EXISTS] items (
				id TEST NOT NULL PRIMARY KEY,
				owner TEXT NOT NULL,
				name TEXT NOT NULL,
				done INTEGER NOT NULL,
				datecreated TEXT NOT NULL,
				datecompleted TEXT,
				FOREIGN KEY(owner) REFERENCES users(username));
	`)

	return &Database{db}, nil
}