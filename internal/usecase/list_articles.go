package usecase

import (
	"context"

	"github.com/ss49919201/diary-article/internal/dicontainer"
	"github.com/ss49919201/diary-article/internal/domain/entity"
	"github.com/ss49919201/diary-article/internal/queryservice"
)

type ListArticles interface {
	Do(context.Context, ListArticlesInput) ([]entity.Article, error)
}

type ListArticlesInput struct{}

type listArticles struct {
	qsListArticles queryservice.ListArticles
}

func NewListArticles() (ListArticles, error) {
	return listArticles{
		qsListArticles: dicontainer.MustInvoke[queryservice.ListArticles](),
	}, nil
}

func (l listArticles) Do(ctx context.Context, in ListArticlesInput) ([]entity.Article, error) {
	return l.qsListArticles.Do(ctx, queryservice.ListArticlesInput{})
}
