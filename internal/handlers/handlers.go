package handlers

import (
	"ProdajaSute/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service.Service
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()
	//auth := router.Group("/auth")
	//{
	//	auth.POST("sign-up")
	//	auth.POST("sign-in")
	//}
	user := router.Group("/authors")
	{
		user.POST("/create", h.Create)
		user.GET("/getList", h.GetList)
		user.GET("/getByID/:id", h.GetByID)
		user.PUT("/update/:id", h.Update)
		user.DELETE("/delete/:id", h.Delete)
	}

	return router
}
