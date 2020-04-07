package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func main() {
	dns := os.Getenv("DB_DATABASE")
	if dns == "" {
		log.Fatal("$DB_DATABASE must be set.")
	}

	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}

	stmt :=
		`
CREATE TABLE messages(
	id INT PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(50) NOT NULL,
	message TEXT NOT NULL,
	created_at DATETIME DEFAULT NOW()
)
`
	if _, err := db.Exec(stmt); err != nil {
		log.Fatal(err)
	}
	log.Print("Database initialized.")
}
