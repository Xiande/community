// DBBLL project DBBLL.go
package DBBLL

import (
	DAL "DBDAL"
	Model "DBModel"
)

type TagBLL struct {
	BaseBLL
}

func NewTagBLL(driverName, conn string) (bll *TagBLL) {
	bll = new(TagBLL)
	//	bll.driverName = driverName
	//	bll.connectionString = conn
	bll.dal = DAL.NewTagDAL(driverName, conn)
	return
}

func (bll *TagBLL) Add(model Model.Tag) (id int) {
	id = bll.dal.Add(model)
	return
}

func (bll *TagBLL) Update(model Model.Tag) {
	bll.dal.Update(model)
}

func (bll *TagBLL) Delete(id int) {
	bll.dal.Delete(id)
}

func (bll *TagBLL) DeleteModel(model *Model.Tag) {
	bll.dal.DeleteModel(model)
}

func (bll *TagBLL) GetById(id int) (model *Model.Tag) {
	m := bll.dal.GetById(id)
	model = m.(*Model.Tag)
	return
}

func (bll *TagBLL) GetByName(name string) (model *Model.Tag) {
	dal := bll.dal.(*DAL.TagDAL)
	m := dal.GetByName(name)
	model = m.(*Model.Tag)
	return
}

func (bll *TagBLL) FilterTag(name string) (names string) {
	dal := bll.dal.(*DAL.TagDAL)
	names = dal.FilterTag(name)

	return
}

func (bll *TagBLL) GetTagsPaging(tagName, userName string, pageIndex, pageSize int) (models []Model.TagExt, count int) {
	dal := bll.dal.(*DAL.TagDAL)
	models, count = dal.GetTagsPaging(tagName, userName, pageIndex, pageSize)

	return
}

func (bll *TagBLL) GetList(whereStr string) (models []Model.Tag) {
	ms := bll.dal.GetList(whereStr)
	models = ms.([]Model.Tag)

	return
}
