package services

import "military/internal/domain"

type SupplyService interface {
	OrderSupply(supplierName string, products map[int]int) (*domain.Supply, error)
	ChangeToDelivered(supplyID int) error
	CancelSupply(supplyID int) error
	GetPendingList() ([]*domain.Supply, error)
	AddToWarehouse(supplyID int) error
}
