package main

import (
	"fmt"
)

// Структура книги
type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
}

// Структура даты аренды
type DateOfRenting struct {
	Day   int
	Month int
	Year  int
}

// Структура читателя
type Reader struct {
	ID             int
	Name           string
	BooksInRentNow map[int]Book
	BooksRented    map[int]Book
	Ban            bool
	DateOfRent     DateOfRenting
}

func main() {
	// Книги русских классиков
	book1 := Book{ID: 1, Title: "Преступление и наказание", Author: "Фёдор Достоевский", Year: 1866}
	book2 := Book{ID: 2, Title: "Евгений Онегин", Author: "Александр Пушкин", Year: 1833}
	book3 := Book{ID: 3, Title: "Война и мир", Author: "Лев Толстой", Year: 1869}
	book4 := Book{ID: 4, Title: "Герой нашего времени", Author: "Михаил Лермонтов", Year: 1840}
	book5 := Book{ID: 5, Title: "Мастер и Маргарита", Author: "Михаил Булгаков", Year: 1967}

	// Пример даты аренды
	date := DateOfRenting{Day: 29, Month: 6, Year: 2025}

	// Читатель с книгами
	reader := Reader{
		ID:   1,
		Name: "Иван Петров",
		BooksInRentNow: map[int]Book{
			book1.ID: book1,
			book2.ID: book2,
			book5.ID: book5,
		},
		BooksRented: map[int]Book{
			book3.ID: book3,
			book4.ID: book4,
		},
		Ban:        false,
		DateOfRent: date,
	}

	// Просто выводим текущие книги в аренде
	fmt.Println("Книги, которые сейчас у читателя на руках:")
	for id, book := range reader.BooksInRentNow {
		fmt.Printf("ID: %d, Название: \"%s\", Автор: %s\n", id, book.Title, book.Author)
	}

	id := 5

	if value, ok := reader.BooksInRentNow[id]; ok {
		fmt.Println(value, "Тревога")
	}
}
