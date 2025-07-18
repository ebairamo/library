package concurrent

import (
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
