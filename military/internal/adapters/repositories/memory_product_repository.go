package repositories

import (
	"errors"
	"military/internal/domain"
)

type MemoryProductRepository struct {
	products map[int]*domain.Product
}

func NewMemoryProductRepository() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: map[int]*domain.Product{
			1: {ID: 1, Name: "Армейский рюкзак", Price: 1500.0, Quantity: 10},
			2: {ID: 2, Name: "Армейский нож", Price: 800.0, Quantity: 15},
			3: {ID: 3, Name: "Полевая фляга", Price: 300.0, Quantity: 30},
			4: {ID: 4, Name: "Тактические перчатки", Price: 450.0, Quantity: 20},
			5: {ID: 5, Name: "Маскировочный костюм", Price: 3000.0, Quantity: 5},
		},
	}
}

func (r *MemoryProductRepository) GetByID(id int) (*domain.Product, error) {
	product, exists := r.products[id]
	if !exists {
		return nil, errors.New("товар не найден")
	}
	return product, nil
}

func (r *MemoryProductRepository) Update(product *domain.Product) error {
	if product == nil {
		return errors.New("товар не может быть nil")
	}
	if _, exists := r.products[product.ID]; !exists {
		return errors.New("товар не найден")
	}
	r.products[product.ID] = product
	return nil
}

func (r *MemoryProductRepository) List() ([]*domain.Product, error) {
	products := make([]*domain.Product, 0, len(r.products))
	for _, product := range r.products {
		products = append(products, product)
	}
	return products, nil
}
