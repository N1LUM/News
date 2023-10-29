package configTest

import (
	"ProdajaSute/internal/config"
	"ProdajaSute/internal/generator"
	"ProdajaSute/internal/handlers"
	"ProdajaSute/internal/models"
	"ProdajaSute/internal/repository"
	"ProdajaSute/internal/service"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime"
)

func Settings() (*repository.Repository, *service.Service) {
	cfg := config.Config()
	configParam := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBname, cfg.SSLmode)
	logrus.Info(configParam)
	db, err := gorm.Open("postgres", configParam)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Fatal("Can't connect to DB")
	} else {
		logrus.Info("DB connection successful!")
	}
	db.AutoMigrate(&models.Paper{}, &models.Author{})

	gen := generator.Generator{}
	repository := repository.NewRepository(db)
	service := service.NewService(repository, gen)
	handler := handlers.Handler{Service: *service}

	router := handler.InitRouter()
	http.ListenAndServe(":5433", router)

	return repository, service
}
