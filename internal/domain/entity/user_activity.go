package entity

type UserActivity struct {
	id     int
	userID int
	label  UserActivityLabel
}

type UserActivityLabel string

const (
	UserActivityLabelCreateArticle UserActivityLabel = "create_article"
	UserActivityLabelCreateUser    UserActivityLabel = "create_user"
)

func NewFromActivity(a Article) UserActivity {
	return UserActivity{
		userID: a.id,
		label:  UserActivityLabelCreateArticle,
	}
}
