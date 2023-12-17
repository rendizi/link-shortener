package shortener

import (
	"database/sql"
	"fmt"
	"log"
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

	db, err := sql.Open("sqlite3", "./database/sqlite3.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	//Create table if required

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS links (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        url TEXT
    )
`)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Get() error {
	rows, err := db.Query("SELECT id, url FROM links")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var url string
		if err = rows.Scan(&id, &url); err != nil {
			return err
		}
		fmt.Println("ID:", id, "URL:", url)
	}
	return nil
}

func Insert(link string) error {
	stmt, err := db.Prepare("INSERT INTO links(url) VALUES(?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(link)
	if err != nil {
		return err
	}
	return nil
}
