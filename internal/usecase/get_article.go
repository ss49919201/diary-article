package usecase

type GetArticle interface {
	Do()
}

type getArticle struct {
	// not implemented
}

func NewGetArticle() (GetArticle, error) {
	return getArticle{
		// not implemented
	}, nil
}

func (_ getArticle) Do() {
	// not implemented
}
