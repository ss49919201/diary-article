package service

import (
	"context"
	"testing"
)

func Test_FindDiary(t *testing.T) {
	s := DiaryService{}
	in := FindDiaryInput{}
	s.FindDiary(context.Background(), &in)
	t.Log("Not implemented")
}
