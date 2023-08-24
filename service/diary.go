package service

import "context"

type Diary struct {
	Title string
}

type DiaryService struct{}

type FindDiaryInput struct{}

type FindDiaryOutput struct {
	Err error
}

func (d *DiaryService) FindDiary(ctx context.Context, input *FindDiaryInput) *FindDiaryOutput {
	return &FindDiaryOutput{}
}
