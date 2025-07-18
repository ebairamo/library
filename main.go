package main

import (
	"fmt"
	"library/internal/adapters/repositories"
	"library/internal/core/services"
	"library/internal/domain"
)

// Функция для вывода информации о книге
func printBook(book *domain.Book) {
	fmt.Printf("ID: %d, Название: \"%s\", Автор: %s, Статус: %s, Цена: %.2f руб.\n",
		book.ID, book.Title, book.Author, book.Status, book.Price)
}

// Функция для вывода информации о читателе
func printReader(reader *domain.Reader) {
	fmt.Println("=== ЧИТАТЕЛЬ ===")
	fmt.Println("ID:", reader.ID)
	fmt.Println("Имя:", reader.Name)
	fmt.Println("Книги на руках:", reader.BooksInRentNow)
	fmt.Println("Всего взято книг:", len(reader.BooksRented))
	fmt.Println("Блокировка:", reader.Ban)
	fmt.Println("Даты взятия книг:")
	for i, bookID := range reader.DateOfRent.BookID {
		fmt.Printf("  Книга ID %d: %s\n", bookID, reader.DateOfRent.DateOfRentingBook[i].Format("02.01.2006"))
	}
	fmt.Println("================")
}

func main() {
	// Инициализация репозиториев
	bookRepo := repositories.NewMemoryBookRepository()
	readerRepo := repositories.NewMemoryReaderRepository()

	// Создание сервиса
	libraryService := services.NewLibraryService(bookRepo, readerRepo)

	// 1. Вывод списка всех книг
	fmt.Println("\n=== СПИСОК ВСЕХ КНИГ ===")
	books, err := bookRepo.List()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	for _, book := range books {
		printBook(book)
	}

	// 2. Вывод списка доступных книг
	fmt.Println("\n=== ДОСТУПНЫЕ КНИГИ ===")
	availableBooks, err := libraryService.GetAvailableBooks()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	for _, book := range availableBooks {
		printBook(book)
	}

	// 3. Вывод информации о читателях
	fmt.Println("\n=== ЧИТАТЕЛИ ===")
	readers, err := readerRepo.List()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	for _, reader := range readers {
		printReader(reader)
	}

	// 4. Выдача книги читателю
	fmt.Println("\n=== ВЫДАЧА КНИГИ ===")
	readerID := 2 // Петров Петр
	bookID := 1   // "Война и мир"

	fmt.Printf("Выдаем книгу ID %d читателю ID %d\n", bookID, readerID)
	err = libraryService.BorrowBook(readerID, bookID)
	if err != nil {
		fmt.Println("Ошибка выдачи книги:", err)
	} else {
		fmt.Println("Книга успешно выдана!")

		// Проверяем обновленную информацию
		reader, _ := readerRepo.GetByID(readerID)
		book, _ := bookRepo.GetByID(bookID)

		fmt.Println("\nОбновленная информация о книге:")
		printBook(book)

		fmt.Println("\nОбновленная информация о читателе:")
		printReader(reader)
	}

	// 5. Получение списка книг читателя
	fmt.Println("\n=== КНИГИ ЧИТАТЕЛЯ ===")
	readerBooks, err := libraryService.GetReaderBooks(readerID)
	if err != nil {
		fmt.Println("Ошибка получения книг читателя:", err)
	} else {
		fmt.Printf("Книги читателя ID %d:\n", readerID)
		for _, book := range readerBooks {
			printBook(book)
		}
	}

	// 6. Возврат книги
	fmt.Println("\n=== ВОЗВРАТ КНИГИ ===")
	fmt.Printf("Возвращаем книгу ID %d от читателя ID %d\n", bookID, readerID)
	err = libraryService.ReturnBook(readerID, bookID)
	if err != nil {
		fmt.Println("Ошибка возврата книги:", err)
	} else {
		fmt.Println("Книга успешно возвращена!")

		// Проверяем обновленную информацию
		reader, _ := readerRepo.GetByID(readerID)
		book, _ := bookRepo.GetByID(bookID)

		fmt.Println("\nОбновленная информация о книге:")
		printBook(book)

		fmt.Println("\nОбновленная информация о читателе:")
		printReader(reader)
	}

	// 7. Отметка книги как потерянной
	fmt.Println("\n=== ОТМЕТКА КНИГИ КАК ПОТЕРЯННОЙ ===")
	bookIDToLose := 4 // "Евгений Онегин"

	// Сначала выдаем книгу
	err = libraryService.BorrowBook(readerID, bookIDToLose)
	if err != nil {
		fmt.Println("Ошибка выдачи книги:", err)
	} else {
		fmt.Printf("Книга ID %d выдана читателю ID %d\n", bookIDToLose, readerID)

		// Затем отмечаем как потерянную
		err = libraryService.MarkBookAsLost(readerID, bookIDToLose)
		if err != nil {
			fmt.Println("Ошибка отметки книги как потерянной:", err)
		} else {
			fmt.Println("Книга успешно отмечена как потерянная!")

			// Проверяем обновленную информацию
			reader, _ := readerRepo.GetByID(readerID)
			book, _ := bookRepo.GetByID(bookIDToLose)

			fmt.Println("\nОбновленная информация о книге:")
			printBook(book)

			fmt.Println("\nОбновленная информация о читателе:")
			printReader(reader)
		}
	}

	// 8. Проверка ограничений (попытка взять книгу заблокированным пользователем)
	fmt.Println("\n=== ПРОВЕРКА ОГРАНИЧЕНИЙ ===")
	bannedReaderID := 3  // Заблокированный Сидоров
	availableBookID := 2 // Доступная книга

	fmt.Printf("Попытка выдачи книги ID %d заблокированному читателю ID %d\n", availableBookID, bannedReaderID)
	err = libraryService.BorrowBook(bannedReaderID, availableBookID)
	fmt.Println("Результат:", err)

	fmt.Println("\n=== ОПЕРАЦИЯ ЗАВЕРШЕНА ===")
}
