package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"server/models"
	"time"
)

type NoteHandler struct {
	DB *sql.DB
}

func (h *NoteHandler) CreateNoteHandler(w http.ResponseWriter, r *http.Request) {

	var note models.Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO notes (title, content, created_at) VALUES ($1, $2, $3) RETURNING id`
	err := h.DB.QueryRow(query, note.Title, note.Content, time.Now()).Scan(&note.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)

}

func (h *NoteHandler) GetNotesHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := h.DB.Query(`SELECT id, title, content, created_at FROM notes`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		notes = append(notes, note)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)

}
