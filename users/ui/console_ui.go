package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"user_manager/concurrent"
	"user_manager/models"
	"user_manager/storage"
)

// RunConsoleUI запускает консольный интерфейс управления пользователями
func RunConsoleUI() {
	users := initDefaultUsers()
	var choice int

	for {
		displayMenu()
		fmt.Print("\nВыберите действие: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			DisplayAllUsers(users)
		case 2:
			newUser, err := addNewUser()
			if err != nil {
				break
			}

			users = append(users, newUser)

		case 3:
			youngerUser, err := findYoungestUser(users)
			if err != nil {
				fmt.Println("Ошибка", err)
				break
			}
			fmt.Printf("Самый младший пользователь %s, Ему %d лет", youngerUser.Name, youngerUser.Age)
		case 4:
			fileName := getFileName()
			storage.SaveUsersToFile(users, fileName)
		case 5:
			fileName := getFileName()
			loadedUsers, err := storage.LoadUsersFromFile(fileName)
			if err != nil {
				break
			}
			if loadedUsers != nil {
				users = loadedUsers
			}
		case 6:
			fmt.Println("Запуск параллельной обработки пользователей...")
			results := concurrent.ProcessUsers(users, func(u models.User) string {
				return fmt.Sprintf("Обработан пользователь: %s %s", u.Name, u.Surname)
			})
			fmt.Println("\nРезультаты обработки:")
			for _, result := range results {
				fmt.Println(result)
			}
		case 7:
			var minAge, maxAge int
			fmt.Print("Введите минимальный возраст: ")
			fmt.Scan(&minAge)
			fmt.Print("Введите максимальный возраст: ")
			fmt.Scan(&maxAge)

			fmt.Printf("Поиск пользователей от %d до %d лет...\n", minAge, maxAge)
			matchedUsers := concurrent.FindUsersByAgeRange(users, minAge, maxAge)

			fmt.Printf("\nНайдено %d пользователей:\n", len(matchedUsers))
			for i, user := range matchedUsers {
				fmt.Printf("%d. %s %s (%d лет)\n", i+1, user.Name, user.Surname, user.Age)
			}
		case 0:
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

// Вспомогательные функции для консольного интерфейса
func displayMenu() {
	fmt.Println("\n--- Меню управления пользователями ---")
	fmt.Println("1. Показать всех пользователей")
	fmt.Println("2. Добавить нового пользователя")
	fmt.Println("3. Найти самого молодого пользователя")
	fmt.Println("4. Сохранить пользователей в файл")
	fmt.Println("5. Загрузить пользователей из файла")
	fmt.Println("6. Параллельная обработка пользователей")
	fmt.Println("7. Параллельная обработка возраста пользователей")
	fmt.Println("0. Выход")
}

func findYoungestUser(users []models.User) (models.User, error) {
	var younger models.User
	if len(users) == 0 {
		return younger, fmt.Errorf("список пустой")
	}

	ageofUser := users[0].Age
	younger = users[0]
	for _, user := range users[1:] {
		if user.Age < ageofUser {
			ageofUser = user.Age
			younger = user
		}
	}
	return younger, nil
}

func addNewUser() (models.User, error) {
	var userString string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите данные о пользователе через пробел, Пример: Александр 25 eee@mail.ru Пушкин")

	// Сброс буфера ввода
	reader.ReadString('\n')

	// Ожидание ввода пользователя
	userString, _ = reader.ReadString('\n')
	userString = strings.TrimSpace(userString)

	// Обработка ввода
	fmt.Println("Вы ввели:", userString)
	userSlice := strings.Split(userString, " ")

	// Проверка достаточности данных
	if len(userSlice) < 4 {
		fmt.Println("Недостаточно данных. Нужно: имя, возраст, email, фамилия")
		return models.User{}, fmt.Errorf("недостаточно данных. Нужно: имя, возраст, email, фамилия")
	}

	// Преобразование возраста
	userAge, err := strconv.Atoi(userSlice[1])
	if err != nil {
		return models.User{}, fmt.Errorf("невозможно преобразовать возраст в число")
	}

	// Создание нового пользователя
	user := models.User{
		Name:    userSlice[0],
		Surname: userSlice[3],
		Age:     userAge,
		Email:   userSlice[2],
	}

	// Добавление в список

	fmt.Println("Операция выполнена успешно. Пользователь добавлен.")
	return user, nil
}

func initDefaultUsers() []models.User {
	users := []models.User{
		{
			Name:    "Александр",
			Age:     25,
			Email:   "eee@mail.ru",
			Surname: "Пушкин",
		},
		{
			Name:    "Владимир",
			Age:     35,
			Email:   "eee@mail.ru",
			Surname: "Маяковский",
		},
		{
			Name:    "Федор",
			Age:     55,
			Email:   "eee@mail.ru",
			Surname: "Достоевский",
		},
	}
	return users
}

func DisplayAllUsers(users []models.User) {
	if len(users) == 0 {
		fmt.Println("Список пользователей пуст")
	} else {
		fmt.Println("Список пользователей:")
		for i, user := range users {
			fmt.Printf("%d. %s %s (%d лет, %s)\n",
				i+1, user.Name, user.Surname, user.Age, user.Email)
		}
	}
}

func getFileName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите название файла в формате json, Пример: lol.json")
	reader.ReadString('\n')
	fileName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода названия файла:", err)
		return ""
	}
	fileName = strings.TrimSpace(fileName)

	// Проверка формата файла
	if !strings.HasSuffix(fileName, ".json") {
		fmt.Println("Файл должен быть формата json")
		return ""
	}
	return fileName
}
