package handlers

import (
	"effective_mobile/pkg/service"

	_ "effective_mobile/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handlers struct {
	service *service.Service
}

func New(service *service.Service) *Handlers {
	return &Handlers{service: service}
}

// Инициализируем пути по которым будут идти запросы
func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.New()
	//Для swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Запрос на добавление новой песни
	router.POST("/create_new_song", h.createNewSong)
	//Получение всех песен
	router.GET("/get_all_songs", h.getAllSongs)
	//Получение информации о песни
	router.GET("/info", h.info)
	//Удаление песни
	router.DELETE("/delete_song", h.deleteSong)
	//Изменение данных в песни
	router.PATCH("/change_song", h.changeSong)
	//Получение песни по куплетам
	router.GET("/get_song_by_verses", h.getSongByVerses)

	return router
}
