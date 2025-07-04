package repositories

import (
	"errors"
	"library/internal/domain"
	"strings"
)

type MemoryBookRepository struct {
	books  map[int]*domain.Book
	nextID int
}

func NewMemoryBookRepository() *MemoryBookRepository {
	books := make(map[int]*domain.Book)

	// Добавляем тестовые книги
	books[1] = &domain.Book{ID: 1, Title: "Война и мир", Author: "Лев Толстой", Status: domain.StatusAvailable, Price: 1000.0, Year: 1869}
	books[2] = &domain.Book{ID: 2, Title: "Преступление и наказание", Author: "Федор Достоевский", Status: domain.StatusAvailable, Price: 850.0, Year: 1866}
	books[3] = &domain.Book{ID: 3, Title: "Мастер и Маргарита", Author: "Михаил Булгаков", Status: domain.StatusBorrowed, Price: 950.0, Year: 1967}
	books[4] = &domain.Book{ID: 4, Title: "Евгений Онегин", Author: "Александр Пушкин", Status: domain.StatusAvailable, Price: 750.0, Year: 1833}
	books[5] = &domain.Book{ID: 5, Title: "Анна Каренина", Author: "Лев Толстой", Status: domain.StatusLost, Price: 900.0, Year: 1877}

	return &MemoryBookRepository{
		books:  books,
		nextID: 6,
	}
}

func (r *MemoryBookRepository) GetByID(id int) (*domain.Book, error) {
	book, exists := r.books[id]
	if !exists {
		return nil, errors.New("книга не найдена")
	}
	return book, nil
}

func (r *MemoryBookRepository) Create(book *domain.Book) error {
	if book == nil {
		return errors.New("книга не может быть пустой")
	}

	if book.ID == 0 {
		book.ID = r.nextID
		r.nextID++
	} else if _, exists := r.books[book.ID]; exists {
		return errors.New("книга с таким ID уже существует")
	}

	r.books[book.ID] = book
	return nil
}

func (r *MemoryBookRepository) Update(book *domain.Book) error {
	if book == nil {
		return errors.New("книга не может быть пустой")
	}

	if _, exists := r.books[book.ID]; !exists {
		return errors.New("книга не найдена")
	}

	r.books[book.ID] = book
	return nil
}

func (r *MemoryBookRepository) Delete(id int) error {
	if _, exists := r.books[id]; !exists {
		return errors.New("книга не найдена")
	}

	delete(r.books, id)
	return nil
}

func (r *MemoryBookRepository) List() ([]*domain.Book, error) {
	books := make([]*domain.Book, 0, len(r.books))
	for _, book := range r.books {
		books = append(books, book)
	}
	return books, nil
}

func (r *MemoryBookRepository) FindByCriteria(criteria string) ([]*domain.Book, error) {
	result := make([]*domain.Book, 0)

	// Проверяем статус
	if criteria == string(domain.StatusAvailable) {
		for _, book := range r.books {
			if book.Status == domain.StatusAvailable {
				result = append(result, book)
			}
		}
		return result, nil
	}

	// Поиск по автору или названию
	for _, book := range r.books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(criteria)) ||
			strings.Contains(strings.ToLower(book.Author), strings.ToLower(criteria)) {
			result = append(result, book)
		}
	}

	return result, nil
}
