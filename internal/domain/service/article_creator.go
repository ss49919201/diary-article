package service

import (
	"github.com/ss49919201/diary-article/internal/domain/entity"
	"github.com/ss49919201/diary-article/internal/domain/repository"
)

type ArticleCreator interface {
	Exec(in CreateArticleInput) error
}

type CreateArticleInput struct {
	UserID       int
	ArticleBody  string
	ArticleTitle string
}

type articleCreator struct {
	article      repository.Article
	userActivity repository.UserActivity
}

func (a *articleCreator) Exec(in CreateArticleInput) error {
	article := entity.NewArticle(in.ArticleTitle, in.ArticleBody, in.UserID)
	if err := a.article.Save(article); err != nil {
		return err
	}

	userActivity := entity.NewFromActivity(article)
	return a.userActivity.Save(userActivity)
}
