package main

import (
	"fmt"
	"military/internal/adapters/repositories"
	"military/internal/core/services"
	"military/internal/domain"
)

// Функция для вывода информации о поставке
func printSupply(supply *domain.Supply) {
	fmt.Printf("ID: %d\n", supply.ID)
	fmt.Printf("Поставщик: %s\n", supply.SupplierName)
	fmt.Printf("Статус: %s\n", supply.Status)
	fmt.Printf("Дата заказа: %s\n", supply.OrderDate.Format("02.01.2006"))

	if !supply.DeliveryDate.IsZero() {
		fmt.Printf("Дата доставки: %s\n", supply.DeliveryDate.Format("02.01.2006"))
	} else {
		fmt.Printf("Дата доставки: не доставлено\n")
	}

	fmt.Println("Товары:")
	for productID, quantity := range supply.Products {
		fmt.Printf("  - Товар ID: %d, Количество: %d\n", productID, quantity)
	}
	fmt.Println("----------------------")
}

// Функция для вывода информации о товаре
func printProduct(product *domain.Product) {
	fmt.Printf("ID: %d, Название: %s, Цена: %.2f руб., Количество: %d\n",
		product.ID, product.Name, product.Price, product.Quantity)
}

func main() {
	fmt.Println("ЗАПУСК СИСТЕМЫ ВОЕНТОРГ - МОДУЛЬ ПОСТАВОК")
	fmt.Println("=======================================")

	// Инициализация репозиториев
	productRepo := repositories.NewMemoryProductRepository()
	supplyRepo := repositories.NewMemorySupplyRepository()

	// Создание сервиса поставок
	supplyService := services.NewSupplyService(supplyRepo, productRepo)

	// 1. Получение списка всех поставок
	fmt.Println("\n=== СПИСОК ВСЕХ ПОСТАВОК ===")
	supplies, err := supplyRepo.List()
	if err != nil {
		fmt.Println("ОШИБКА получения списка поставок:", err)
		return
	}

	for _, supply := range supplies {
		printSupply(supply)
	}

	// 2. Получение списка ожидаемых поставок
	fmt.Println("\n=== ОЖИДАЕМЫЕ ПОСТАВКИ ===")
	pendingSupplies, err := supplyService.GetPendingList()
	if err != nil {
		fmt.Println("ОШИБКА получения списка ожидаемых поставок:", err)
		return
	}

	for _, supply := range pendingSupplies {
		printSupply(supply)
	}

	// 3. Вывод текущего состояния склада
	fmt.Println("\n=== ТЕКУЩЕЕ СОСТОЯНИЕ СКЛАДА ===")
	products, err := productRepo.List()
	if err != nil {
		fmt.Println("ОШИБКА получения списка товаров:", err)
		return
	}

	for _, product := range products {
		printProduct(product)
	}

	// 4. Создание новой поставки
	fmt.Println("\n=== СОЗДАНИЕ НОВОЙ ПОСТАВКИ ===")
	newSupplyProducts := map[int]int{
		1: 3, // 3 единицы товара с ID 1
		2: 5, // 5 единиц товара с ID 2
	}

	newSupply, err := supplyService.OrderSupply("КрасноармейскийСнаб", newSupplyProducts)
	if err != nil {
		fmt.Println("ОШИБКА создания поставки:", err)
		return
	}

	fmt.Println("Создана новая поставка:")
	printSupply(newSupply)

	// 5. Изменение статуса поставки на "доставлено"
	fmt.Println("\n=== ДОСТАВКА ПОСТАВКИ ===")
	// Доставляем первую ожидаемую поставку (должна быть с ID 1)
	supplyToDeliver := 1
	if len(pendingSupplies) > 0 {
		supplyToDeliver = pendingSupplies[0].ID
	}

	fmt.Printf("Отмечаем поставку ID %d как доставленную\n", supplyToDeliver)
	err = supplyService.ChangeToDelivered(supplyToDeliver)
	if err != nil {
		fmt.Println("ОШИБКА доставки поставки:", err)
	} else {
		fmt.Println("Поставка успешно отмечена как доставленная!")

		// Получаем обновленную информацию о поставке
		updatedSupply, err := supplyRepo.GetByID(supplyToDeliver)
		if err == nil {
			printSupply(updatedSupply)
		}
	}

	// 6. Добавление товаров на склад из доставленной поставки
	fmt.Println("\n=== ДОБАВЛЕНИЕ ТОВАРОВ НА СКЛАД ===")
	fmt.Printf("Добавляем товары из поставки ID %d на склад\n", supplyToDeliver)
	err = supplyService.AddToWarehouse(supplyToDeliver)
	if err != nil {
		fmt.Println("ОШИБКА добавления товаров на склад:", err)
	} else {
		fmt.Println("Товары успешно добавлены на склад!")
	}

	// 7. Проверка обновленного состояния склада
	fmt.Println("\n=== ОБНОВЛЕННОЕ СОСТОЯНИЕ СКЛАДА ===")
	updatedProducts, err := productRepo.List()
	if err != nil {
		fmt.Println("ОШИБКА получения списка товаров:", err)
		return
	}

	for _, product := range updatedProducts {
		printProduct(product)
	}

	// 8. Отмена новой поставки
	fmt.Println("\n=== ОТМЕНА ПОСТАВКИ ===")
	fmt.Printf("Отменяем поставку ID %d\n", newSupply.ID)
	err = supplyService.CancelSupply(newSupply.ID)
	if err != nil {
		fmt.Println("ОШИБКА отмены поставки:", err)
	} else {
		fmt.Println("Поставка успешно отменена!")

		// Получаем обновленную информацию о поставке
		cancelledSupply, err := supplyRepo.GetByID(newSupply.ID)
		if err == nil {
			printSupply(cancelledSupply)
		}
	}

	// 9. Проверка отказоустойчивости - попытка некорректных операций
	fmt.Println("\n=== ПРОВЕРКА ОБРАБОТКИ ОШИБОК ===")

	// Попытка доставить отмененную поставку
	fmt.Printf("Попытка доставить отмененную поставку ID %d:\n", newSupply.ID)
	err = supplyService.ChangeToDelivered(newSupply.ID)
	fmt.Println("Результат:", err)

	// Попытка добавить на склад не доставленную поставку
	fmt.Println("Попытка добавить на склад товары из отмененной поставки:")
	err = supplyService.AddToWarehouse(newSupply.ID)
	fmt.Println("Результат:", err)

	fmt.Println("\n=== ОПЕРАЦИЯ ЗАВЕРШЕНА ===")
}
