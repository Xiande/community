// DBBLL project DBBLL.go
package DBBLL

import (
	DAL "DBDAL"
	Model "DBModel"
	"fmt"
)

type FavoriteQuestionBLL struct {
	BaseBLL
}

func NewFavoriteQuestionBLL(driverName, conn string) (bll *FavoriteQuestionBLL) {
	bll = new(FavoriteQuestionBLL)
	//	bll.driverName = driverName
	//	bll.connectionString = conn
	bll.dal = DAL.NewFavoriteQuestionDAL(driverName, conn)
	return
}

func (bll *FavoriteQuestionBLL) Add(model Model.Favoritequestion) (id int) {
	id = bll.dal.Add(model)
	return
}

func (bll *FavoriteQuestionBLL) Update(model Model.Favoritequestion) {
	bll.dal.Update(model)
}

func (bll *FavoriteQuestionBLL) Delete(id int) {
	bll.dal.Delete(id)
}

func (bll *FavoriteQuestionBLL) DeleteModel(model *Model.Favoritequestion) {
	bll.dal.DeleteModel(model)
}

func (bll *FavoriteQuestionBLL) GetById(id int) (model *Model.Favoritequestion) {
	m := bll.dal.GetById(id)
	model = m.(*Model.Favoritequestion)
	return
}

func (bll *FavoriteQuestionBLL) GetByQuestionId(qid int) (model *Model.Favoritequestion) {
	ms := bll.dal.GetList(fmt.Sprintf("QuestionId = %d", qid))
	models := ms.([]Model.Favoritequestion)
	if models == nil || len(models) == 0 {
		model = nil
	} else {
		model = &models[0]
	}

	return
}

func (bll *FavoriteQuestionBLL) GetListByUesrName(pIndex, pSize int, userName string) (models []Model.Favoritequestionext, count int) {
	dal := bll.dal.(*DAL.FavoriteQuestionDAL)
	models, count = dal.GetFavoriteQuestionList(pIndex, pSize, userName)

	return
}

func (bll *FavoriteQuestionBLL) GetFavoriteQuestion(userName string, qid int) (model *Model.Favoritequestion) {
	model = nil
	whereStr := fmt.Sprintf("UserName = '%s' And QuestionId = %d", userName, qid)
	ms := bll.dal.GetList(whereStr)
	models := ms.([]Model.Favoritequestion)
	if len(models) > 0 {
		model = &models[0]
	}
	return
}

func (bll *FavoriteQuestionBLL) GetList(whereStr string) (models []Model.Favoritequestion) {
	ms := bll.dal.GetList(whereStr)
	models = ms.([]Model.Favoritequestion)

	return
}
