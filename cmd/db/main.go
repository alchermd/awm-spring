package main

import (
	"github.com/alchermd/awm/internal/data"
	"log"
)

func main() {
	db, err := data.New()
	if err != nil {
		log.Fatal(err)
	}

	stmt :=
		`
CREATE TABLE messages(
	id INT PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(50) NOT NULL,
	message TEXT NOT NULL,
	reply_me BOOL NOT NULL,
	created_at DATETIME DEFAULT NOW()
)
`
	if _, err := db.Exec(stmt); err != nil {
		log.Fatal(err)
	}
	log.Print("Database initialized.")
}
