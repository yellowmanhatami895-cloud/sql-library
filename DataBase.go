package main

import (
	"fmt"
	"log"
	"strconv"

	_ "github.com/glebarez/go-sqlite"
)

func createDB() {

	db.Exec("PRAGMA foreign_keys = ON")
	db.Exec("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY AUTOINCREMENT,name TEXT NOT NULL,price INT NOT NULL,quantity INT NOT NULL);")
	db.Exec("CREATE TABLE IF NOT EXISTS staff (id INTEGER PRIMARY KEY AUTOINCREMENT,name TEXT NOT NULL,password TEXT NOT NULL,role TEXT NOT NULL DEFAULT 'Librarian');")
	db.Exec("CREATE TABLE IF NOT EXISTS customer (id INTEGER PRIMARY KEY AUTOINCREMENT,name TEXT NOT NULL,inventory INT NOT NULL,role TEXT NOT NULL DEFAULT 'customer');")
	db.Exec("CREATE TABLE IF NOT EXISTS safekeeping (id INTEGER PRIMARY KEY AUTOINCREMENT,bookID INT NOT NULL,customerID INT NOT NULL,FOREIGN KEY (bookID) REFERENCES books(id),FOREIGN KEY (customerID) REFERENCES customer(id),);")
	fmt.Println("data base connected")
}

func checkDB() bool {

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM staff").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		return false
	}
	return true
}
func insertStaffDB(record []string) error {
	if len(record) != 4 {
		return fmt.Errorf("Wrong length")
	}
	id, err := strconv.Atoi(record[0])
	if err != nil {
		return fmt.Errorf("id most be integer")
	}
	db.Exec("INSERT INTO staff VALUES(?,?,?,?) ", id, record[1], record[2], record[3])
	return nil
}
func checkStaff(username string, password string) bool {
	var pass string

	err := db.QueryRow("SELECT password FROM staff WHERE name = ?", username).Scan(&pass)
	if err != nil {
		log.Fatal(err)
	}
	if pass == password {
		return true
	}
	return false
}
