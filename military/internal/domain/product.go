package domain

import "errors"

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

func (p *Product) Reserve(amount int) error {
	if p.Quantity < amount {
		return errors.New("недостаточно товара на складе")
	}
	p.Quantity -= amount
	return nil
}
