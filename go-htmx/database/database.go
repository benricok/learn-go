package database

import (
	"database/sql"
	"fmt"
	"go-htmx/user"
	"time"

	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

//type Database struct {
var	Db *sql.DB

func NewDatabase(url string) error {
	db, err := sql.Open("libsql", url)
	if err != nil {
		return err
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

	Db = db
	return nil
}

func GetUser(username string) (*user.User,  error) {
	res, err := Db.Query(`SELECT * FROM users WHERE username = $1`, username)
	if err != nil {
		return nil, err
	}

	res.Next()
    var password string
    var surname string
    var name string
    var datecreated time.Time

    err = res.Scan(&password, &name, &surname, &datecreated)
    if err != nil {
        return nil, fmt.Errorf("unable to scan db row: %+v", err)
    }

    res.Close()

    return &user.User{
        Username: username,
        Password: password,
        Name: name,
        Surname: surname,
        Created: datecreated,
    }, nil
}

func AddUser(u *user.User) (error) {
	_, err := Db.Exec(`INSERT INTO users VALUES ($1, $2, $3, $4, $5)`, u.Username,
										u.Password, u.Name, u.Surname, u.Created)
	
	//rowsAffected, err := res.RowsAffected()
	//if err != nil {
	//	return err
	//}

	//fmt.Printf("Book created successfully (%d row affected)\n", rowsAffected)
	return err
}