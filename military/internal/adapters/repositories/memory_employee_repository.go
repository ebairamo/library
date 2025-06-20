package repositories

import (
	"errors"
	"military/internal/domain"
)

type MemoryEmployeeRepository struct {
	employees map[int]*domain.Employee
}

func NewMemoryEmployeeRepository() *MemoryEmployeeRepository {
	return &MemoryEmployeeRepository{
		employees: map[int]*domain.Employee{
			1: {ID: 1, Name: "Петров", Rank: "Майор", Budget: 5000.0, OrderHistory: []int{}},
			2: {ID: 2, Name: "Сидоров", Rank: "Капитан", Budget: 3000.0, OrderHistory: []int{}},
			3: {ID: 3, Name: "Иванов", Rank: "Лейтенант", Budget: 2000.0, OrderHistory: []int{}},
		},
	}
}

func (r *MemoryEmployeeRepository) GetByID(id int) (*domain.Employee, error) {
	employee, exists := r.employees[id]
	if !exists {
		return nil, errors.New("сотрудник не найден")
	}
	return employee, nil
}

func (r *MemoryEmployeeRepository) Update(employee *domain.Employee) error {
	if employee == nil {
		return errors.New("сотрудник не может быть nil")
	}
	if _, exists := r.employees[employee.ID]; !exists {
		return errors.New("сотрудник не найден")
	}
	r.employees[employee.ID] = employee
	return nil
}
