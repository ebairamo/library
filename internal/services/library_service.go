// файл: internal/core/services/order_service.go

package services

import (
	"errors"
	"shop/internal/domain"
	"shop/internal/ports/repositories"
	"time"
)

type orderService struct {
	orderRepo    repositories.OrderRepository
	productRepo  repositories.ProductRepository
	customerRepo repositories.CustomerRepository
}

func NewOrderService(
	orderRepo repositories.OrderRepository,
	productRepo repositories.ProductRepository,
	customerRepo repositories.CustomerRepository,
) *orderService {
	return &orderService{
		orderRepo:    orderRepo,
		productRepo:  productRepo,
		customerRepo: customerRepo,
	}
}

func (s *orderService) CreateOrder(customerID int, productIDs []int) (*domain.Order, error) {
	customer, err := s.customerRepo.GetByID(customerID)
	if err != nil {
		return nil, errors.New("покупатель не найден")
	}

	if customer.IsBlocked {
		return nil, errors.New("покупатель заблокирован")
	}

	var products []*domain.Product
	var totalPrice float64 = 0

	for _, productID := range productIDs {
		product, err := s.productRepo.GetByID(productID)
		if err != nil {
			return nil, errors.New("товар не найден")
		}

		if product.Quantity <= 0 {
			return nil, errors.New("товар отсутствует на складе")
		}

		product.Quantity--
		products = append(products, product)
		totalPrice += product.Price

		err = s.productRepo.Update(product)
		if err != nil {
			return nil, errors.New("ошибка обновления товара")
		}
	}

	order := &domain.Order{
		CustomerID: customerID,
		Products:   products,
		TotalPrice: totalPrice,
		Status:     domain.StatusNew,
		CreatedAt:  time.Now(),
	}

	err = s.orderRepo.Create(order)
	if err != nil {
		return nil, errors.New("ошибка сохранения заказа")
	}

	customer.OrderCount++
	err = s.customerRepo.Update(customer)
	if err != nil {
		return nil, errors.New("ошибка обновления данных покупателя")
	}

	return order, nil
}

func (s *orderService) CancelOrder(orderID int) error {
	order, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return errors.New("заказ не найден")
	}

	if order.Status != domain.StatusNew {
		return errors.New("нельзя отменить заказ в текущем статусе")
	}

	order.Status = domain.StatusCancelled
	order.UpdatedAt = time.Now()

	for _, product := range order.Products {
		product.Quantity++
		err = s.productRepo.Update(product)
		if err != nil {
			return errors.New("ошибка обновления товара")
		}
	}

	err = s.orderRepo.Update(order)
	if err != nil {
		return errors.New("ошибка обновления заказа")
	}

	return nil
}

func (s *orderService) GetCustomerOrders(customerID int) ([]*domain.Order, error) {
	_, err := s.customerRepo.GetByID(customerID)
	if err != nil {
		return nil, errors.New("покупатель не найден")
	}

	orders, err := s.orderRepo.FindByCustomerID(customerID)
	if err != nil {
		return nil, errors.New("ошибка получения заказов")
	}

	return orders, nil
}
