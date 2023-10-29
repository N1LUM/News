package paperService

import (
	"ProdajaSute/internal/generator"
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
