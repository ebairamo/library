package models

import "fmt"

// User представляет пользователя системы
type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
}

// SayHello - метод приветствия
func (u User) SayHello() {
	fmt.Printf("Привет! Меня зовут %s, мне %d лет.\n", u.Name, u.Age)
}

// FullName возвращает полное имя пользователя
func (u User) FullName() string {
	return u.Surname + " " + u.Name
}
