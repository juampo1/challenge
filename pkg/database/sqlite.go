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
		"senderId"	INTEGER NOT NULL,
		"recipientId"	INTEGER NOT NULL,
		"created_at" TEXT NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT)
	)`)

	userTableSt.Exec()
	messageTableSt.Exec()

	return db, nil
}
