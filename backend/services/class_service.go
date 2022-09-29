package services

import (
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/types"
)

type ClassService struct {
	ClassRepository *repositories.ClassRepository
}

func NewClassService(classRepository *repositories.ClassRepository) *ClassService {
	return &ClassService{
		ClassRepository: classRepository,
	}
}

func (classService *ClassService) ListClass(query types.ListClassQuery) ([]types.ClassDto, error) {
	classes, err := classService.ClassRepository.ListClass(query)
	if err != nil {
		return make([]types.ClassDto, 0), err
	}
	classList := make([]types.ClassDto, len(classes))
	for i, class := range classes {
		classList[i] = types.ClassDto{
			ClassID:   class.ClassID,
			ClassName: class.ClassName,
		}
	}
	return classList, nil
}
