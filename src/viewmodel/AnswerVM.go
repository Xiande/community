package viewmodel

import (
	"DBModel"
)

type AnswerVM struct {
	DBModel.Answer
	AnswerContent string
	CanBest       bool
	CanExpert     bool
	AuthorEmail   string
	PhotoImgSrc   string
	IsFollowed    bool
	IsAuthor      bool
	IsAdmin       bool
	IsModerator   bool
}
