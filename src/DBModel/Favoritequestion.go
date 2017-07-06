package DBModel

import (
	"time"
)

type Favoritequestion struct {
	Id         int `beedb:"PK"`
	QuestionId int
	UserName   string
	CreateDate time.Time
	CreateBy   string
}

type Favoritequestionext struct {
	Id           int `beedb:"PK"`
	Title        string
	UserName     string
	DisplayName  string
	Tags         string
	EmailNotice  bool
	ViewedCount  int
	BestCount    int
	AnswersCount int
	VotedCount   int
	IsExpert     bool
	ExpertTime   time.Time
	LanguageType string
	CreateDate   time.Time
	CreateBy     string

	FavoriteQuestionId int
	FavoriteAnswerId   int
	AnswerId           int
}
