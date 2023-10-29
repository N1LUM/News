package repository

import (
	"ProdajaSute/internal/models"
	"ProdajaSute/internal/repository/authorRepository"
	"ProdajaSute/internal/repository/paperRepository"
	"github.com/jinzhu/gorm"
)

type PaperRepository interface {
	Create(news *models.Paper) (*models.Paper, error)
	GetList() ([]*models.Paper, error)
	GetByID(id string) (*models.Paper, error)
	Update(news *models.Paper, id string) (*models.Paper, error)
	Delete(id string) (bool, error)
}
type AuthorRepository interface {
	Create(author *models.Author) (*models.Author, error)
	GetList() ([]*models.Author, error)
	GetByID(id string) (*models.Author, error)
	Update(author *models.Author, id string) (*models.Author, error)
	Delete(id string) (bool, error)
}

type Repository struct {
	PaperRepository
	AuthorRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		PaperRepository:  paperRepository.NewPaperRepository(db),
		AuthorRepository: authorRepository.NewAuthorRepository(db),
	}
}
