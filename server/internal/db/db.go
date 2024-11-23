package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQL() (*sql.DB, error) {
	dsn := "myuser:mypassword@tcp(docker.for.mac.localhost:3306)/mydatabase"

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
	defer db.Close()

	err = db.Ping()
	log.Println("Testing the connection")
	if err != nil {
		log.Fatalf("Bad connection: %s", err)
	} else {
		log.Println("Connection is good")
	}

	sqlFile := "../init.sql"

	script, err := os.ReadFile(sqlFile)

	if err != nil {
		log.Fatalf("Unable to read sql init script file: %s", err)
	}

	commands := string(script)

	_, err = db.Exec(commands)

	if err != nil {
		log.Fatalf("Unable to execute sql init script: %s", err)
	}

	return db, nil
}
