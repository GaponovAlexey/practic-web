package handler

import (
	"github.com/GaponovAlexey/practic-web/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createLists)
			lists.GET("/", h.getAllILists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateLists)
			lists.DELETE("/:id", h.deleteLists)

			items := router.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT(":item_id", h.updateItem)
				items.DELETE(":item_id", h.deleteItem)
			}
		}
	}
	return router
}
