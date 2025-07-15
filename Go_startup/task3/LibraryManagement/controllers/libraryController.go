package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

func RunConsole(library *services.Library) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n******Library Management System******")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Add Member")
		fmt.Println("0. Exit")
		fmt.Print("Enter choice: ")

		choice := readLine(reader)

		switch choice {
		case "1":
			fmt.Print("Enter Book ID: ")
			id, _ := strconv.Atoi(readLine(reader))
			fmt.Print("Enter Title: ")
			title := readLine(reader)
			fmt.Print("Enter Author: ")
			author := readLine(reader)
			library.AddBook(models.Book{ID: id, Title: title, Author: author, Status: "Available"})
			fmt.Println("Book added.")
		case "2":
			fmt.Print("Enter Book ID to remove: ")
			id, _ := strconv.Atoi(readLine(reader))
			library.RemoveBook(id)
			fmt.Println("Book removed.")
		case "3":
			fmt.Print("Enter Book ID to borrow: ")
			bookID, _ := strconv.Atoi(readLine(reader))
			fmt.Print("Enter Member ID: ")
			memberID, _ := strconv.Atoi(readLine(reader))
			if err := library.BorrowBook(bookID, memberID); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book borrowed.")
			}
		case "4":
			fmt.Print("Enter Book ID to return: ")
			bookID, _ := strconv.Atoi(readLine(reader))
			fmt.Print("Enter Member ID: ")
			memberID, _ := strconv.Atoi(readLine(reader))
			if err := library.ReturnBook(bookID, memberID); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned.")
			}
		case "5":
			books := library.ListAvailableBooks()
			fmt.Println("Available Books:")
			for _, book := range books {
				fmt.Printf("%d: %s by %s\n", book.ID, book.Title, book.Author)
			}
		case "6":
			fmt.Print("Enter Member ID: ")
			memberID, _ := strconv.Atoi(readLine(reader))
			books := library.ListBorrowedBooks(memberID)
			fmt.Printf("Borrowed Books for Member %d:\n", memberID)
			for _, book := range books {
				fmt.Printf("%d: %s by %s\n", book.ID, book.Title, book.Author)
			}
		case "7":
			fmt.Print("Enter Member ID: ")
			id, _ := strconv.Atoi(readLine(reader))
			fmt.Print("Enter Member Name: ")
			name := readLine(reader)
			library.AddMember(models.Member{ID: id, Name: name})
			fmt.Println("Member added.")
		case "0":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}
func readLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}