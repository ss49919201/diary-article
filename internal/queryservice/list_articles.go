package queryservice

import (
	"context"

	"github.com/ss49919201/diary-article/internal/domain/entity"
)

type ListArticles interface {
	Do(context.Context, ListArticlesInput) ([]entity.Article, error)
}

type ListArticlesInput struct{}

type listArticles struct {
}

func NewListArticles() (ListArticles, error) {
	return listArticles{}, nil
}

func (_ listArticles) Do(context.Context, ListArticlesInput) ([]entity.Article, error) {
	// Mock
	return []entity.Article{
		{}, {},
	}, nil
}
