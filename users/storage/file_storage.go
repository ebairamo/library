package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"user_manager/models"
)

// SaveUsersToFile сохраняет список пользователей в файл
func SaveUsersToFile(users []models.User, filename string) error {
	if len(users) == 0 {
		return fmt.Errorf("список пользователей пуст")
	}
	jsonData, err := json.Marshal(users)
	if err != nil {
		return fmt.Errorf("ошибка при сериализации: %w", err)
	}
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("ошибка при записи файла: %w", err)
	}
	fmt.Println("Операция выполнена успешно. Пользователи сохранены в файл:", filename)
	return nil
}

// LoadUsersFromFile загружает список пользователей из файла
func LoadUsersFromFile(filename string) ([]models.User, error) {
	var users []models.User
	if filename == "" {
		return nil, fmt.Errorf("имя файла не указано")
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	err = json.Unmarshal(file, &users)
	if err != nil {
		return nil, fmt.Errorf("ошибка преобразования JSON в структуру: %w", err)
	}

	fmt.Println("Операция выполнена успешно. Данные загружены из файла:", filename)
	return users, nil
}
