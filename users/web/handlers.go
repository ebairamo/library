package web

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"user_manager/models"
	"user_manager/storage"
)

// Добавить глобальную переменную для хранения пользователей
var users []models.User

// Добавить функцию для установки списка пользователей
func SetUsers(userList []models.User) {
	users = userList
}

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "sdfsdf")
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`--- Меню управления пользователями ---
	1. Показать всех пользователей
	2. Добавить нового пользователя
	3. Найти самого молодого пользователя
	4. Сохранить пользователей в файл
	5. Загрузить пользователей из файла
	6. Параллельная обработка пользователей
	7. Параллельная обработка возраста пользователей
	0. Выход`))
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	// Использовать глобальную переменную users вместо вызова ui.InitDefaultUsers()
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		// Логируем ошибку для внутренней диагностики
		log.Printf("Ошибка при кодировании JSON: %v", err)

		// Отправляем клиенту сообщение об ошибке с кодом 500 (Internal Server Error)
		http.Error(w, "Ошибка при формировании ответа", http.StatusInternalServerError)
		return
	}

}
func findYoungestUser(w http.ResponseWriter, r *http.Request) {
	var younger models.User
	if len(users) == 0 {
		return
	}

	ageofUser := users[0].Age
	younger = users[0]
	for _, user := range users[1:] {
		if user.Age < ageofUser {
			ageofUser = user.Age
			younger = user
		}
	}
	err := json.NewEncoder(w).Encode(younger)
	if err != nil {
		// Логируем ошибку для внутренней диагностики
		log.Printf("Ошибка при кодировании JSON: %v", err)

		// Отправляем клиенту сообщение об ошибке с кодом 500 (Internal Server Error)
		http.Error(w, "Ошибка при формировании ответа", http.StatusInternalServerError)
		return
	}
}
func saveUser(w http.ResponseWriter, r *http.Request) {
	body := rBodyToString(w, r)

	err := json.NewEncoder(w).Encode(string(body))
	if err != nil {
		// Логируем ошибку для внутренней диагностики
		log.Printf("Ошибка при кодировании JSON: %v", err)

		// Отправляем клиенту сообщение об ошибке с кодом 500 (Internal Server Error)
		http.Error(w, "Ошибка при формировании ответа", http.StatusInternalServerError)
		return
	}
	err = storage.SaveUsersToFile(users, body)
	if err != nil {
		// Логируем ошибку для внутренней диагностики
		log.Printf("Ошибка при кодировании JSON: %v", err)

		// Отправляем клиенту сообщение об ошибке с кодом 500 (Internal Server Error)
		http.Error(w, "Ошибка при формировании ответа", http.StatusInternalServerError)
		return
	}

}

func loadUser(w http.ResponseWriter, r *http.Request) {
	body := rBodyToString(w, r)

	err := json.NewEncoder(w).Encode(string(body))
	if err != nil {
		// Логируем ошибку для внутренней диагностики
		log.Printf("Ошибка при кодировании JSON: %v", err)

		// Отправляем клиенту сообщение об ошибке с кодом 500 (Internal Server Error)
		http.Error(w, "Ошибка при формировании ответа", http.StatusInternalServerError)
		return
	}
	loadedUsers, err := storage.LoadUsersFromFile(body)
	if err != nil {
		log.Printf("Ошибка в Загрузке пользователей: %v", err)
		http.Error(w, "Ошибка при формировании ответа", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(loadedUsers)
	if err != nil {
		// Логируем ошибку для внутренней диагностики
		log.Printf("Ошибка при кодировании JSON: %v", err)

		// Отправляем клиенту сообщение об ошибке с кодом 500 (Internal Server Error)
		http.Error(w, "Ошибка при формировании ответа", http.StatusInternalServerError)
		return
	}
}

func rBodyToString(w http.ResponseWriter, r *http.Request) string {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return ""
	}
	defer r.Body.Close() // Важно закрыть Body после чтения
	bodyString := string(body)
	return bodyString
}
