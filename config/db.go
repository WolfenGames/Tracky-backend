package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb(dbSource string) *sql.DB {
	db, err := sql.Open("sqlite3", dbSource)
	if err != nil {
		log.Fatalf("Failed to connect to SQLite db: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("failed to ping SQLite :%v", err)
	}

	log.Println("Connected to SQLite database successfully")

	DB = db
	CreateTables(DB)
	return db
}

func CreateTables(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS Trackable (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tool TEXT NOT NULL,                 -- e.g., VSCode, Chrome
		metadata TEXT,                      -- JSON or a string (e.g., file path, URL)
		start_time DATETIME NOT NULL,
		end_time DATETIME,
		duration INTEGER,                   -- Duration in seconds
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
	log.Println("Tables created")
}

func GetDB() *sql.DB {
	if DB == nil {
		log.Fatalf("DB not initialized")
	}
	return DB
}
