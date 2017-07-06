package DBModel

type Questioncontent struct {
	Id              int `beedb:"PK"`
	QuestionId      int
	QuestionContent string
}
