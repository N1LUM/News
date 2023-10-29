package tests

import (
	"ProdajaSute/internal/models"
	"ProdajaSute/tests/configTest"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	_, ser := configTest.Settings()
	author := models.AuthorInput{
		Username: "john_doe",
		Name:     "John",
		Lastname: "Doe",
		Sex:      "Male",
		Age:      30,
		BornDate: time.Date(1992, time.January, 15, 0, 0, 0, 0, time.UTC),
		Country:  "USA",
		City:     "New York",
	}
	res, _ := ser.AuthorService.Create(&author)
	logrus.Info(res)
}
