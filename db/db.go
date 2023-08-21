package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Get_addresses()[]string {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("DB_PATH environment variable is not set")
	}
	// Connect to SQLite database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// SQL query
	query := `SELECT '0x' || substr(input,35,40) AS address from transactions
	WHERE to_address = '0xcf205808ed36593aa40a44f10c7f7c2f67d4a4d4'
	AND SUBSTR(input, 1, 10) = '0x6945b123'
	AND '0x' || substr(input,35,40) = from_address
	ORDER BY blockNumber desc;`

	// Execute query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var addresses []string
	for rows.Next() {
		var address string
		if err := rows.Scan(&address); err != nil {
			log.Fatal(err)
		}
		addresses = append(addresses, address)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Print addresses
	return addresses
}
