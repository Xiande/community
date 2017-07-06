// DBBLL project DBBLL.go
package DBBLL

import (
	DAL "DBDAL"
	Model "DBModel"
)

type QuestionBLL struct {
	BaseBLL
}

func NewQuestionBLL(driverName, conn string) (bll *QuestionBLL) {
	bll = new(QuestionBLL)
	//	bll.driverName = driverName
	//	bll.connectionString = conn
	bll.dal = DAL.NewQuestionDAL(driverName, conn)
	return
}

func (bll *QuestionBLL) Add(model Model.Question) (id int) {
	id = bll.dal.Add(model)
	return
}

func (bll *QuestionBLL) AddTx(model Model.Question, content string) (id int) {
	dal := bll.dal.(*DAL.QuestionDAL)
	id = dal.AddTx(model, content)
	return
}

func (bll *QuestionBLL) Update(model Model.Question) {
	bll.dal.Update(model)
}

func (bll *QuestionBLL) Delete(id int) {
	bll.dal.Delete(id)
}

func (bll *QuestionBLL) DeleteModel(model *Model.Question) {
	bll.dal.DeleteModel(model)
}

func (bll *QuestionBLL) GetById(id int) (model *Model.Question) {
	m := bll.dal.GetById(id)
	model = m.(*Model.Question)
	return
}

func (bll *QuestionBLL) GetList(whereStr string) (models []Model.Question) {
	ms := bll.dal.GetList(whereStr)
	models = ms.([]Model.Question)

	return
}

func (bll *QuestionBLL) GetPagingList(pIndex, pSize int, keyOne, keyTwo, sType, sortField, lang, curUserName string) (models []Model.Question, count int) {
	dal := bll.dal.(*DAL.QuestionDAL)
	models, count = dal.GetPagingList(pIndex, pSize, keyOne, keyTwo, sType, sortField, lang, curUserName)

	return
}

func (bll *QuestionBLL) GetPagingListByTag(pIndex, pSize int, keyOne, keyTwo, tagName, sortField, lang, curUserName string) (models []Model.Question, count int) {
	dal := bll.dal.(*DAL.QuestionDAL)
	models, count = dal.GetPagingListByTag(pIndex, pSize, keyOne, keyTwo, tagName, sortField, lang, curUserName)

	return
}

func (bll *QuestionBLL) GetQuestionByVotedAnswer(pIndex, pSize int, curUserName string) (models []Model.Favoritequestionext, count int) {
	dal := bll.dal.(*DAL.QuestionDAL)
	models, count = dal.GetQuestionByVotedAnswer(pIndex, pSize, curUserName)

	return
}

func (bll *QuestionBLL) GetQuestionByFavoriteAnswer(pIndex, pSize int, curUserName string) (models []Model.Favoritequestionext, count int) {
	dal := bll.dal.(*DAL.QuestionDAL)
	models, count = dal.GetQuestionByFavoriteAnswer(pIndex, pSize, curUserName)

	return
}

func (bll *QuestionBLL) GetFixedCount() (count int) {
	dal := bll.dal.(*DAL.QuestionDAL)
	count = dal.GetFixedCount()

	return
}

func (bll *QuestionBLL) GetUserQuestionTop(userName string, top int) (models []Model.Question) {
	dal := bll.dal.(*DAL.QuestionDAL)
	models = dal.GetUserQuestionTop(userName, top)

	return
}
