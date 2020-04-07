package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/alchermd/awm/internal/data"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("web/templates/index.html"))
	t.Execute(w, nil)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var payload data.Message
	p, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	json.Unmarshal(p, &payload)

	db, err := data.New()
	if err != nil {
		log.Print(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	res, err := db.Exec("INSERT INTO messages(name, message, reply_me) VALUES(?, ?, ?)", payload.Name, payload.Message, payload.ReplyMe)
	if err != nil {
		log.Print(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Print(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	var newMessage data.Message
	row := db.QueryRow("SELECT id, name, message, reply_me, created_at FROM messages WHERE id = ?", id)
	err = row.Scan(&newMessage.Id, &newMessage.Name, &newMessage.Message, &newMessage.ReplyMe, &newMessage.CreatedAt)
	if err != nil {
		log.Print(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(&newMessage)
	if err != nil {
		log.Print(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(j))
}
