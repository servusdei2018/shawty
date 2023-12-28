package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/servusdei2018/shawty/pkg/id"
)

// DB contains database data.
type DB struct {
	db *sql.DB
	// cached count of all URLs stored
	count int
	// whether the count needs to be refreshed
	countRefresh bool
}

// Init initializes the database.
func (d *DB) Init() (err error) {
	d.db, err = sql.Open("sqlite3", "./shawty.db")
	if err != nil {
		return
	}
	statement := `
	CREATE TABLE IF NOT EXISTS urls(
		id TEXT PRIMARY KEY,
		url TEXT
	);
	`
	d.countRefresh = true
	_, err = d.db.Exec(statement)
	return
}

// Close closes the underlying database connection.
func (d *DB) Close() error {
	return d.db.Close()
}

// Delete removes a URL by shortID.
func (d *DB) Delete(shortID string) (err error) {
	statement := `
	DELETE FROM urls WHERE id = ?;
	`
	d.countRefresh = true
	_, err = d.db.Exec(statement, shortID)
	return err
}

// Get retrieves a URL by shortID.
func (d *DB) Get(shortID string) (url string, err error) {
	statement := `
	SELECT url FROM urls WHERE id = ?;
	`
	err = d.db.QueryRow(statement, shortID).Scan(&url)
	return
}

// Stats contains database statistics.
type Stats struct {
	// Count contains the number of URLs stored.
	Count int
}

// Stats returns database stats.
func (d *DB) Stats() (stats Stats) {
	if !d.countRefresh {
		stats.Count = d.count
		return
	}
	statement := `
	SELECT COUNT(*) FROM urls;
	`
	d.db.QueryRow(statement).Scan(&stats.Count)
	d.count = stats.Count
	d.countRefresh = false
	return
}

// Store stores a URL, and returns a shortID.
func (d *DB) Store(url string) (shortID string, err error) {
	shortID = id.New()
	statement := `
	INSERT INTO urls (id, url) VALUES (?, ?);
	`
	d.countRefresh = true
	_, err = d.db.Exec(statement, shortID, url)
	return
}
