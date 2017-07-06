// DBBLL project DBBLL.go
package DBBLL

import (
	DAL "DBDAL"
	Model "DBModel"
	"fmt"
)

type AnswerBLL struct {
	BaseBLL
}

func NewAnswerBLL(driverName, conn string) (bll *AnswerBLL) {
	bll = new(AnswerBLL)
	//	bll.driverName = driverName
	//	bll.connectionString = conn
	bll.dal = DAL.NewAnswerDAL(driverName, conn)
	return
}

func (bll *AnswerBLL) Add(model Model.Answer) (id int) {
	id = bll.dal.Add(model)
	return
}

func (bll *AnswerBLL) AddTx(model Model.Answer, content string) (id int) {
	dal := bll.dal.(*DAL.AnswerDAL)
	id = dal.AddTx(model, content)
	return
}

func (bll *AnswerBLL) Update(model Model.Answer) {
	bll.dal.Update(model)
}

func (bll *AnswerBLL) Delete(id int) {
	bll.dal.Delete(id)
}

func (bll *AnswerBLL) DeleteModel(model *Model.Answer) {
	bll.dal.DeleteModel(model)
}

func (bll *AnswerBLL) GetById(id int) (model *Model.Answer) {
	m := bll.dal.GetById(id)
	model = m.(*Model.Answer)
	return
}

func (bll *AnswerBLL) GetBestByQuestionId(qid int) (model *Model.Answer) {
	ms := bll.dal.GetList(fmt.Sprintf("IsBest = 1 And QuestionId = %d", qid))
	models := ms.([]Model.Answer)
	if models == nil || len(models) == 0 {
		model = nil
	} else {
		model = &models[0]
	}

	return
}

func (bll *AnswerBLL) GetExpertByQuestionId(qid int) (model *Model.Answer) {
	ms := bll.dal.GetList(fmt.Sprintf("IsExpert = 1 And QuestionId = %d", qid))
	models := ms.([]Model.Answer)
	if models == nil || len(models) == 0 {
		model = nil
	} else {
		model = &models[0]
	}

	return
}

func (bll *AnswerBLL) GetListByQuestionId(qid int) (models []Model.Answer) {
	ms := bll.dal.GetList(fmt.Sprintf("QuestionId = %d", qid))
	models = ms.([]Model.Answer)

	return
}

func (bll *AnswerBLL) GetListByUserName(userName string) (models []Model.Answer) {
	ms := bll.dal.GetList(fmt.Sprintf("UserName = %s", userName))
	models = ms.([]Model.Answer)

	return
}

func (bll *AnswerBLL) GetLastReply(userName string, top int) (models []Model.Answer) {
	dal := bll.dal.(*DAL.AnswerDAL)
	models = dal.GetLastReply(userName, top)
	return
}
func (bll *AnswerBLL) GetPagingListByQuestionId(qid, aid, pageIndex, pageSize int, sortField, sortOrder string) (models []Model.Answer, count int) {
	dal := bll.dal.(*DAL.AnswerDAL)
	ms, count := dal.GetPagingByQuestionId(qid, aid, pageIndex, pageSize, sortField, sortOrder)
	models = ms.([]Model.Answer)

	return
}

func (bll *AnswerBLL) GetList(whereStr string) (models []Model.Answer) {
	ms := bll.dal.GetList(whereStr)
	models = ms.([]Model.Answer)

	return
}

func (bll *AnswerBLL) ExpertAnswer(qid, aid int) {
	dal := bll.dal.(*DAL.AnswerDAL)

	dal.ExpertAnswer(qid, aid)
}

func (bll *AnswerBLL) BestAnswer(qid, aid int) {
	dal := bll.dal.(*DAL.AnswerDAL)

	dal.BestAnswer(qid, aid)
}
