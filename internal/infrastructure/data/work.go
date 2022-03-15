package data

import (
	"github.com/comeonjy/go-layout/internal/domain/entity"
	"github.com/comeonjy/go-layout/internal/domain/repository"
)

type workRepo struct {
	data *Data
}

func NewWorkRepo(data *Data) repository.WorkRepo {
	return &workRepo{data: data}
}

func (w workRepo) Get(id int) (*entity.WorkModel, error) {
	return &entity.WorkModel{
		ID:  id,
		Url: "https://www.a.com",
	}, nil
}
