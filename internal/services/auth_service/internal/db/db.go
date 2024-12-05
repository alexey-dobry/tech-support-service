package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/alexey-dobry/tech-support-platform/internal/services/auth_service/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func NewMySQL(cfg *config.Config) (*sql.DB, error) {
	dsn := cfg.AuthServeer.MySqlDsn

	var db *sql.DB
	var err error

	maxRetries := 10
	delay := 3 * time.Second

	for i := range maxRetries {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			break
		}

		log.Printf("Store connection retry: %d of %d", i+1, maxRetries)

		time.Sleep(delay)
	}

	if err != nil {
		log.Fatalf("Undable to connect: %s", err)
	} else {
		log.Println("Successfully connected")
	}

	err = db.Ping()
	log.Println("Testing the connection")
	if err != nil {
		log.Fatalf("Bad connection: %s", err)
	} else {
		log.Println("Connection is good")
	}

	return db, nil
}
