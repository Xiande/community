package DBModel

import (
	"time"
)

type Answervotedrecord struct {
	Id         int `beedb:"PK"`
	QuestionId int
	AnswerId   int
	UserName   string
	CreateDate time.Time
	CreateBy   string
}
