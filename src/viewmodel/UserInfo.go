package viewmodel

import (
	"DBBLL"
	"DBModel"
	"common"
)

type UserInfo struct {
	Base       *DBModel.User
	Statistic  *DBModel.Userstatistic
	IsFollowed bool
}
type Reply struct {
	Question *DBModel.Question
	Answer   *DBModel.Answer
}

func (user *UserInfo) LastQuestions() *[]DBModel.Question {
	bll := DBBLL.NewQuestionBLL(common.Config.DB, common.Config.DBConn)
	qs, _ := bll.GetPagingList(1, 10, "", "", "my", "", "all", user.Base.Username)

	return &qs
}

func (user *UserInfo) LastReplies() *[]Reply {
	aBll := DBBLL.NewAnswerBLL(common.Config.DB, common.Config.DBConn)
	qBll := DBBLL.NewQuestionBLL(common.Config.DB, common.Config.DBConn)
	as := aBll.GetLastReply(user.Base.Username, 10)
	var rs []Reply
	for _, a := range as {
		var r Reply
		r.Answer = &a
		r.Question = qBll.GetById(a.QuestionId)

		rs = append(rs, r)
	}

	return &rs
}
