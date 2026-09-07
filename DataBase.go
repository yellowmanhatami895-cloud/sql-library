package main

import (
	"fmt"
	"log"
	"strconv"

	_ "github.com/glebarez/go-sqlite"
)

func createDB() {

	db.Exec("PRAGMA foreign_keys = ON")
	db.Exec("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY AUTOINCREMENT,name TEXT NOT NULL,price INT NOT NULL,quantity INT NOT NULL,stat TEXT DEFAULT 'Free');")
	db.Exec("CREATE TABLE IF NOT EXISTS staff (id INTEGER PRIMARY KEY AUTOINCREMENT,name TEXT NOT NULL,password TEXT NOT NULL,role TEXT NOT NULL DEFAULT 'Librarian');")
	db.Exec("CREATE TABLE IF NOT EXISTS customer (id INTEGER PRIMARY KEY AUTOINCREMENT,name TEXT NOT NULL,inventory INT NOT NULL,password TEXT NOT NULL,stat TEXT DEFAULT 'Free');")
	db.Exec("CREATE TABLE IF NOT EXISTS safekeeping (id INTEGER PRIMARY KEY AUTOINCREMENT,bookID INT NOT NULL,customerID INT NOT NULL,FOREIGN KEY (bookID) REFERENCES books(id),FOREIGN KEY (customerID) REFERENCES customer(id));")
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
func insertBookDB(record []string) error {
	if len(record) != 5 {
		return fmt.Errorf("Wrong length")
	}
	id, err := strconv.Atoi(record[0])
	if err != nil {
		return fmt.Errorf("id most be integer")
	}

	price, err := strconv.Atoi(record[2])

	if err != nil {
		fmt.Println(err)
		return err
	}
	quantity, err := strconv.Atoi(record[3])

	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = db.Exec("INSERT INTO books VALUES(?,?,?,?,?) ", id, record[1], price, quantity, record[4])

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func printBookDB() {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var price string
		var quantity string
		var stat string
		err := rows.Scan(&id, &name, &price, &quantity, &stat)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(id, name, price, quantity, stat)
	}

}
func removeFromBookDB(id string) {
	var name string
	db.QueryRow("SELECT name FROM books WHERE id = ?", id).Scan(&name)
	if name == "" {
		fmt.Println("id not found")
		return
	}
	_, err := db.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name + " Has been deleted")

}
func editBookDB(id int) {
	var row string
	var name string
	var price string
	var quantity string
	var stat string

	err := db.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&id, &name, &price, &quantity, &stat)
	if name == "" {
		fmt.Println("id not found")
	}
	if err != nil {
		fmt.Println(err)

	}
	row = strconv.Itoa(id) + "  " + name + "  " + price + "  " + quantity + "  " + stat
	fmt.Println(row)
	fmt.Println("Enter new row")
	name = getInput("Enter name")
	price = getInput("Enter price")
	quantity = getInput("Enter quantity")
	stat = getInput("Enter stat")
	_, err = db.Exec("UPDATE books SET name = ?,price = ? ,quantity = ?, stat = ? WHERE id = ?", name, price, quantity, stat, id)
	if err != nil {
		fmt.Println(err)
		return
	}

}
func printOneRecordBook(id int) {
	var name string
	var price int
	var quantity string
	var stat string
	err := db.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&id, &name, &price, &quantity, &stat)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name + " " + strconv.Itoa(price) + " " + quantity + " " + stat)
}
func printOneRecordStaff(id int) {
	var name string
	var password string
	var role string
	err := db.QueryRow("SELECT * FROM staff WHERE id = ?", id).Scan(&id, &name, &password, &role)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name + " " + password + " " + role)

}

