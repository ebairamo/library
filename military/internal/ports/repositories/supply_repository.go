package repositories

import "military/internal/domain"

type SupplyRepository interface {
	GetByID(id int) (*domain.Supply, error)
	Create(*domain.Supply) error
	Update(*domain.Supply) error
	List() ([]*domain.Supply, error)
	FindPending() ([]*domain.Supply, error)
}
