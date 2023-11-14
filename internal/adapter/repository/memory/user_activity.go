package memory

import (
	"github.com/ss49919201/diary-article/internal/domain/entity"
	"github.com/ss49919201/diary-article/internal/domain/repository"
)

var userActivityMap = map[int]entity.UserActivity{}

type userActivity struct{}

func NewUserActivity() repository.UserActivity {
	return &userActivity{}
}

func (u *userActivity) Save(draft entity.UserActivity) error {
	userActivityMap[draft.ID()] = draft
	return nil
}
