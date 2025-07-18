package concurrent

import (
	"fmt"
	"sync"
	"user_manager/models"
)

// Другие импорты по необходимости

// ProcessUsers обрабатывает пользователей параллельно и возвращает результаты
// operation - функция, которая применяется к каждому пользователю
func ProcessUsers(users []models.User, operation func(models.User) string) []string {
	var results []string
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(users))
	for _, user := range users {
		go func(u models.User) {
			defer wg.Done()
			ch <- operation(u)
		}(user)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for value := range ch {
		results = append(results, value)
	}
	return results
}

// FindUsersByAgeRange параллельно ищет пользователей в диапазоне возрастов
func FindUsersByAgeRange(users []models.User, minAge, maxAge int) []models.User {
	var results []models.User

	if minAge > maxAge {
		fmt.Println("минимальный возвраст выше максимального")
		return nil
	}
	ch := make(chan models.User)
	var wg sync.WaitGroup
	wg.Add(len(users))
	for _, user := range users {
		go func(u models.User) {
			defer wg.Done()
			if minAge <= u.Age && u.Age <= maxAge {
				ch <- u
			}
		}(user)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for value := range ch {
		results = append(results, value)
	}
	// Ваш код здесь
	return results
}
