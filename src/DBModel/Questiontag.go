package DBModel

//"time"

type Questiontag struct {
	Id         int `beedb:"PK"`
	QuestionId int
	TagId      int
}
