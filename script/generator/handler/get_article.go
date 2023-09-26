package main

import "net/http"

type GetArticle interface {
	Do(w http.ResponseWriter, r *http.Request)
}

type getArticle struct {
	// not implemented
}

func NewGetArticle() (GetArticle, error) {
	return getArticle{
		// not implemented
	}, nil
}

func (_ getArticle) Do(w http.ResponseWriter, r *http.Request) {
	// not implemented
}
