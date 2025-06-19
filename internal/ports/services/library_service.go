package services

import "library/internal/domain"

type LibraryService interface {
	BorrowBook(readerID int, bookID int) error
	ReturnBook(readerID int, bookID int) error
	MarkBookAsLost(readerID int, bookID int) error
	GetReaderBooks(readerID int) ([]*domain.Book, error)
	GetAvailableBooks() ([]*domain.Book, error)
}
