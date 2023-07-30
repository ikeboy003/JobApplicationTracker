package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// sql.Open open database driver
// perform a blank import to use the coorrrect db driver
func GetDBConnection() (*sql.DB, error) {
	connStr := "user=postgres password=postgres1 dbname=JobTrackerDB sslmode=verify-full"
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {

		return nil, err
	}
	defer dbConn.Close()

	pingErr := dbConn.Ping()

	if pingErr != nil {
		fmt.Print(err)
	}
	fmt.Println("Connected!")
	return dbConn, nil
}
