// DBBLL project DBBLL.go
package DBBLL

import (
	DAL "DBDAL"
	Model "DBModel"
	"fmt"
)

type AnswerContentBLL struct {
	BaseBLL
}

func NewAnswerContentBLL(driverName, conn string) (bll *AnswerContentBLL) {
	bll = new(AnswerContentBLL)
	//	bll.driverName = driverName
	//	bll.connectionString = conn
	bll.dal = DAL.NewAnswerContentDAL(driverName, conn)
	return
}

func (bll *AnswerContentBLL) Add(model Model.Answercontent) (id int) {
	id = bll.dal.Add(model)
	return
}

func (bll *AnswerContentBLL) Update(model Model.Answercontent) {
	bll.dal.Update(model)
}

func (bll *AnswerContentBLL) Delete(id int) {
	bll.dal.Delete(id)
}

func (bll *AnswerContentBLL) DeleteModel(model *Model.Answercontent) {
	bll.dal.DeleteModel(model)
}

func (bll *AnswerContentBLL) GetById(id int) (model *Model.Answercontent) {
	m := bll.dal.GetById(id)
	model = m.(*Model.Answercontent)
	return
}

func (bll *AnswerContentBLL) GetByAnswerId(qid int) (model *Model.Answercontent) {
	ms := bll.dal.GetList(fmt.Sprintf("AnswerId = %d", qid))
	models := ms.([]Model.Answercontent)
	if models == nil || len(models) == 0 {
		model = nil
	} else {
		model = &models[0]
	}

	return
}

func (bll *AnswerContentBLL) GetList(whereStr string) (models []Model.Answercontent) {
	ms := bll.dal.GetList(whereStr)
	models = ms.([]Model.Answercontent)

	return
}
