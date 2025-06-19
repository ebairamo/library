package repositories

import "library/internal/domain"

type BookRepository interface {
	GetByID(id int) (*domain.Book, error)
	Create(book *domain.Book) error
	Update(book *domain.Book) error
	Delete(id int) error
	List() ([]*domain.Book, error)
	FindByCriteria(criteria string) ([]*domain.Book, error)
}
