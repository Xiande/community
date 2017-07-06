package DBModel

//"time"

type Tag struct {
	Id          int `beedb:"PK"`
	TagName     string
	IsSystemTag bool
}

type TagExt struct {
	Id               int
	TagName          string
	IsSystemTag      bool
	IsMyFavorite     bool
	TagQuestionCount int
}
