package repository

import "github.com/ss49919201/diary/internal/domain/entity"

type Article interface {
	Create(draft entity.Article) error
}
