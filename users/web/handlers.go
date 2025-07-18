package web

import (
	"fmt"
	"net/http"
	"user_manager/models"
	"user_manager/ui"
)

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
	ui.DisplayAllUsers([]models.User{})
}
