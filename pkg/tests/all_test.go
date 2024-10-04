package tests

import (
	"effective_mobile/handlers"
	"effective_mobile/pkg/repository"
	"effective_mobile/pkg/service"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/magiconair/properties/assert"
	"github.com/sirupsen/logrus"
)

// Что бы тесты работали нужно инициализировать все сервисы как в main
func initServices() *handlers.Handlers {
	godotenv.Load("../../cmd/.env")

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
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handlers.New(services)
	return handlers
}

// ТЕСТЫ НА /create_new_song (в test_table находятся сами тесты. Здесь они запускаются)
// Тест работает только с уже запущенным на localhost сервере.
func TestHandler_create_new_song(t *testing.T) {
	//Поднимаем все сервисы
	handlers := initServices()
	router := handlers.InitRoutes()

	//В цикле проходимся по каждому тесту и запускаем его
	for _, test := range getTestTableCreateNewSong() {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/create_new_song", strings.NewReader(test.requestBody))

			w := httptest.NewRecorder()
			//Выполняется запрос
			router.ServeHTTP(w, req)

			//Проверяем ответ
			assert.Equal(t, test.expectedCode, w.Code)
		})
	}
}

// Тест на получении всех песен с пагинацией
func TestHandler_get_all_songs(t *testing.T) {
	//Поднимаем все сервисы
	handlers := initServices()
	router := handlers.InitRoutes()

	//В цикле проходимся по каждому тесту и запускаем его
	for _, test := range getTestTableGetAllSongs() {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/get_all_songs", strings.NewReader(test.requestBody))

			w := httptest.NewRecorder()
			//Выполняется запрос
			router.ServeHTTP(w, req)

			//Проверяем ответ
			assert.Equal(t, test.expectedCode, w.Code)
		})
	}
}

// Тест на получении песни
func TestHandler_info(t *testing.T) {
	//Поднимаем все сервисы
	handlers := initServices()
	router := handlers.InitRoutes()

	//В цикле проходимся по каждому тесту и запускаем его
	for _, test := range getTestTableInfo() {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/info", strings.NewReader(test.requestBody))

			w := httptest.NewRecorder()
			//Выполняется запрос
			router.ServeHTTP(w, req)

			//Проверяем ответ
			assert.Equal(t, test.expectedCode, w.Code)
		})
	}
}

// Тест на удаление песни
func TestHandler_deleteSong(t *testing.T) {
	//Поднимаем все сервисы
	handlers := initServices()
	router := handlers.InitRoutes()

	//В цикле проходимся по каждому тесту и запускаем его
	for _, test := range getTestTableDeleteSong() {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest("DELETE", "/delete_song", strings.NewReader(test.requestBody))

			w := httptest.NewRecorder()
			//Выполняется запрос
			router.ServeHTTP(w, req)

			//Проверяем ответ
			assert.Equal(t, test.expectedCode, w.Code)
		})
	}
}

// Тест на получении куплетов песни с пагинации
func TestHandler_getSongByVerses(t *testing.T) {
	//Поднимаем все сервисы
	handlers := initServices()
	router := handlers.InitRoutes()

	//В цикле проходимся по каждому тесту и запускаем его
	for _, test := range getTestTableSongByVerses() {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/get_song_by_verses", strings.NewReader(test.requestBody))

			w := httptest.NewRecorder()
			//Выполняется запрос
			router.ServeHTTP(w, req)

			//Проверяем ответ
			assert.Equal(t, test.expectedCode, w.Code)
		})
	}
}
