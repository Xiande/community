// DBBLL project DBBLL.go
package DBBLL

import (
	DAL "DBDAL"
	Model "DBModel"
	"fmt"
)

type FavoriteTagBLL struct {
	BaseBLL
}

func NewFavoriteTagBLL(driverName, conn string) (bll *FavoriteTagBLL) {
	bll = new(FavoriteTagBLL)
	//	bll.driverName = driverName
	//	bll.connectionString = conn
	bll.dal = DAL.NewFavoriteTagDAL(driverName, conn)
	return
}

func (bll *FavoriteTagBLL) Add(model Model.Favoritetag) (id int) {
	id = bll.dal.Add(model)
	return
}

func (bll *FavoriteTagBLL) Update(model Model.Favoritetag) {
	bll.dal.Update(model)
}

func (bll *FavoriteTagBLL) Delete(id int) {
	bll.dal.Delete(id)
}

func (bll *FavoriteTagBLL) DeleteModel(model *Model.Favoritetag) {
	bll.dal.DeleteModel(model)
}

func (bll *FavoriteTagBLL) GetById(id int) (model *Model.Favoritetag) {
	m := bll.dal.GetById(id)
	model = m.(*Model.Favoritetag)
	return
}

func (bll *FavoriteTagBLL) GetByTagId(tagId int) (model *Model.Favoritetag) {
	ms := bll.dal.GetList(fmt.Sprintf("TagId = %d", tagId))
	models := ms.([]Model.Favoritetag)
	if models == nil || len(models) == 0 {
		model = nil
	} else {
		model = &models[0]
	}

	return
}

func (bll *FavoriteTagBLL) GetFavoriteTag(userName string, tagId int) (model *Model.Favoritetag) {
	model = nil
	whereStr := fmt.Sprintf("UserName = '%s' And TagId = %d", userName, tagId)
	ms := bll.dal.GetList(whereStr)
	models := ms.([]Model.Favoritetag)
	if len(models) > 0 {
		model = &models[0]
	}
	return
}

func (bll *FavoriteTagBLL) GetList(whereStr string) (models []Model.Favoritetag) {
	ms := bll.dal.GetList(whereStr)
	models = ms.([]Model.Favoritetag)

	return
}

func (bll *FavoriteTagBLL) GetFavoriteTagsAndUseCounts(pageIndex, pageSize int, userName string) (models []Model.Favoritetagcount, count int) {
	dal := bll.dal.(*DAL.FavoriteTagDAL)
	models, count = dal.GetFavoriteTagsAndUseCounts(pageIndex, pageSize, userName)

	return
}
