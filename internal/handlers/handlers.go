package handlers

import (
	"html/template"
	"net/http"
	"fmt"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("web/templates/index.html"))
	t.Execute(w, nil)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"message": "hello"}`)
}