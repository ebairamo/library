package repositories

import "library/internal/domain"

type ReaderRepository interface {
	GetByID(id int) (*domain.Reader, error)
	Create(*domain.Reader) error
	Update(*domain.Reader) error
	Delete(id int) error
	List() ([]*domain.Reader, error)
	FindByCriteria(criteria string) ([]*domain.Reader, error)
}
