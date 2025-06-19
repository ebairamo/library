package domain

import (
	"fmt"
	"time"
)

type DateOfRenting struct {
	DateOfRentingBook []time.Time
	BookID            []int
}

type RentDates struct {
	ReaderRented []DateOfRenting
}

type Reader struct {
	ID             int
	Name           string
	BooksInRentNow []int
	BooksRented    []int
	Ban            bool
	DateOfRent     DateOfRenting
}

func (r *Reader) ReaderBorrowBook(bookID int) error {
	if r.Ban {
		return fmt.Errorf("Пользователь заблокирован")
	}
	if len(r.BooksInRentNow) >= 10 {
		return fmt.Errorf("Лимит книг привышен")
	}
	for i := 0; i < len(r.BooksInRentNow); i++ {
		if r.BooksInRentNow[i] == bookID {
			return fmt.Errorf("Книга уже арендована вами")
		}
	}
	r.BooksInRentNow = append(r.BooksInRentNow, bookID)
	r.DateOfRent.DateOfRentingBook = append(r.DateOfRent.DateOfRentingBook, time.Now())
	r.DateOfRent.BookID = append(r.DateOfRent.BookID, bookID)

	r.BooksRented = append(r.BooksRented, bookID)
	return nil
}
