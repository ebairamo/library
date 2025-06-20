package services

import (
	"errors"
	"military/internal/domain"
	"military/internal/ports/repositories"
	servicesPort "military/internal/ports/services"
)

type storeService struct {
	productRepo  repositories.ProductRepository
	employeeRepo repositories.EmployeeRepository
}

func NewStoreService(
	productRepo repositories.ProductRepository,
	employeeRepo repositories.EmployeeRepository,
) servicesPort.StoreService {
	return &storeService{
		productRepo:  productRepo,
		employeeRepo: employeeRepo,
	}
}

// OrderProduct обрабатывает заказ товара сотрудником
func (s *storeService) OrderProduct(employeeID int, productID int, amount int) error {
	// 1. Находим сотрудника
	employee, err := s.employeeRepo.GetByID(employeeID)
	if err != nil {
		return errors.New("сотрудник не найден")
	}

	// 2. Находим товар
	product, err := s.productRepo.GetByID(productID)
	if err != nil {
		return errors.New("товар не найден")
	}

	// 3. Проверяем, достаточно ли товара на складе
	if product.Quantity < amount {
		return errors.New("недостаточно товара на складе")
	}

	// 4. Считаем общую стоимость
	totalPrice := product.Price * float64(amount)

	// 5. Проверяем, хватает ли бюджета
	if !employee.CanAfford(totalPrice) {
		return errors.New("недостаточно бюджета у сотрудника")
	}

	// 6. Резервируем товар (уменьшаем количество)
	err = product.Reserve(amount)
	if err != nil {
		return err
	}

	// 7. Создаем заказ для сотрудника
	err = employee.MakeOrder(productID, totalPrice)
	if err != nil {
		// Возвращаем товар на склад в случае ошибки
		product.Quantity += amount
		return err
	}

	// 8. Сохраняем изменения в репозиториях
	err = s.productRepo.Update(product)
	if err != nil {
		// Возвращаем деньги сотруднику в случае ошибки
		employee.Budget += totalPrice
		employee.OrderHistory = employee.OrderHistory[:len(employee.OrderHistory)-1]
		return errors.New("ошибка обновления товара")
	}

	err = s.employeeRepo.Update(employee)
	if err != nil {
		// Возвращаем товар на склад в случае ошибки
		product.Quantity += amount
		s.productRepo.Update(product)
		return errors.New("ошибка обновления данных сотрудника")
	}

	return nil
}

func (s *storeService) GetAvailableProducts() ([]*domain.Product, error) {
	allAvailableProducts, err := s.productRepo.List()
	if err != nil {
		return nil, errors.New("ошибка получения доступных продуктов")
	}
	return allAvailableProducts, nil
}

func (s *storeService) GetEmployeeOrders(employeeID int) ([]int, error) {
	employee, err := s.employeeRepo.GetByID(employeeID)
	if err != nil {
		return nil, errors.New("Ошибка получения айди сотрудника")
	}
	return employee.OrderHistory, nil
}
