// DBBLL project DBBLL.go
package DBBLL

import (
	DAL "DBDAL"
	Model "DBModel"
	"fmt"
)

type FavoriteAnswerBLL struct {
	BaseBLL
}

func NewFavoriteAnswerBLL(driverName, conn string) (bll *FavoriteAnswerBLL) {
	bll = new(FavoriteAnswerBLL)
	//	bll.driverName = driverName
	//	bll.connectionString = conn
	bll.dal = DAL.NewFavoriteAnswerDAL(driverName, conn)
	return
}

func (bll *FavoriteAnswerBLL) Add(model Model.Favoriteanswer) (id int) {
	id = bll.dal.Add(model)
	return
}

func (bll *FavoriteAnswerBLL) Update(model Model.Favoriteanswer) {
	bll.dal.Update(model)
}

func (bll *FavoriteAnswerBLL) Delete(id int) {
	bll.dal.Delete(id)
}

func (bll *FavoriteAnswerBLL) DeleteModel(model *Model.Favoriteanswer) {
	bll.dal.DeleteModel(model)
}

func (bll *FavoriteAnswerBLL) GetById(id int) (model *Model.Favoriteanswer) {
	m := bll.dal.GetById(id)
	model = m.(*Model.Favoriteanswer)
	return
}

func (bll *FavoriteAnswerBLL) GetByAnswerId(answerId int) (model *Model.Favoriteanswer) {
	ms := bll.dal.GetList(fmt.Sprintf("AnswerId = %d", answerId))
	models := ms.([]Model.Favoriteanswer)
	if models == nil || len(models) == 0 {
		model = nil
	} else {
		model = &models[0]
	}

	return
}

func (bll *FavoriteAnswerBLL) GetFavoriteAnswer(userName string, answerId int) (model *Model.Favoriteanswer) {
	model = nil
	whereStr := fmt.Sprintf("UserName = '%s' And AnswerId = %d", userName, answerId)
	ms := bll.dal.GetList(whereStr)
	models := ms.([]Model.Favoriteanswer)
	if len(models) > 0 {
		model = &models[0]
	}
	return
}

func (bll *FavoriteAnswerBLL) GetList(whereStr string) (models []Model.Favoriteanswer) {
	ms := bll.dal.GetList(whereStr)
	models = ms.([]Model.Favoriteanswer)

	return
}
