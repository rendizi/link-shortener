package shortener

import (
	"database/sql"
	"errors"
	"log"
	"module_name/encrypt/encrypt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	db, err = Connect()
	if err != nil {
		log.Fatal(err)
	}
}

func Connect() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "pkg/database/sqlite3.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS links (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        url TEXT,
        shorted TEXT
    )
`)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Get(short string) (string, error) {
	rows, err := db.Query("SELECT id, url,shorted FROM links")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var url, shorted string
		if err = rows.Scan(&id, &url, &shorted); err != nil {
			return "", err
		}
		if shorted == short {
			return url, nil
		}
	}
	return "", errors.New("Something went wrong")
}

func Insert(link string) (string, error) {
	stmt, err := db.Prepare("INSERT INTO links(url,shorted) VALUES(?,?)")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	last, err := Last()
	if err != nil {
		return "", err
	}
	short := encrypt.Encrypt(last)
	_, err = stmt.Exec(link, short)
	if err != nil {
		return "", err
	}
	return short, nil
}

func Last() (string, error) {
	rows, err := db.Query("SELECT id, url,shorted FROM links")
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var id int
	var url string
	shorted := "A"
	for rows.Next() {
		if err = rows.Scan(&id, &url, &shorted); err != nil {
			return "", err
		}
	}
	if shorted == "A" {
		shorted = "a"
	}
	return shorted, nil
}
