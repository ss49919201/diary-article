package repository

import "github.com/ss49919201/diary/internal/domain/entity"

type User interface {
	Create(draft entity.User) error
}
