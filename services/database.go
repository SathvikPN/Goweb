package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // postgres golang driver // NOTE: use only init functions and nothing else (_)
)

func InitDB() error {
	db := connectDB()
	return createTables(db)
}

func connectDB() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// open connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err) // NOTE: panic allows deferred func, log.Fatal call os.Exit doesnot allow defer func run
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connection Success!")

	return db // return connection
}

func createTables(db *sql.DB) error {
	startQuery := `CREATE TABLE IF NOT EXISTS `
	tableName := `posts`
	endQuery := `(
		post_id SERIAL,
		title TEXT,
		body TEXT,
		PRIMARY KEY (post_id));`

	_, err := db.Exec(startQuery + tableName + endQuery)
	if err != nil {
		fmt.Println("failed to create table", tableName, "error:", err)
		panic(err)
	}

	fmt.Println("tables create Succcess!")
	return nil
}
