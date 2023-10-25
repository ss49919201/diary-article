package repository

import "github.com/ss49919201/diary/internal/domain/entity"

type UserActivity interface {
	Create(draft entity.UserActivity) error
}
