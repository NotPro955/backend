package main

import (
	"log"
	"os"

	"github.com/NotPro955/backend/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: mirgrate-<up/down>")
	}

	cfg := config.MustLoad()

	m, err := migrate.New(
		"file://migrations",
		cfg.Databaseurl,
	)
	if err != nil {
		log.Fatalf("migration.new: %v", err)
	}

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatal("Migrate up:", err)
		}
		log.Println("up called")
	case "down":
		if err := m.Down(); err != nil {
			log.Fatal("Migrate down:", err)
		}
		log.Println("down called")
	default:
		log.Fatalf("unknow command: %v", os.Args[1])
	}

	log.Printf("SQL table created")
}
