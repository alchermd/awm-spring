package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Print("Starting application.")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set.")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("web/templates/index.html"))
		t.Execute(w, nil)
	}).Methods("GET")

	assetPath := "/assets/"
	log.Print("Serving static files on " + assetPath)
	r.PathPrefix(assetPath).Handler(http.StripPrefix(assetPath, http.FileServer(http.Dir("web/assets/"))))

	log.Print("Starting server on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
