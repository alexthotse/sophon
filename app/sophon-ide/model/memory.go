package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/asg017/sqlite-vec-go-bindings/go"
)

type SophonMemory struct {
	db *sql.DB
}

func NewSophonMemory(dbPath string) (*SophonMemory, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Load sqlite-vec extension
	// Note: This usually requires a specific load command depending on the driver.
	// For this mock implementation, we'll assume it's available.
	_, err = db.Exec("CREATE VIRTUAL TABLE IF NOT EXISTS vec_memory USING vec0(embedding float[1536], metadata text)")
	if err != nil {
		log.Printf("Warning: Could not create virtual table (sqlite-vec might not be loaded): %v", err)
	}

	return &SophonMemory{db: db}, nil
}

func (m *SophonMemory) Save(embedding []float32, metadata string) error {
	_, err := m.db.Exec("INSERT INTO vec_memory(embedding, metadata) VALUES (?, ?)", embedding, metadata)
	return err
}

func (m *SophonMemory) Search(embedding []float32, limit int) ([]string, error) {
	rows, err := m.db.Query("SELECT metadata FROM vec_memory WHERE embedding MATCH ? ORDER BY distance LIMIT ?", embedding, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []string
	for rows.Next() {
		var meta string
		if err := rows.Scan(&meta); err != nil {
			return nil, err
		}
		results = append(results, meta)
	}
	return results, nil
}
