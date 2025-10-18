package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "Ru0810/15"
	}
	database := os.Getenv("DB_NAME")
	if database == "" {
		database = "fayda_queue_system"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, database)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("❌ Failed to open database connection:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	log.Println("✅ Connected to MySQL Database!")
}
