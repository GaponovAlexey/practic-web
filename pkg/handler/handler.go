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
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}
	api := router.Group("/api", )
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createLists)
			lists.GET("/", h.getAllILists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateLists)
			lists.DELETE("/:id", h.deleteLists)

			items := router.Group(":id/items", nil)
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/", h.getItemById)
				items.PUT("/", h.updateItem)
				items.DELETE("/", h.deleteItem)
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