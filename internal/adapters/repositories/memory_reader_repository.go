package repositories

import (
	"errors"
	"library/internal/domain"
	"strings"
	"time"
)

type MemoryReaderRepository struct {
	readers map[int]*domain.Reader
	nextID  int
}

func NewMemoryReaderRepository() *MemoryReaderRepository {
	readers := make(map[int]*domain.Reader)

	// Добавляем тестовых читателей
	readers[1] = &domain.Reader{
		ID:             1,
		Name:           "Иванов Иван",
		BooksInRentNow: []int{3}, // У него уже взята книга с ID 3
		BooksRented:    []int{3},
		Ban:            false,
		DateOfRent: domain.DateOfRenting{
			DateOfRentingBook: []time.Time{time.Now().Add(-7 * 24 * time.Hour)}, // Неделю назад
			BookID:            []int{3},
		},
	}

	readers[2] = &domain.Reader{
		ID:             2,
		Name:           "Петров Петр",
		BooksInRentNow: []int{},
		BooksRented:    []int{},
		Ban:            false,
		DateOfRent: domain.DateOfRenting{
			DateOfRentingBook: []time.Time{},
			BookID:            []int{},
		},
	}

	readers[3] = &domain.Reader{
		ID:             3,
		Name:           "Сидоров Сидор",
		BooksInRentNow: []int{},
		BooksRented:    []int{5}, // Ранее брал книгу 5 (она потеряна)
		Ban:            true,     // Заблокирован за потерю книги
		DateOfRent: domain.DateOfRenting{
			DateOfRentingBook: []time.Time{},
			BookID:            []int{},
		},
	}

	return &MemoryReaderRepository{
		readers: readers,
		nextID:  4,
	}
}

func (r *MemoryReaderRepository) GetByID(id int) (*domain.Reader, error) {
	reader, exists := r.readers[id]
	if !exists {
		return nil, errors.New("читатель не найден")
	}
	return reader, nil
}

func (r *MemoryReaderRepository) Create(reader *domain.Reader) error {
	if reader == nil {
		return errors.New("читатель не может быть пустым")
	}

	if reader.ID == 0 {
		reader.ID = r.nextID
		r.nextID++
	} else if _, exists := r.readers[reader.ID]; exists {
		return errors.New("читатель с таким ID уже существует")
	}

	r.readers[reader.ID] = reader
	return nil
}

func (r *MemoryReaderRepository) Update(reader *domain.Reader) error {
	if reader == nil {
		return errors.New("читатель не может быть пустым")
	}

	if _, exists := r.readers[reader.ID]; !exists {
		return errors.New("читатель не найден")
	}

	r.readers[reader.ID] = reader
	return nil
}

func (r *MemoryReaderRepository) Delete(id int) error {
	if _, exists := r.readers[id]; !exists {
		return errors.New("читатель не найден")
	}

	delete(r.readers, id)
	return nil
}

func (r *MemoryReaderRepository) List() ([]*domain.Reader, error) {
	readers := make([]*domain.Reader, 0, len(r.readers))
	for _, reader := range r.readers {
		readers = append(readers, reader)
	}
	return readers, nil
}

func (r *MemoryReaderRepository) FindByCriteria(criteria string) ([]*domain.Reader, error) {
	result := make([]*domain.Reader, 0)

	for _, reader := range r.readers {
		if strings.Contains(strings.ToLower(reader.Name), strings.ToLower(criteria)) {
			result = append(result, reader)
		}
	}

	return result, nil
}
