package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ss49919201/diary/internal/dicontainer"
	"github.com/ss49919201/diary/internal/usecase"
)

type ListArticles interface {
	Do(w http.ResponseWriter, r *http.Request)
}

type listArticles struct {
	listArticlesUsecase usecase.ListArticles
}

func NewListArticles() (ListArticles, error) {
	return listArticles{
		listArticlesUsecase: dicontainer.MustInvoke[usecase.ListArticles](),
	}, nil
}

func (l listArticles) Do(w http.ResponseWriter, r *http.Request) {
	articles, err := l.listArticlesUsecase.Do(r.Context(), usecase.ListArticlesInput{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, _ := json.Marshal(articles)

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
