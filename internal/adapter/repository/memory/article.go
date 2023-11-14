package memory

import (
	"github.com/ss49919201/diary-article/internal/domain/entity"
	"github.com/ss49919201/diary-article/internal/domain/repository"
)

var articleMap = map[int]entity.Article{}

type article struct{}

func NewArticle() repository.Article {
	return &article{}
}

func (a *article) Save(draft entity.Article) error {
	articleMap[draft.ID()] = draft
	return nil
}
