package main

import (
	"fmt"
	"library/internal/domain"
)

func main() {
	book := domain.Book{ID: 20, Title: "Война и мир"}
	book2 := domain.Book{ID: 21, Title: "Война и мир2"}
	book3 := domain.Book{ID: 22, Title: "Война и мир"}
	book4 := domain.Book{ID: 23, Title: "Война и мир2"}
	book5 := domain.Book{ID: 24, Title: "Война и мир"}
	book6 := domain.Book{ID: 25, Title: "Война и мир2"}
	book7 := domain.Book{ID: 26, Title: "Война и мир"}
	book8 := domain.Book{ID: 27, Title: "Война и мир2"}
	book9 := domain.Book{ID: 28, Title: "Война и мир"}
	book10 := domain.Book{ID: 29, Title: "Война и мир2"}
	book11 := domain.Book{ID: 30, Title: "Война и мир"}
	book12 := domain.Book{ID: 31, Title: "Война и мир"}

	// Создать читателя
	reader := domain.Reader{ID: 1, Name: "Иванов"}

	reader.ReaderBorrowBook(book2.ID)
	reader.ReaderBorrowBook(book3.ID)
	reader.ReaderBorrowBook(book4.ID)
	reader.ReaderBorrowBook(book5.ID)
	reader.ReaderBorrowBook(book6.ID)
	reader.ReaderBorrowBook(book7.ID)
	reader.ReaderBorrowBook(book8.ID)
	reader.ReaderBorrowBook(book9.ID)
	reader.ReaderBorrowBook(book10.ID)
	reader.ReaderBorrowBook(book11.ID)
	reader.ReaderBorrowBook(book12.ID)
	err := reader.ReaderBorrowBook(book.ID)
	if err != nil {
		fmt.Println(err)
	}
	PrintReader(reader)
}

func PrintReader(r domain.Reader) {
	fmt.Println("=== ЧИТАТЕЛЬ ===")
	fmt.Println("ID:", r.ID)
	fmt.Println("Имя:", r.Name)
	fmt.Println("Книги на рукахId:", r.BooksInRentNow)
	fmt.Println("Всего взято книг:", len(r.BooksRented))
	fmt.Println("Блокировка:", r.Ban)
	fmt.Println("Даты взятия книг:")
	for i, bookID := range r.DateOfRent.BookID {
		fmt.Printf("  Книга ID %d: %s\n", bookID, r.DateOfRent.DateOfRentingBook[i].Format("02.01.2006"))
	}
	fmt.Println("================")
}
