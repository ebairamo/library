package repositories

import (
	"errors"
	"military/internal/domain"
	"time"
)

type MemorySupplyRepository struct {
	supplies map[int]*domain.Supply
	nextID   int
}

func NewMemorySupplyRepository() *MemorySupplyRepository {
	supplies := make(map[int]*domain.Supply)
	supplies[1] = &domain.Supply{
		ID:           1,
		SupplierName: "ВоенТоргСнаб",
		Products: map[int]int{
			1: 5,  // 5 единиц товара с ID 1
			3: 10, // 10 единиц товара с ID 3
		},
		Status:       domain.StatusExpected,           // или StatusExpected, в зависимости от вашей модели
		OrderDate:    time.Now().Add(-48 * time.Hour), // 2 дня назад
		DeliveryDate: time.Time{},                     // Пустая дата
	}

	supplies[2] = &domain.Supply{
		ID:           2,
		SupplierName: "АрмТехПоставка",
		Products: map[int]int{
			2: 3, // 3 единицы товара с ID 2
			4: 7, // 7 единиц товара с ID 4
		},
		Status:       domain.StatusDelivered,
		OrderDate:    time.Now().Add(-72 * time.Hour), // 3 дня назад
		DeliveryDate: time.Now().Add(-24 * time.Hour), // 1 день назад
	}

	supplies[3] = &domain.Supply{
		ID:           3,
		SupplierName: "ТактикСнаб",
		Products: map[int]int{
			5: 2, // 2 единицы товара с ID 5
		},
		Status:       domain.StatusCancelled,
		OrderDate:    time.Now().Add(-96 * time.Hour), // 4 дня назад
		DeliveryDate: time.Time{},                     // Пустая дата
	}
	return &MemorySupplyRepository{
		supplies: supplies,
		nextID:   4,
	}
}

func (r *MemorySupplyRepository) GetByID(id int) (*domain.Supply, error) {
	supply, exists := r.supplies[id]
	if !exists {
		return nil, errors.New("Id не существует")
	}
	return supply, nil
}
