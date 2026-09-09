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

func getIntInput(title string) int {
	input := getInput(title)
	intInput, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return intInput
}

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
			insertIntoStaffDB([]string{"1", username, password, "Manager"})

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
								input = getMenu("Staffs", []string{"Add ", "Delete", "Edit", "Check list", "Check one"})
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

										insertIntoStaffDB([]string{id, name, password, role})

									}
								case 2:
									for {
										id := getIntInput("Enter unique id ")
										if id == 0 {
											break
										}
										removeFromStaffDB(id)
										if err != nil {
											fmt.Println(err)
										}

									}
								case 3:
									for {
										id := getIntInput("Enter id")
										if id == 0 {
											fmt.Println(err)
											break
										}

										editFromStaffDB(id)

									}
								case 4:
									printStaffDB()
								case 5:
									for {
										id := getIntInput("Enter id")
										if id == 0 {
											break
										}

										printOneRecordStaff(id)
									}
								case 0:
									break ManagerSTFLoop
								}
							}
						case 2:
						ManagerCTSloop:

							for {
								input = getMenu("Customers", []string{"Add", "Delete", "Edit", "Check list", "Check one"})
								switch input {
								case 0:
									break ManagerCTSloop
								case 1:
									id := getInput("Enter unique id or None")
									name := getInput("Enter name")
									inventory := getInput("Enter inventory")
									password := getInput("Enter password")
									stat := getInput("Enter stat")
									insertIntoCustomerDB([]string{id, name, inventory, password, stat})
								case 2:
									for {
										id := getIntInput("Enter id")
										if id == 0 {
											break
										}
										removeFromCustomerDB(id)

									}

								case 3:
									for {
										id := getIntInput("Enter id")
										if id == 0 {
											break
										}

										editFromCustomerDB(id)
									}
								case 4:
									printCustomerDB()
								case 5:
									for {
										id := getIntInput("Enter id")
										if id == 0 {
											break
										}

										printOneRecordCustomer(id)
									}
								}
							}
						case 3:
						ManagerBKSloop:
							for {
								input = getMenu("Books", []string{"Add", "Delete", "Edit", "Check list", "Check one"})
								switch input {
								case 1:
									for {
										id := getInput("Enter id")
										if id == "0" {
											break
										}
										name := getInput("Enter name")
										if name == "0" {
											break
										}
										price := getInput("Enter price")
										if price == "0" {
											break
										}

										quantity := getInput("Enter quantity")
										if quantity == "0" {
											break
										}
										stat := getInput("Enter stat")
										if stat == "0" {
											break
										}
										insertIntoBookDB([]string{id, name, price, quantity, stat})
									}

								case 2:
									for {
										id := getIntInput("Enter id")
										if id == 0 {
											break
										}
										removeFromBookDB(id)
									}
								case 3:
									for {
										id := getIntInput("Enter id")
										if id == 0 {
											break
										}
										editFromBookDB(id)
									}
								case 4:
									printBookDB()

								case 5:
									for {
										id := getIntInput("Enter id")
										if id == 0 {
											break
										}
										printOneRecordBook(id)
									}
								case 0:
									break ManagerBKSloop
								}
							}
						case 0:
							break ManagerLoop
						}
					}
				case "Librarian":
				LibrarianMainLoop:
					for {
						input := getMenu("Librarian", []string{"List books", "Search books", "Delete books", "Edit books", "Add book", "Issue books", "Receive returned books"})
						switch input {
						case 0:
							break LibrarianMainLoop
						case 1:
							printBookDB()
						case 2:
						SearchBook:
							for {
								id := getIntInput("Enter id")
								if id == 0 {
									break SearchBook
								}
								printOneRecordBook(id)
							}
						case 3:
						DeleteBook:
							for {
								id := getIntInput("Enter id")
								if id == 0 {
									break DeleteBook
								}
								removeFromBookDB(id)
							}
						case 4:
						EditBook:
							for {
								id := getIntInput("Enter id")
								if id == 0 {
									break EditBook
								}
								editFromBookDB(id)
							}
						case 5:
						AddBook:
							for {
								id := getInput("Enter id")
								if id == "0" {
									break AddBook
								}
								name := getInput("Enter name")
								if id == "0" {
									break AddBook
								}
								price := getInput("Enter price")
								if price == "0" {
									break AddBook
								}
								quantity := getInput("Enter quantity")
								if quantity == "0" {
									break AddBook
								}
								stat := getInput("Enter stat")
								if stat == "0" {
									break AddBook
								}
								insertIntoBookDB([]string{id, name, price, quantity, stat})

							}
						case 6:
						IssueBook:
							for {
								bookID := getIntInput("Enter book id")
								if bookID == 0 {
									break IssueBook
								}
								customerID := getIntInput("Enter book id")
								if customerID == 0 {
									break IssueBook
								}
								issueBook(bookID, customerID)
							}
						case 7:
						ReceiveBooks:
							for {
								bookID := getIntInput("Enter book id")
								if bookID == 0 {
									break ReceiveBooks
								}
								customerID := getIntInput("Enter book id")
								if customerID == 0 {
									break ReceiveBooks
								}
								id := getIntInput("Enter safekeeping id ")
								if id == 0 {
									break ReceiveBooks
								}
								receiveBooks(bookID, customerID, id)
							}
						}
					}
				case "Cashier":
				}
			}
		}
	}
}
