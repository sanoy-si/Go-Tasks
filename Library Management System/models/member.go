package models

type Member struct{
	ID int64
	Name string
	BorrowedBooks []Book
}