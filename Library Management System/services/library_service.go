package services

import (
	"fmt"
	"Library_Management_System/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

func (library *Library) AddBook(book models.Book) {
	library.Books[book.ID] = book
	book.Status = "Available"

}

func (library *Library) RemoveBook(bookID int) {
	delete(library.Books, bookID)
}

func(library *Library) BorrowBook(bookId int, memberID int) error{
	book, bookExists := library.Books[bookId]
	member, memberExists := library.Members[memberID]

	if !bookExists{
		return fmt.Errorf("a book with ID %v doesnot exist", bookId)
	}

	if book.Status == "Borrowed"{
		return fmt.Errorf("the book with ID %v is already borrowed", bookId)
	}

	if !memberExists{
		return fmt.Errorf("a member with ID %v doesnot exist", memberID)
	}

	book.Status = "Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	return nil

} 


func(library *Library) ReturnBook(bookId int, memberID int) error{
	book, bookExists := library.Books[bookId]
	member, memberExists := library.Members[memberID]

	if !bookExists{
		return fmt.Errorf("a book with ID %v doesnot exist", bookId)
	}
	
	if !memberExists{
		return fmt.Errorf("a member with ID %v doesnot exist", memberID)
	}
	
	bookIndex := -1
	for i, borrowedBook :=  range member.BorrowedBooks{
		if borrowedBook.ID == bookId{
			bookIndex = i
		}
	}

	if bookIndex == -1{
		return fmt.Errorf("the book with ID %v is not borrowed by the member having ID %v", bookId, memberID)
	}

	book.Status = "Available"
	member.BorrowedBooks = append(member.BorrowedBooks[:bookIndex], member.BorrowedBooks[bookIndex + 1:]...)
	return nil

}


func (library *Library) ListAvailableBooks() []models.Book{
	var availableBooks []models.Book
	for _, book := range library.Books{
		if book.Status == "Available"{
			availableBooks = append(availableBooks, book)
		}
	}

	return availableBooks
}


func (library *Library) ListBorrowedBooks(memberID int) []models.Book{
	member := library.Members[memberID]
	return member.BorrowedBooks
}

