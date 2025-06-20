package domain

import "errors"

type Employee struct {
	ID           int
	Name         string
	Rank         string
	Budget       float64
	OrderHistory []int
}

func (e *Employee) CanAfford(price float64) bool {
	return e.Budget >= price
}

func (e *Employee) MakeOrder(productID int, price float64) error {
	if !e.CanAfford(price) {
		return errors.New("недостаточно бюджета")
	}
	e.Budget -= price
	e.OrderHistory = append(e.OrderHistory, productID)
	return nil
}
