package repositories

import "military/internal/domain"

type ProductRepository interface {
	GetByID(id int) (*domain.Product, error)
	Update(*domain.Product) error
	List() ([]*domain.Product, error)
}
