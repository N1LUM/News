package paperService

import (
	"ProdajaSute/internal/generator"
	"ProdajaSute/internal/models"
	"ProdajaSute/internal/repository"
	"github.com/sirupsen/logrus"
	"runtime"
	"time"
)

type PaperService struct {
	rep repository.PaperRepository
	gen generator.Generator
}

func NewPaperService(rep repository.PaperRepository, gen generator.Generator) *PaperService {
	return &PaperService{
		rep: rep,
		gen: gen,
	}
}

func (p *PaperService) Create(paperInput *models.PaperCreateInput, author *models.Author) (*models.Paper, error) {
	paper := models.Paper{
		ID:            p.gen.GenerateID(),
		TitleImageURL: paperInput.TitleImageURL,
		Title:         paperInput.Title,
		Subtitle:      paperInput.Subtitle,
		Content:       paperInput.Content,
		AuthorID:      paperInput.AuthorID,
		Author:        *author,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		DeletedAt:     nil,
	}
	result, err := p.rep.Create(&paper)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Paper couldn't CREATE")
		return nil, err
	}
	return result, nil
}

func (p *PaperService) GetList() ([]*models.Paper, error) {
	result, err := p.rep.GetList()
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Paper List didn't get")
		return nil, err
	}
	return result, nil
}

func (p *PaperService) GetByID(id string) (*models.Paper, error) {
	result, err := p.rep.GetByID(id)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Paper by ID didn't get")
		return nil, err
	}
	return result, nil
}

func (p *PaperService) Update(paperInput *models.PaperUpdateInput, id string) (*models.Paper, error) {
	paper, err := p.rep.GetByID(id)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Paper by ID didn't get")
		return nil, err
	}
	if paperInput.TitleImageURL != nil {
		paper.TitleImageURL = *paperInput.TitleImageURL
	}
	if paperInput.Title != nil {
		paper.Title = *paperInput.Title
	}
	if paperInput.Subtitle != nil {
		paper.Subtitle = *paperInput.Subtitle
	}
	if paperInput.Content != nil {
		paper.Content = *paperInput.Content
	}
	result, err := p.rep.Update(paper, id)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Paper by ID didn't update")
		return nil, err
	}
	return result, nil
}

func (p *PaperService) Delete(id string) (bool, error) {
	result, err := p.rep.Delete(id)
	logrus.Info(id)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.WithFields(logrus.Fields{
			"file":  file,
			"line":  line - 2,
			"error": err,
		}).Error("Paper didn't delete")
		return false, err
	}
	return result, nil
}
