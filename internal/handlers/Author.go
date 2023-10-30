package handlers

import (
	"ProdajaSute/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime"
)

func (h *Handler) CreateAuthor(ctx *gin.Context) {
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

func (h *Handler) GetListAuthor(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, gin.H{"message": "Author list got successful", "author": result})
}

func (h *Handler) GetByIDAuthor(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, gin.H{"message": "Author got successful", "author": result})
}
func (h *Handler) UpdateAuthor(ctx *gin.Context) {
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
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Author by ID didn't update")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Author updated successful", "author": result})
}
func (h *Handler) DeleteAuthor(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, gin.H{"message": "Author deleted successful", "author": result})
}
