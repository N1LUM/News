package authorService

import (
	"ProdajaSute/internal/generator"
	"ProdajaSute/internal/models"
	"ProdajaSute/internal/repository"
	"github.com/sirupsen/logrus"
	"runtime"
	"time"
)

type AuthorService struct {
	rep repository.AuthorRepository
	gen generator.Generator
}

func NewAuthorService(rep repository.AuthorRepository, gen generator.Generator) *AuthorService {
	return &AuthorService{
		rep: rep,
		gen: gen,
	}
}

func (a *AuthorService) Create(authorInput *models.AuthorCreateInput) (*models.Author, error) {
	bornDate, err := time.Parse("2006-01-02", authorInput.BornDate)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("BornDate didn't parse")
		return nil, err
	}
	author := models.Author{
		ID:        a.gen.GenerateID(),
		Username:  authorInput.Username,
		Name:      authorInput.Name,
		Lastname:  authorInput.Lastname,
		Sex:       authorInput.Sex,
		BornDate:  bornDate,
		Country:   authorInput.Country,
		City:      authorInput.City,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
	result, err := a.rep.Create(&author)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Author didn't create")
		return nil, err
	}
	return result, nil
}
func (a *AuthorService) GetList() ([]*models.Author, error) {
	result, err := a.rep.GetList()
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Author List didn't get")
		return nil, err
	}
	return result, nil
}

func (a *AuthorService) GetByID(id string) (*models.Author, error) {
	result, err := a.rep.GetByID(id)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Author by ID didn't get")
		return nil, err
	}
	return result, nil
}
func (a *AuthorService) Update(authorInput *models.AuthorUpdateInput, id string) (*models.Author, error) {
	author, err := a.rep.GetByID(id)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Author by ID didn't get")
		return nil, err
	}
	if authorInput.Username != nil {
		author.Username = *authorInput.Username
	}
	if authorInput.Name != nil {
		author.Name = *authorInput.Name
	}
	if authorInput.Lastname != nil {
		author.Lastname = *authorInput.Lastname
	}
	if authorInput.Sex != nil {
		author.Sex = *authorInput.Sex
	}
	if authorInput.BornDate != nil {
		bornDate, err := time.Parse("2006-01-02", *authorInput.BornDate)
		if err != nil {
			_, file, line, _ := runtime.Caller(0)
			logrus.WithFields(logrus.Fields{
				"file":  file,
				"line":  line - 2,
				"error": err,
			}).Error("BornDate didn't parse")
			return nil, err
		}
		author.BornDate = bornDate
	}
	if authorInput.Country != nil {
		author.Country = *authorInput.Country
	}
	if authorInput.City != nil {
		author.City = *authorInput.City
	}
	result, err := a.rep.Update(author, id)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Author by ID didn't update")
		return nil, err
	}
	return result, nil
}

func (a *AuthorService) Delete(id string) (bool, error) {
	result, err := a.rep.Delete(id)
	logrus.Info(id)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Author didn't delete")
		return false, err
	}
	return result, nil
}
