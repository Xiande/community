package DBModel

type Followuser struct {
	Id               int `beedb:"PK"`
	UserName         string
	FollowedUserName string
}
