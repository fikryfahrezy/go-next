package database

import (
	"log/slog"
	"sync"
)

type DB struct {
	mu sync.Mutex
	V  map[string]map[string]any
}

func NewDB() (*DB, error) {
	return &DB{V: make(map[string]map[string]any)}, nil
}

func (db *DB) Close() error {
	slog.Info("Closing database connection")
	return nil
}

func (db *DB) Health() error {
	return nil
}

func (db *DB) BeginTx() error {
	db.mu.Lock()
	return nil
}

func (db *DB) Commit() error {
	db.mu.Unlock()
	return nil
}
