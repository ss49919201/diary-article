package repository

import "github.com/ss49919201/diary-article/internal/domain/entity"

type Article interface {
	Save(draft entity.Article) error
}
