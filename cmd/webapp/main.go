package main

import (
	"github.com/alchermd/awm/internal/handlers"
	"github.com/gorilla/mux"
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
	r.HandleFunc("/", handlers.Home).Methods("GET")

	assetPath := "/assets/"
	log.Print("Serving static files on " + assetPath)
	r.PathPrefix(assetPath).Handler(http.StripPrefix(assetPath, http.FileServer(http.Dir("web/assets/"))))

	log.Print("Starting server on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
