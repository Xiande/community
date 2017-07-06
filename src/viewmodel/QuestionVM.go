package viewmodel

import (
	"DBModel"
)

type QuestionVM struct {
	DBModel.Question
	QuestionContent string
	BestAnswer      AnswerVM
	ExpertAnswer    AnswerVM
	AuthorEmail     string
	PhotoImgSrc     string
	IsAuthor        bool
	IsFollowed      bool
	IsAdmin         bool
	IsModerator     bool
}
