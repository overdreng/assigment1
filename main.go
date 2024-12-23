package main

import (
	"bufio"
	"fmt"
	library "github.com/AdiMukhamedzhan/Assignment1/Library/Library"
	"os"
	"strings"

	"github.com/AdiMukhamedzhan/Assignment1/Library" // Импорт пакета Library
)

func main() {
	lib := library.Library{Books: make(map[string]library.Book)}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("1. Add")
		fmt.Println("2. Borrow")
		fmt.Println("3. Return")
		fmt.Println("4. List")
		fmt.Println("5. Exit")

		fmt.Print("Enter your choice: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter book ID: ")
			scanner.Scan()
			id := scanner.Text()

			fmt.Print("Enter book title: ")
			scanner.Scan()
			title := scanner.Text()

			fmt.Print("Enter book author: ")
			scanner.Scan()
			author := scanner.Text()

			book := library.Book{
				ID:         id,
				Title:      title,
				Author:     author,
				IsBorrowed: false,
			}
			lib.AddBook(book)
			fmt.Println("Book added successfully.")

		case "2":
			fmt.Print("Enter book ID to borrow: ")
			scanner.Scan()
			id := scanner.Text()
			lib.BorrowBook(id)
			fmt.Println("Book borrowed successfully.")

		case "3":
			fmt.Print("Enter book ID to return: ")
			scanner.Scan()
			id := scanner.Text()
			lib.ReturnBook(id)
			fmt.Println("Book returned successfully.")

		case "4":
			books := lib.ListBooks()
			for _, book := range books {
				fmt.Printf("ID: %s, Title: %s, Author: %s, IsBorrowed: %t\n", book.ID, book.Title, book.Author, book.IsBorrowed)
			}

		case "5":
			fmt.Println("Exiting program.")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
