package handler

import (
	"net/http"

	"github.com/GaponovAlexey/practic-web/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", nil)
		auth.POST("/sign-in", nil)
	}
	api := router.Group("/api", nil)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", nil)
			lists.GET("/", nil)
			lists.GET("/:id", nil)
			lists.PUT("/:id", nil)
			lists.DELETE("/:id", nil)

			items := router.Group(":id/items", nil)
			{
				items.POST("/", nil)
				items.GET("/", nil)
				items.GET("/", nil)
				items.PUT("/", nil)
				items.DELETE("/", nil)
			}
		}
	}
	return router
}

//test

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}