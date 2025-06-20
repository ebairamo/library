package main

import (
	"fmt"
	"military/internal/adapters/repositories"
	"military/internal/core/services"
)

func main() {
	// Создаем репозитории с тестовыми данными
	productRepo := repositories.NewMemoryProductRepository()
	employeeRepo := repositories.NewMemoryEmployeeRepository()

	// Создаем сервис
	storeService := services.NewStoreService(productRepo, employeeRepo)

	// Получаем доступные товары
	products, err := storeService.GetAvailableProducts()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("=== ДОСТУПНЫЕ ТОВАРЫ ===")
	for _, p := range products {
		fmt.Printf("ID: %d, Название: %s, Цена: %.2f, Количество: %d\n",
			p.ID, p.Name, p.Price, p.Quantity)
	}

	// Делаем заказ
	employeeID := 1 // Майор Петров
	productID := 2  // Армейский нож
	amount := 2     // 2 штуки

	fmt.Printf("\nПытаемся заказать товар ID:%d в количестве %d для сотрудника ID:%d\n",
		productID, amount, employeeID)

	err = storeService.OrderProduct(employeeID, productID, amount)
	if err != nil {
		fmt.Println("Ошибка заказа:", err)
	} else {
		fmt.Println("Заказ успешно оформлен!")
	}

	// Проверяем историю заказов
	orders, err := storeService.GetEmployeeOrders(employeeID)
	if err != nil {
		fmt.Println("Ошибка получения истории:", err)
		return
	}

	fmt.Println("\n=== ИСТОРИЯ ЗАКАЗОВ ===")
	fmt.Printf("Сотрудник ID:%d заказал товары с ID: %v\n", employeeID, orders)

	// Получаем обновленные данные о сотруднике
	employee, err := employeeRepo.GetByID(employeeID)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("\nОставшийся бюджет: %.2f\n", employee.Budget)
}
