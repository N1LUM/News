package handlers

import (
	"ProdajaSute/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime"
)

func (h *Handler) Create(ctx *gin.Context) {
	var author models.AuthorCreateInput
	err := ctx.ShouldBind(&author)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Gin couldn't convert JSON to struct")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	result, err := h.Service.AuthorService.Create(&author)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Author didn't create")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Author created successful", "author": result})
}

func (h *Handler) GetList(ctx *gin.Context) {
	result, err := h.Service.AuthorService.GetList()
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Author List didn't get")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) GetByID(ctx *gin.Context) {
	result, err := h.Service.AuthorService.GetByID(ctx.Param("id"))
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Author by ID didn't get")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, result)
}
func (h *Handler) Update(ctx *gin.Context) {
	var author models.AuthorUpdateInput
	err := ctx.ShouldBind(&author)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Gin couldn't convert JSON to struct")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	result, err := h.Service.AuthorService.Update(&author, ctx.Param("id"))
	ctx.JSON(http.StatusOK, result)
}
func (h *Handler) Delete(ctx *gin.Context) {
	result, err := h.Service.AuthorService.Delete(ctx.Param("id"))
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Author didn't delete")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, result)
}
