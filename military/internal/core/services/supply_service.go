package services

import (
	"errors"
	"math/rand/v2"
	"military/internal/domain"
	"military/internal/ports/repositories"
	"time"
)

type supplyService struct {
	supplyRepo  repositories.SupplyRepository
	productRepo repositories.ProductRepository
}

func NewSupplyService(
	supplyRepo repositories.SupplyRepository,
	productRepo repositories.ProductRepository,
) *supplyService {
	return &supplyService{
		supplyRepo:  supplyRepo,
		productRepo: productRepo,
	}
}

func (s *supplyService) OrderSupply(supplierName string, products map[int]int) (*domain.Supply, error) {
	if supplierName == "" {
		return nil, errors.New("Имя поставщика не может быть пустым")
	}
	if len(products) == 0 {
		return nil, errors.New("Список продуктов не может быть пустым")
	}

	supplyID := rand.IntN(1000)

	supply := &domain.Supply{
		ID:           supplyID,
		SupplierName: supplierName,
		Products:     products,
		Status:       domain.StatusExpected,
		OrderDate:    time.Now(),
		DeliveryDate: time.Time{},
	}
	err := s.supplyRepo.Create(supply)
	if err != nil {
		return nil, errors.New("Ошибка создания поставки")
	}

	return supply, nil
}

func (s *supplyService) ChangeToDelivered(supplyID int) error {
	supply, err := s.supplyRepo.GetByID(supplyID)
	if err != nil {
		return errors.New("Ошибка получения supplyID")
	}
	err = supply.Deliver()
	if err != nil {
		return errors.New("Ошибка смены статуса поставки")
	}
	err = s.supplyRepo.Update(supply)
	if err != nil {
		return errors.New("Ошибка обнавления стутуса поставки")
	}

	return nil
}

func (s *supplyService) CancelSupply(supplyID int) error {
	supply, err := s.supplyRepo.GetByID(supplyID)
	if err != nil {
		return errors.New("Ошибка получения supplyID")
	}
	err = supply.Cancel()
	if err != nil {
		return errors.New("Ошибка смены статуса")
	}
	err = s.supplyRepo.Update(supply)
	if err != nil {
		return errors.New("Ошибка обнавления статуса поставки")
	}
	return nil
}

func (s *supplyService) GetPendingList() ([]*domain.Supply, error) {
	pendingList, err := s.supplyRepo.FindPending()
	if err != nil {
		return nil, errors.New("Ошибка поиска листа ожидания")
	}

	return pendingList, nil
}

func (s *supplyService) AddToWarehouse(supplyID int) error {
	supply, err := s.supplyRepo.GetByID(supplyID)
	if err != nil {
		return errors.New("Ошибка получения supplyID")
	}
	if supply.Status != domain.StatusDelivered {
		return errors.New("Невозможно добавить на склад так как статус не Доставлен")
	}
	products := supply.Products
	for productID, quantity := range products {
		product, err := s.productRepo.GetByID(productID)
		if err != nil {
			return errors.New("Не удается получить продукт по айди")
		}
		product.Quantity += quantity
		err = s.productRepo.Update(product)
		if err != nil {
			return errors.New("Ошибка обнавления продукта ")
		}
	}

	return nil
}
