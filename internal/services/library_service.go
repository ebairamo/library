package services

import (
	"errors"
	"library/internal/domain"
	"library/internal/ports/repositories"
)

type libraryService struct {
	bookRepo   repositories.BookRepository
	readerRepo repositories.ReaderRepository
}

func NewLibraryService(
	bookRepo repositories.BookRepository,
	readerRepo repositories.ReaderRepository,
) *libraryService {
	return &libraryService{
		bookRepo:   bookRepo,
		readerRepo: readerRepo,
	}
}

func (s *libraryService) BorrowBook(readerID int, bookID int) error {
	// 1. Находим читателя
	reader, err := s.readerRepo.GetByID(readerID)
	if err != nil {
		return errors.New("читатель не найден")
	}

	// 2. Находим книгу
	book, err := s.bookRepo.GetByID(bookID)
	if err != nil {
		return errors.New("книга не найдена")
	}

	// 3. Проверяем, доступна ли книга
	if book.Status != domain.StatusAvailable {
		return errors.New("книга недоступна для выдачи")
	}

	// 4. Выдаем книгу читателю
	err = reader.ReaderBorrowBook(bookID)
	if err != nil {
		return err
	}

	// 5. Меняем статус книги
	err = book.Borrow()
	if err != nil {
		return err
	}

	// 6. Сохраняем изменения
	err = s.bookRepo.Update(book)
	if err != nil {
		return errors.New("ошибка обновления книги")
	}

	err = s.readerRepo.Update(reader)
	if err != nil {
		return errors.New("ошибка обновления данных читателя")
	}

	return nil
}

func (s *libraryService) ReturnBook(readerID int, bookID int) error {
	reader, err := s.readerRepo.GetByID(readerID)
	if err != nil {
		return errors.New("читатель не найден")
	}
	book, err := s.bookRepo.GetByID(bookID)
	if err != nil {
		return errors.New("книга не найдена")
	}
	exist := false
	for _, theBook := range reader.BooksInRentNow {
		if theBook == bookID {
			exist = true
		}
	}
	if exist != true {
		return errors.New("такой книги нет у читателя")
	}
	newBooksInRent := make([]int, 0, len(reader.BooksInRentNow)-1)
	for _, id := range reader.BooksInRentNow {
		if id != bookID {
			newBooksInRent = append(newBooksInRent, id)
		}
	}
	reader.BooksInRentNow = newBooksInRent

	var bookIndex int = -1
	for i, id := range reader.DateOfRent.BookID {
		if id == bookID {
			bookIndex = i
			break
		}
	}
	if bookIndex >= 0 {

		reader.DateOfRent.BookID = append(reader.DateOfRent.BookID[:bookIndex], reader.DateOfRent.BookID[bookIndex+1:]...)
		reader.DateOfRent.DateOfRentingBook = append(reader.DateOfRent.DateOfRentingBook[:bookIndex], reader.DateOfRent.DateOfRentingBook[bookIndex+1:]...)
	}
	err = book.ReturnBook()
	if err != nil {
		return err
	}
	err = s.bookRepo.Update(book)
	if err != nil {
		return errors.New("ошибка обнавления книги")
	}
	err = s.readerRepo.Update(reader)
	if err != nil {
		return errors.New("ошибка обнавления данных читателя")
	}

	return nil
}

func (s *libraryService) GetAvailableBooks() ([]*domain.Book, error) {
	criteria := string(domain.StatusAvailable)
	availableBooks, err := s.bookRepo.FindByCriteria(criteria)
	if err != nil {
		return nil, errors.New("нет такогь критерия")
	}
	return availableBooks, nil
}

func (s *libraryService) GetReaderBooks(readerID int) ([]*domain.Book, error) {
	reader, err := s.readerRepo.GetByID(readerID)
	if err != nil {
		return nil, errors.New("ошибка получения айди читателя")
	}

	booksOfReader := reader.BooksInRentNow
	books := make([]*domain.Book, 0, len(booksOfReader))
	for _, bookID := range booksOfReader {
		book, err := s.bookRepo.GetByID(bookID)
		if err != nil {
			return nil, errors.New("ошибка получения книги по айди")
		}
		books = append(books, book)
	}
	return books, nil
}
