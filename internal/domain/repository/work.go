package repository

import (
	"github.com/comeonjy/go-layout/internal/domain/entity"
)

type WorkRepo interface {
	Get(id int) (*entity.WorkModel, error)
}
