package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"server/config"
	"server/handlers"
)

func main() {

	cfg := config.LoadConfig()

	connStr := "user=" + cfg.DBUser + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " sslmode=disable port=" + cfg.DBPort

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Error connection to database: %v", err)
	}

	defer db.Close()

	noteHandler := &handlers.NoteHandler{DB: db}

	http.HandleFunc("/notes", noteHandler.GetNotesHandler)
	http.HandleFunc("/notes/create", noteHandler.CreateNoteHandler)

	log.Printf("Server starting on port ...")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Serve failed: %v", err)
	}

}
