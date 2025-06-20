package domain

import (
	"errors"
	"time"
)

type SupplyStatus string

const (
	StatusExpected  SupplyStatus = "Expected"
	StatusDelivered SupplyStatus = "Delivered"
	StatusCancelled SupplyStatus = "Cancelled"
)

type Supply struct {
	ID           int
	SupplierName string
	Products     map[int]int
	Status       SupplyStatus
	OrderDate    time.Time
	DeliveryDate time.Time
}

func (s *Supply) Deliver() error {
	if s.Status == StatusCancelled || s.Status == StatusDelivered {
		return errors.New("Невозможно изменить статус на доставлено")
	}
	s.Status = StatusDelivered
	s.DeliveryDate = time.Now()
	return nil
}

func (s *Supply) Cancel() error {
	if s.Status == StatusCancelled || s.Status == StatusDelivered {
		return errors.New("Невозможно изменить статус на отменено")
	}
	s.Status = StatusCancelled
	return nil
}
