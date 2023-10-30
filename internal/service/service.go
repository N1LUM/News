package service

import (
	"ProdajaSute/internal/generator"
	"ProdajaSute/internal/models"
	"ProdajaSute/internal/repository"
	"ProdajaSute/internal/service/authorService"
	"ProdajaSute/internal/service/paperService"
)

type PaperService interface {
	Create(paperInput *models.PaperCreateInput, authorID *models.Author) (*models.Paper, error)
	GetList() ([]*models.Paper, error)
	GetByID(id string) (*models.Paper, error)
	Update(paperInput *models.PaperUpdateInput, id string) (*models.Paper, error)
	Delete(id string) (bool, error)
}
type AuthorService interface {
	Create(authorInput *models.AuthorCreateInput) (*models.Author, error)
	GetList() ([]*models.Author, error)
	GetByID(id string) (*models.Author, error)
	Update(authorInput *models.AuthorUpdateInput, id string) (*models.Author, error)
	Delete(id string) (bool, error)
}

type Service struct {
	PaperService
	AuthorService
}

func NewService(rep *repository.Repository, gen generator.Generator) *Service {
	return &Service{
		PaperService:  paperService.NewPaperService(rep.PaperRepository, gen),
		AuthorService: authorService.NewAuthorService(rep.AuthorRepository, gen),
	}
}
