package db

import (
	"database/sql"
	"log"
	"os"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getConnectionString() string {
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	return fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s", user, name, password, host, port, sslmode)
}

// Connect to PostgreSQL database
func connect() *sql.DB {
	connStr := getConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Get_addresses() []string {
	db := connect()
	defer db.Close()

	query := `SELECT '0x' || substr(input,35,40) AS address from transactions
	WHERE to_address = '0xcf205808ed36593aa40a44f10c7f7c2f67d4a4d4'
	AND SUBSTR(input, 1, 10) = '0x6945b123'
	AND '0x' || substr(input,35,40) = from_address
	ORDER BY blockNumber desc;`

	return executeQuery(db, query)
}

func Get_addresses_from_block(lastProcessedBlock int) []string {
	db := connect()
	defer db.Close()

	query := `SELECT '0x' || substr(input,35,40) AS address from transactions
	WHERE to_address = '0xcf205808ed36593aa40a44f10c7f7c2f67d4a4d4'
	AND SUBSTR(input, 1, 10) = '0x6945b123'
	AND '0x' || substr(input,35,40) = from_address
	AND blockNumber > ?
	ORDER BY blockNumber desc;`

	return executeQuery(db, query, lastProcessedBlock)
}


func executeQuery(db *sql.DB, query string, args ...interface{}) []string {
	rows, err := db.Query(query, args...)
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

	return addresses
}
