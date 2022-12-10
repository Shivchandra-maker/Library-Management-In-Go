package driver

import (
	"database/sql"
	"fmt"
	"log"

	// package sql-driver
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	// Open the driver to datasource
	db, err := sql.Open("mysql", "root:root@123@tcp(localhost:3306)/library")
	if err != nil {
		fmt.Println("Failed during Connection")
		log.Fatal(err)
	}

	// Checking Connection To DB
	err = db.Ping()
	if err != nil {
		fmt.Println("Ping Giving error ")
		log.Fatal(err)
	}

	fmt.Println("Connection Established..!!")

	return db, err
}
