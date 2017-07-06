package viewmodel

import (
	"DBModel"
)

type FavoriteQuestionVM struct {
	DBModel.Favoritequestionext
	PhotoImgSrc string
	IsAuthor    bool
	IsFollowed  bool
	IsAdmin     bool
	IsModerator bool
}
