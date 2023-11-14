package repository

import "github.com/ss49919201/diary-article/internal/domain/entity"

type UserActivity interface {
	Save(draft entity.UserActivity) error
}
