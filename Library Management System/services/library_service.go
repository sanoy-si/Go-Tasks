package services

import (
	"fmt"
	"Library_Management_System/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int64)
	BorrowBook(bookID int64, memberID int64) error
	ReturnBook(bookID int64, memberID int64) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int64) []models.Book
}

type Library struct {
	Books   map[int64]models.Book
	Members map[int64]models.Member
}

func (library *Library) AddMember(member models.Member){
	library.Members[member.ID] = member
}

func (library *Library) AddBook(book models.Book) {
	library.Books[book.ID] = book
}

func (library *Library) RemoveBook(bookID int64) {
	delete(library.Books, bookID)
}

func(library *Library) BorrowBook(bookId int64, memberID int64) error{
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
	library.Books[bookId] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	library.Members[memberID] = member
	return nil

} 


func(library *Library) ReturnBook(bookId int64, memberID int64) error{
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
	library.Books[bookId] = book
	library.Members[memberID] = member
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


func (library *Library) ListBorrowedBooks(memberID int64) []models.Book{
	member := library.Members[memberID]
	return member.BorrowedBooks
}

