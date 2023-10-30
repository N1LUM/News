package handlers

import (
	"ProdajaSute/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"path/filepath"
	"runtime"
)

func (h *Handler) CreatePaper(ctx *gin.Context) {
	var paper models.PaperCreateInput
	service := h.Service.AuthorService

	err := ctx.ShouldBind(&paper)
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

	file, err := paper.TitleImage.Open()
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("ImageTitle wasn't read")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	defer file.Close()
	filepath := filepath.Join("content/images", paper.TitleImage.Filename)
	err = ctx.SaveUploadedFile(&paper.TitleImage, filepath)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("ImageTitle wasn't save")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	paper.TitleImageURL = filepath
	author, err := service.GetByID(paper.AuthorID)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Paper by ID didn't get")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	result, err := h.Service.PaperService.Create(&paper, author)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Paper didn't create")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Paper created successful", "paper": result})
}

func (h *Handler) GetListPaper(ctx *gin.Context) {
	result, err := h.Service.PaperService.GetList()

	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Paper List didn't get")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Paper list got successful", "Paper's list": result})
}

func (h *Handler) GetByIDPaper(ctx *gin.Context) {
	result, err := h.Service.PaperService.GetByID(ctx.Param("id"))
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Paper by ID didn't get")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Paper got successful", "paper": result})
}
func (h *Handler) UpdatePaper(ctx *gin.Context) {
	var paper models.PaperUpdateInput
	err := ctx.ShouldBind(&paper)
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
	result, err := h.Service.PaperService.Update(&paper, ctx.Param("id"))
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Paper by ID didn't update")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Paper updated successful", "paper": result})
}
func (h *Handler) DeletePaper(ctx *gin.Context) {
	result, err := h.Service.PaperService.Delete(ctx.Param("id"))
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Paper didn't delete")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Paper deleted successful", "paper": result})
}
