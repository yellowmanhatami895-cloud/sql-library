package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func getInput(title string) string {
	fmt.Printf("-----%s-----\n", title)

	if !scanner.Scan() {
		return ""
	}

	return strings.TrimSpace(scanner.Text())
}

func getMenu(title string, subjects []string) int {
	fmt.Printf("-----%s-----\n", title)
	for i, s := range subjects {
		fmt.Printf("%d.%s\n", i+1, s)
	}
	input := getInput("Enter menu index")

	intInput, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Enter menu index")
		return -1
	}
	if intInput < 1 || intInput > len(subjects) {
		fmt.Printf("Enter menu index between 1 to %d\n", len(subjects))
	}
	return intInput
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("sqlite", "db/Library.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createDB()
mainLoop:
	for {
		check := checkDB()
		if check == false {
			fmt.Println("DataBase is Empty")
			username := getInput("Enter First username")
			password := getInput("Enter First password")
			err := insertStaffDB([]string{"1", username, password, "Manager"})
			if err != nil {
				fmt.Println(err)
			}
		}
		input := getMenu("Library", []string{"Customer", "Staff"})
		switch input {
		case 0:
			break mainLoop
		case 1:

		case 2:
			username := getInput("Enter user name")
			password := getInput("Enter password")
			if b, s := checkStaff(username, password); b == true {
				switch s {

				case "Manager":

				ManagerLoop:

					for {
						fmt.Println("Welcome ", username)
						input := getMenu("Manager panel", []string{"Staff", "Customers", "Books"})
						switch input {
						case 1:
						ManagerSTFLoop:
							for {
								input = getMenu("Staff", []string{"Add ", "Delete", "Edit", "Check list", "Check one"})
								switch input {
								case 1:
									for {
										id := getInput("Enter unique id or None")
										if id == "0" {
											break
										}
										name := getInput("Enter name")
										if name == "0" {
											break
										}
										password := getInput("Enter password")
										if password == "0" {
											break
										}
										role := getInput("Enter role")
										if role == "0" {
											break
										}

										insertStaffDB([]string{id, name, password, role})

									}
								case 2:
									for {
										id := getInput("Enter unique id ")
										if id == "0" {
											break
										}
										err := removeFromStaffDB(id)
										if err != nil {
											fmt.Println(err)
										}

									}
								case 3:
									for {
										id := getInput("Enter id")
										if id == "0" {
											fmt.Println(err)
											break
										}
										intID, err := strconv.Atoi(id)
										if err != nil {
											fmt.Println(err)
											break
										}
										err = editStaffDB(intID)
										if err != nil {
											fmt.Println(err)
										}
									}
								case 4:
									printStaffDB()
								case 5:
									for {
										id := getInput("Enter id")
										if id == "0" {
											break
										}
										intID, err := strconv.Atoi(id)
										if err != nil {
											fmt.Println(err)
											continue
										}
										printOneRecordStaff(intID)
									}
								case 0:
									break ManagerSTFLoop
								}
							}
						case 2:
						ManagerCTSloop:

							for {
								input = getMenu("Customer", []string{"Add", "Delete", "Edit", "Check list", "Check one"})
								switch input {
								case 0:
									break ManagerCTSloop
								case 1:
									id := getInput("Enter unique id or None")
									name := getInput("Enter name")
									inventory := getInput("Enter inventory")
									password := getInput("Enter password")
									stat := getInput("Enter stat")
									insertCustomerDB([]string{id, name, inventory, password, stat})
								case 2:
									for {
										id := getInput("Enter id")
										if id == "0" {
											break
										}
										removeFromCustomerDB(id)

									}

								case 3:
									for {
										id := getInput("Enter id")
										if id == "0" {
											break
										}
										intID, err := strconv.Atoi(id)
										if err != nil {
											fmt.Println(err)
										}
										editCustomerDB(intID)
									}
								case 4:
									printCustomerDB()
								case 5:
									for {
										id := getInput("Enter id")
										if id == "0" {
											break
										}
										intID, err := strconv.Atoi(id)
										if err != nil {
											fmt.Println(err)
											continue
										}
										printOneRecordCustomer(intID)
									}
								}
							}
						case 3:
							for {
								input = getMenu("Customer", []string{"Add", "Delete", "Edit", "Check list"})
							}
						case 0:
							break ManagerLoop
						}
					}
				case "Librarian":

				case "Cashier":
				}
			}
		}
	}
}
