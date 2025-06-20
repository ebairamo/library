package services

import "military/internal/domain"

type StoreService interface {
	OrderProduct(employeeID int, productID int, amount int) error
	GetAvailableProducts() ([]*domain.Product, error)
	GetEmployeeOrders(employeeID int) ([]int, error)
}