func editStaffDB(id int) {
	var row string
	var name string
	var password string
	var role string

	err := db.QueryRow("SELECT * FROM staff WHERE id = ?", id).Scan(&id, &name, &password, &role)
	if err != nil {
		fmt.Println(err)
		return
	}
	row = strconv.Itoa(id) + "  " + name + "  " + password + "  " + role
	fmt.Println(row)
	fmt.Println("Enter new row")
	name = getInput("Enter name")
	password = getInput("Enter password")
	role = getInput("Enter role")
	_, err = db.Exec("UPDATE staff SET name = ?, password = ?, role = ? WHERE id = ?", name, password, role, id)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func insertStaffDB(record []string) {
	if len(record) != 4 {
		fmt.Println("Wrong length")
		return
	}
	id, err := strconv.Atoi(record[0])
	if err != nil {
		fmt.Println("id most be integer")
		return
	}
	_, err = db.Exec("INSERT INTO staff VALUES(?,?,?,?) ", id, record[1], record[2], record[3])
	if err != nil {
		fmt.Println(err)
		return
	}

}

func removeFromStaffDB(id string) {
	var name string
	db.QueryRow("SELECT name FROM staff WHERE id = ?", id).Scan(&name)
	if name == "" {
		fmt.Println("id not found")
		return
	}
	_, err := db.Exec("DELETE FROM staff WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name + "Has been deleted")

}
func checkStaff(username string, password string) (bool, string) {
	var pass string
	var role string
	err := db.QueryRow("SELECT password,role FROM staff WHERE name = ?", username).Scan(&pass, &role)
	if pass == "" {
		fmt.Println("user not found")
	}
	if err != nil {
		fmt.Println(err)
	}
	if pass == password {
		return true, role
	}
	fmt.Println("wrong password")
	return false, ""
}
func printStaffDB() {
	rows, err := db.Query("SELECT * FROM staff")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var password string
		var role string
		err := rows.Scan(&id, &name, &password, &role)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(id, name, password, role)
	}

}
func removeFromCustomerDB(id string) {
	var name string
	db.QueryRow("SELECT name FROM customer WHERE id = ?", id).Scan(&name)
	if name == "" {
		fmt.Println("id not found")
		return
	}
	_, err := db.Exec("DELETE FROM customer WHERE id = ?", id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name + " Has been deleted")

}
func editCustomerDB(id int) {
	var row string
	var name string
	var inventory string
	var password string
	var stat string

	err := db.QueryRow("SELECT * FROM customer WHERE id = ?", id).Scan(&id, &name, &inventory, &password, &stat)
	if name == "" {
		fmt.Println("id not found")
	}
	if err != nil {
		fmt.Println(err)

	}
	row = strconv.Itoa(id) + "  " + name + "  " + inventory + "  " + password + "  " + stat
	fmt.Println(row)
	fmt.Println("Enter new row")
	name = getInput("Enter name")
	inventory = getInput("Enter inventory")
	password = getInput("Enter password")
	stat = getInput("Enter stat")
	_, err = db.Exec("UPDATE customer SET name = ?,inventory = ? ,password = ?, stat = ? WHERE id = ?", name, inventory, password, stat, id)
	if err != nil {
		fmt.Println(err)
		return
	}

}
func printCustomerDB() {
	rows, err := db.Query("SELECT * FROM customer")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var inventory string
		var password string
		var stat string
		err := rows.Scan(&id, &name, &inventory, &password, &stat)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(id, name, inventory, password, stat)
	}

}

func printOneRecordCustomer(id int) {
	var name string
	var inventory int
	var password string
	var stat string
	err := db.QueryRow("SELECT * FROM customer WHERE id = ?", id).Scan(&id, &name, &inventory, &password, &stat)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name + " " + strconv.Itoa(inventory) + " " + password + " " + stat)
}
func insertCustomerDB(record []string) {
	if len(record) != 5 {
		fmt.Println("Wrong length ")
		return
	}
	id, err := strconv.Atoi(record[0])
	if err != nil {
		fmt.Println("id most bbe integer")
		return
	}

	inventory, err := strconv.Atoi(record[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = db.Exec("INSERT INTO customer VALUES(?,?,?,?,?) ", id, record[1], inventory, record[3], record[4])

	if err != nil {
		fmt.Println(err)
		return
	}

}
