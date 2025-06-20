package repositories

import "military/internal/domain"

type EmployeeRepository interface {
	GetByID(id int) (*domain.Employee, error)
	Update(*domain.Employee) error
}
