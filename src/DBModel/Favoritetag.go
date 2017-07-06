package DBModel

import (
	"time"
)

type Favoritetag struct {
	Id         int `beedb:"PK"`
	TagId      int
	UserName   string
	CreateDate time.Time
	CreateBy   string
}
