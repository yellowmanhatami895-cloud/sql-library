package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func getInput(title string) string {
	fmt.Printf("-----%s-----\n", title)
	scanner.Scan()
	s := scanner.Text()
	return s
}

func getMenu(title string, subjects []string) int {
	fmt.Printf("-----%s-----\n", title)
	for i, s := range subjects {
		fmt.Printf("%d.%s\n", i+1, s)
	}
	input := getInput("Enter menu index")
	if input == "0" {
		return 0
	}
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
		input := getMenu("Library", []string{"Customer", "Staff", "Manager"})
		switch input {
		case 0:
			break mainLoop
		case 1:

		case 2:

		case 3:
			username := getInput("Enter user name")
			password := getInput("Enter password")
			if checkStaff(username, password) {
			ManagerLoop:
				for {
					fmt.Println("Welcome ", username)
					input := getMenu("Manager panel", []string{"Staff", "Customers", "Books"})
					switch input {
					case 1:
					ManagerSTFLoop:
						for {
							input = getMenu("Staff", []string{"Add staff", "Delete staff", "Check list"})
							switch input {
							case 1:
								id := getInput("Enter unique id")
								name := getInput("Enter name")
								password := getInput("Enter password")
								role := getInput("Enter role")
								insertStaffDB([]string{id, name, password, role})
							case 2:

							case 3:

							case 0:
								break ManagerSTFLoop
							}
						}
					case 2:

					case 3:

					case 0:
						break ManagerLoop
					}
				}
			}
		}
	}
}
