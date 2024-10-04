package main

import (
	"effective_mobile/elements"
	"effective_mobile/handlers"
	"effective_mobile/pkg/repository"
	"effective_mobile/pkg/service"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// @title Effective Mobile API
// @version 1.0
// @description Go Music library gin framework

// @host localhost:8080
// @BasePath /

func main() {
	godotenv.Load("../config/.env")

	//Подключаемся к PostgreSQL
	db, err := repository.NewPostgresDB(repository.Config{
		//Подстваляем данные из config
		Host:     os.Getenv("db_host"),
		Port:     os.Getenv("db_port"),
		UserName: os.Getenv("db_username"),
		DBName:   os.Getenv("db_dbname"),
		SSLMode:  os.Getenv("db_sslmode"),
		Password: os.Getenv("db_password"),
	})

	//Проверяем наличие ошибки
	if err != nil {
		logrus.Fatalf("Cant connect to postgreSQL. Err: %s", err.Error())
	}
	fmt.Println("Succes connect to postgreSQL")

	//Инициализируем репозиторий(Здесь находсятся все функции, которые будут взаимодействовать с БД).
	//Функции будут запаскаться через services
	repos := repository.NewRepository(db)

	//Инициализируем сервисы(Сервисы будут использоваться для запуска, если понадобиться дополнительных функций
	//перед взаимодействием с БД и через сервисы будут запускаться функции для взаимодействия с БД).
	//Функции сервисов запускаются из Handlers
	services := service.NewService(repos)

	//Инициализируем handler(здесь мы задаём пути запроса и также указываем какие функции запустяться по данному пути)
	handlers := handlers.New(services)
	//Создаем сервер
	server := new(elements.Server)
	//Запускаем сервер
	err = server.Run(os.Getenv("port"), handlers.InitRoutes())
	if err != nil {
		logrus.Fatalf("error while runnig server: %s", err.Error())
	}
}
