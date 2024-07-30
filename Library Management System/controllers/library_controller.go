package controllers

import (
	"Library_Management_System/models"
	"Library_Management_System/services"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
var library = services.Library{make(map[int64]models.Book), make(map[int64]models.Member)}

func recieveInt(inputText string) int64{
	fmt.Print(inputText)
	scanner.Scan()
	input := scanner.Text()
	value, err := strconv.ParseInt(input, 10, 64)

	for err != nil{
		fmt.Print("Please Enter a vaild integer input: ")
		scanner.Scan()
		input := scanner.Text()
		value, err = strconv.ParseInt(input, 10, 32)

	}

	return value
}


func AddBook(){
	id := recieveInt("Enter the id of the book: ")

	fmt.Print("Enter the Title of the book: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("Enter the Author of the book: ")
	scanner.Scan()
	author := scanner.Text()

	var newBook models.Book
	newBook.ID = id
	newBook.Tiltle = title
	newBook.Author = author
	newBook.Status = "Available"


	library.AddBook(newBook)
	fmt.Println("Book added Successfully.")
}


func RemoveBook(){
	id := recieveInt("Enter the id of the book: ")
	library.RemoveBook(id)
}

func BorrowBook(){
	bookId := recieveInt("Enter the id of the book to borrow: ")
	memberId := recieveInt("Enter the id of the member: ")

	if err := library.BorrowBook(bookId, memberId); err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("The Book with the id %v borrowed by the member with the id %v successfully", bookId, memberId)
	}
}

func ReturnBook(){
	bookId := recieveInt("Enter the id of the book to return: ")
	memberId := recieveInt("Enter the id of the member: ")

	if err := library.ReturnBook(bookId, memberId); err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("The Book with the id %v returned by the member with the id %v successfully\n", bookId, memberId)
	}
}

func ListAvailableBooks(){
	fmt.Println(library.ListAvailableBooks())
}

func ListBorrowedBooks(){
	memberId := recieveInt("Enter the id of the member: ")
	fmt.Println(library.ListBorrowedBooks(memberId))
}

func AddMember(){
	memberId := recieveInt("Enter the id of the member: ")

	fmt.Print("Enter the Name of the member: ")
	scanner.Scan()
	name := scanner.Text()

	var newMember = models.Member{}
	newMember.ID = memberId
	newMember.Name = name

	library.AddMember(newMember)
}

func Welcome(){
	for {
		fmt.Println("Welcome to Library Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Add Member")
		fmt.Println("8. Exit")
		fmt.Print("Enter the number associated with the action you want to perform: ")
		

		scanner.Scan()
		input := scanner.Text()
		choice, err := strconv.ParseInt(input, 10, 64)

		for err != nil{
			fmt.Print("Invalid input. Please Enter the numbers that are available in the menu(1 - 7): ")
			scanner.Scan()
			input := scanner.Text()
			choice, err = strconv.ParseInt(input, 10, 64)

		}

		switch choice{
		case 1:
			AddBook()
		
		case 2:
			RemoveBook()

		case 3:
			BorrowBook()
		
		case 4:
			ReturnBook()
		
		case 5:
			ListAvailableBooks()
		
		case 6:
			ListBorrowedBooks()

		case 7:
			AddMember()
		
		case 8:
			os.Exit(0) 
		
		default:
			fmt.Print("Please enter a value withint the range (1, 7)")
		}
	}
	
}



