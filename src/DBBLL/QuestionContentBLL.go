// DBBLL project DBBLL.go
package DBBLL

import (
	DAL "DBDAL"
	Model "DBModel"
	"fmt"
)

type QuestionContentBLL struct {
	BaseBLL
}

func NewQuestionContentBLL(driverName, conn string) (bll *QuestionContentBLL) {
	bll = new(QuestionContentBLL)
	//	bll.driverName = driverName
	//	bll.connectionString = conn
	bll.dal = DAL.NewQuestionContentDAL(driverName, conn)
	return
}

func (bll *QuestionContentBLL) Add(model Model.Questioncontent) (id int) {
	id = bll.dal.Add(model)
	return
}

func (bll *QuestionContentBLL) Update(model Model.Questioncontent) {
	bll.dal.Update(model)
}

func (bll *QuestionContentBLL) Delete(id int) {
	bll.dal.Delete(id)
}

func (bll *QuestionContentBLL) DeleteModel(model *Model.Questioncontent) {
	bll.dal.DeleteModel(model)
}

func (bll *QuestionContentBLL) GetById(id int) (model *Model.Questioncontent) {
	m := bll.dal.GetById(id)
	model = m.(*Model.Questioncontent)
	return
}

func (bll *QuestionContentBLL) GetByQuestionId(qid int) (model *Model.Questioncontent) {
	ms := bll.dal.GetList(fmt.Sprintf("QuestionId = %d", qid))
	models := ms.([]Model.Questioncontent)
	if models == nil || len(models) == 0 {
		model = nil
	} else {
		model = &models[0]
	}

	return
}

func (bll *QuestionContentBLL) GetList(whereStr string) (models []Model.Questioncontent) {
	ms := bll.dal.GetList(whereStr)
	models = ms.([]Model.Questioncontent)

	return
}
