package DBModel

import (
	"time"
)

type Favoritetagcount struct {
	TagId         int
	FavoriteTagId int
	TagName       string
	CreateDate    time.Time
	TagCounts     int
}
