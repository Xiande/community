package DBModel

import (
	"time"
)

type Answer struct {
	Id          int `beedb:"PK"`
	QuestionId  int
	UserName    string
	DisplayName string
	IsBest      bool
	BestTime    time.Time
	IsExpert    bool
	ExpertTime  time.Time
	VotedCount  int
	CreateDate  time.Time
	CreateBy    string
}
