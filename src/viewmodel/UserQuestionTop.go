package viewmodel

import (
	"time"
)

type UserQuestionTop struct {
	Id          int
	UserName    string
	DisplayName string
	Title       string
	Photo       string
	CreateDate  time.Time
}
