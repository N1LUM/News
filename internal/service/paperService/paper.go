package paperService

import (
	"ProdajaSute/internal/generator"
	"ProdajaSute/internal/models"
	"ProdajaSute/internal/repository"
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

func (p *PaperService) Create(paperInput *models.PaperCreateInput) (*models.Paper, error) {
	return nil, nil
}

func (a *PaperService) GetList() ([]*models.Paper, error) {
	return nil, nil
}

func (a *PaperService) GetByID(id string) (*models.Paper, error) {
	return nil, nil
}

func (a *PaperService) Update(paperInput *models.PaperUpdateInput, id string) (*models.Paper, error) {
	return nil, nil
}

func (a *PaperService) Delete(id string) (bool, error) {
	return true, nil
}
