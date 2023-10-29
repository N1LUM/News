package authorRepository

import (
	"ProdajaSute/internal/models"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"runtime"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (p *AuthorRepository) Create(author *models.Author) (*models.Author, error) {
	result := p.db.Create(&author)
	if result.Error != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": result.Error,
		}).Error("Author couldn't CREATE")
		return nil, result.Error
	}
	logrus.Info("Author created successful")
	return author, nil
}
func (p *AuthorRepository) GetList() ([]*models.Author, error) {
	var author []*models.Author
	result := p.db.Find(&author)
	if result.Error != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": result.Error,
		}).Error("Author couldn't GET LIST")
		return nil, result.Error
	}
	logrus.Info("Author list got successful")
	return author, nil
}
func (p *AuthorRepository) GetByID(id string) (*models.Author, error) {
	author := models.Author{}
	result := p.db.Where("id = ?", id).Find(&author)
	if result.Error != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": result.Error,
		}).Error("Author by ID didn't get")
		return nil, result.Error
	}
	logrus.Info("Author by ID got successful")
	return &author, nil
}
func (p *AuthorRepository) Update(author *models.Author, id string) (*models.Author, error) {
	result := p.db.Where("id = ?", id).Save(&author)
	if result.Error != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": result.Error,
		}).Error("Author couldn't UPDATE")
		return nil, result.Error
	}
	logrus.Info("Author by ID update successful")
	return author, nil
}
func (p *AuthorRepository) Delete(id string) (bool, error) {
	author := models.Author{}
	result := p.db.Where("id = ?", id).Delete(&author)
	logrus.Info(result.Value)
	if result.Error != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": result.Error,
		}).Error("Author couldn't DELETE")
		return false, result.Error
	}
	logrus.Info("Author by ID delete successful")
	return true, nil
}
