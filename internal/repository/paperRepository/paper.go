package paperRepository

import (
	"ProdajaSute/internal/models"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"runtime"
)

type PaperRepository struct {
	db *gorm.DB
}

func NewPaperRepository(db *gorm.DB) *PaperRepository {
	return &PaperRepository{
		db: db,
	}
}

func (p *PaperRepository) Create(news *models.Paper) (*models.Paper, error) {
	result := p.db.Create(&news)
	if result.Error != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": result.Error,
		}).Error("Author couldn't CREATE")
		return nil, result.Error
	}
	logrus.Info("Article created successful")
	return news, nil
}
func (p *PaperRepository) GetList() ([]*models.Paper, error) {
	var news []*models.Paper
	result := p.db.Find(&news)
	if result.Error != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": result.Error,
		}).Error("Author couldn't GET LIST")
		return nil, result.Error
	}
	logrus.Info("Article list got successful")
	return news, nil
}
func (p *PaperRepository) GetByID(id string) (*models.Paper, error) {
	var news *models.Paper
	result := p.db.Where("id = ?", id).Find(&news)
	if result.Error != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": result.Error,
		}).Error("Author couldn't GET BY ID")
		return nil, result.Error
	}
	logrus.Info("Article by ID got successful")
	return news, nil
}
func (p *PaperRepository) Update(news *models.Paper, id string) (*models.Paper, error) {
	result := p.db.Where("id = ?", id).Save(&news)
	if result.Error != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": result.Error,
		}).Error("Author couldn't UPDATE")
		return nil, result.Error
	}
	logrus.Info("Article by ID update successful")
	return news, nil
}
func (p *PaperRepository) Delete(id string) (bool, error) {
	news := models.Paper{}
	result := p.db.Where("id = ?", id).Delete(&news)
	if result.Error != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": result.Error,
		}).Error("Author couldn't DELETE")
		return false, result.Error
	}
	logrus.Info("Article by ID update successful")
	return true, nil
}
