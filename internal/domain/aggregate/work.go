package aggregate

import (
	"github.com/comeonjy/go-layout/internal/domain/entity"
	"github.com/comeonjy/go-layout/internal/domain/repository"
)

type WorkUseCase struct {
	repo repository.WorkRepo
}

func NewWorkUseCase(repo repository.WorkRepo) *WorkUseCase {
	return &WorkUseCase{
		repo: repo,
	}
}

func (u WorkUseCase) GetInfo(id int) (*entity.WorkModel, error) {
	get, err := u.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return get, nil
}
