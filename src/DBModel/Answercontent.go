package DBModel

type Answercontent struct {
	Id            int `beedb:"PK"`
	AnswerId      int
	AnswerContent string
}
