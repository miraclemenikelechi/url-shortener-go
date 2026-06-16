package memory

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB() *sql.DB {
	DATABASE_URL := os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		log.Fatal("database url not set")
	}

	db, err := sql.Open("pgx", DATABASE_URL)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("cannot ping database", err)
	}

	log.Println("connected to database")
	return db
}
