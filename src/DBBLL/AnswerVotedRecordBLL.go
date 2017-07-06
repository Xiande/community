// DBBLL project DBBLL.go
package DBBLL

import (
	DAL "DBDAL"
	Model "DBModel"
	"fmt"
)

type AnswerVotedRecordBLL struct {
	BaseBLL
}

func NewAnswerVotedRecordBLL(driverName, conn string) (bll *AnswerVotedRecordBLL) {
	bll = new(AnswerVotedRecordBLL)
	//	bll.driverName = driverName
	//	bll.connectionString = conn
	bll.dal = DAL.NewAnswerVotedRecordDAL(driverName, conn)
	return
}

func (bll *AnswerVotedRecordBLL) Add(model Model.Answervotedrecord) (id int) {
	id = bll.dal.Add(model)
	return
}

func (bll *AnswerVotedRecordBLL) Update(model Model.Answervotedrecord) {
	bll.dal.Update(model)
}

func (bll *AnswerVotedRecordBLL) Delete(id int) {
	bll.dal.Delete(id)
}

func (bll *AnswerVotedRecordBLL) DeleteModel(model *Model.Answervotedrecord) {
	bll.dal.DeleteModel(model)
}

func (bll *AnswerVotedRecordBLL) GetById(id int) (model *Model.Answervotedrecord) {
	m := bll.dal.GetById(id)
	model = m.(*Model.Answervotedrecord)
	return
}

func (bll *AnswerVotedRecordBLL) GetAnswerVotedRcord(userName string, answerId int) (model *Model.Answervotedrecord) {
	model = nil
	whereStr := fmt.Sprintf("UserName = '%s' And AnswerId = %d", userName, answerId)
	ms := bll.dal.GetList(whereStr)
	models := ms.([]Model.Answervotedrecord)
	if len(models) > 0 {
		model = &models[0]
	}
	return
}

func (bll *AnswerVotedRecordBLL) GetList(whereStr string) (models []Model.Answervotedrecord) {
	ms := bll.dal.GetList(whereStr)
	models = ms.([]Model.Answervotedrecord)

	return
}
