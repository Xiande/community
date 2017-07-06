package viewmodel

import (
	"DBModel"
)

type UserStatistic struct {
	DBModel.Userstatistic
	PhotoStr    string
	BadgesCount int
}
