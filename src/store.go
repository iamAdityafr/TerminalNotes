package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Note struct {
	ID    int64
	Title string
	Body  string
}

type Store struct {
	conn *sql.DB
}

func (s *Store) Init() error {
	var err error
	s.conn, err = sql.Open("sqlite3", "./notes.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	// defer s.conn.Close()

	if err = s.conn.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Database connection established successfully")

	notes_table := `CREATE TABLE IF NOT EXISTS notes (
		id integer not null primary key,
		title text not null,
		body text not null
	);`

	if _, err := s.conn.Exec(notes_table); err != nil {
		return fmt.Errorf("error in creating notes table: %v", err)
	}

	return nil
}

func (s *Store) GetNotes() ([]Note, error) {
	records, err := s.conn.Query("SELECT * FROM notes")
	if err != nil {
		return nil, fmt.Errorf("unable to query notes: %v", err)
	}

	notes := []Note{}
	defer records.Close()
	for records.Next() {
		note := Note{}
		if err := records.Scan(&note.ID, &note.Title, &note.Body); err != nil {
			return nil, fmt.Errorf("error scanning note: %v", err)
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (s *Store) DeleteNote(id int64) error {
	deleteQuery := `DELETE FROM notes WHERE id = ?`
	result, err := s.conn.Exec(deleteQuery, id)
	if err != nil {
		return fmt.Errorf("failed to delete note with ID %d: %v", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no note found with ID %d", id)
	}

	return nil
}

func (s *Store) SaveNote(note Note) error {
	if note.ID == 0 {
		note.ID = time.Now().UTC().Unix()
	}

	insertQuery := `INSERT INTO notes (id, title, body)
	VALUES (?, ?, ?)
	ON CONFLICT(id) DO UPDATE
	SET title=excluded.title, body=excluded.body;`

	if _, err := s.conn.Exec(insertQuery, note.ID, note.Title, note.Body); err != nil {
		return fmt.Errorf("failed to save note: %w", err)
	}

	return nil
}
