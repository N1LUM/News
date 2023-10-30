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
		user.POST("/create", h.CreateAuthor)
		user.GET("/getList", h.GetListAuthor)
		user.GET("/getByID/:id", h.GetByIDAuthor)
		user.PUT("/update/:id", h.UpdateAuthor)
		user.DELETE("/delete/:id", h.DeleteAuthor)
	}
	paper := router.Group("/papers")
	{
		paper.POST("/create", h.CreatePaper)
		paper.GET("/getList", h.GetListPaper)
		paper.GET("/getByID/:id", h.GetByIDPaper)
		paper.PUT("/update/:id", h.UpdatePaper)
		paper.DELETE("/delete/:id", h.DeletePaper)
	}

	return router
}
