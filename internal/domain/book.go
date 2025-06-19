package domain

import "fmt"

type BookStatus string

const (
	StatusAvailable BookStatus = "available"
	StatusBorrowed  BookStatus = "borrowed"
	StatusLost      BookStatus = "lost"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Price  float64
	Status BookStatus
	Year   int
}

func (b *Book) Borrow() error {
	if b == nil {
		return fmt.Errorf("Книга не существует")
	}

	if b.Title == "" {
		return fmt.Errorf("У книги нет названия")
	}
	if b.Price < 0 {
		return fmt.Errorf("Цена указанна не корректно или ее нет")
	}
	switch b.Status {
	case StatusAvailable:
		b.Status = StatusBorrowed
		return nil
	case StatusBorrowed:
		return fmt.Errorf("Ошибка выдачи, Книга уже выдана")
	case StatusLost:
		return fmt.Errorf("Ошибка выдачи, Книга потеряна")
	default:
		return fmt.Errorf("Ошибка, неизвестный статус: %s", b.Status)
	}
}

func (b *Book) ReturnBook() error {
	if b == nil {
		return fmt.Errorf("Книга не существует")
	}
	if b.Status == StatusAvailable {
		return fmt.Errorf("Книга доступна, ошибка возврата")
	}
	if b.Status == StatusLost {
		return fmt.Errorf("Книга потеряна, невозможно вернуть, обратитесь к библиотекарю")
	}
	if b.Status == StatusBorrowed {
		b.Status = StatusAvailable
		return nil
	}
	return fmt.Errorf("Что то не так с возвратом книги")
}

func (b *Book) LostBook() error {
	if b == nil {
		return fmt.Errorf("Книга не существует")
	}
	if b.Status == StatusAvailable {
		return fmt.Errorf("Книга на полке, он не может быть потеряна")
	}
	if b.Status == StatusLost {
		return fmt.Errorf("Книга уже числится потеряной")
	}
	if b.Status == StatusBorrowed {
		b.Status = StatusLost
		return nil
	}
	return fmt.Errorf("Ошибка, неизвестный статус: %s", b.Status)
}
