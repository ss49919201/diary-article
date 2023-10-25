package entity

type Article struct {
	id     int
	title  string
	body   string
	userID int
}

func NewArticle(title, body string, userID int) Article {
	return Article{
		title:  title,
		body:   body,
		userID: userID,
	}
}
