package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./challege.db")

	if err != nil {
		return nil, err
	}

	userTableSt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS "user" (
		"id"	INTEGER NOT NULL UNIQUE,
		"username"	TEXT NOT NULL UNIQUE,
		"password"	TEXT NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT)
	)`)

	messageTableSt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS "message" (
		"id"	INTEGER NOT NULL UNIQUE,
		"sender_id"	INTEGER NOT NULL,
		"recipient_id"	INTEGER NOT NULL,
		"created_at" TEXT NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT)
	)`)

	contentTableSt, _ := db.Prepare(`CREATE TABLE "content" (
		"id"	INTEGER NOT NULL UNIQUE,
		"message_id"	INTEGER NOT NULL UNIQUE,
		"content_type"	TEXT NOT NULL,
		"text"	TEXT,
		FOREIGN KEY("message_id") REFERENCES "user"("id"),
		PRIMARY KEY("id" AUTOINCREMENT)
	)`)

	userTableSt.Exec()
	messageTableSt.Exec()
	contentTableSt.Exec()

	return db, nil
}
