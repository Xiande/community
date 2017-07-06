// DBBLL project DBBLL.go
package DBBLL

import (
	DAL "DBDAL"
	Model "DBModel"
)

type UserBLL struct {
	BaseBLL
}

func NewUserBLL(driverName, conn string) (bll *UserBLL) {
	bll = new(UserBLL)
	//	bll.driverName = driverName
	//	bll.connectionString = conn
	bll.dal = DAL.NewUserDAL(driverName, conn)
	return
}

func (bll *UserBLL) Add(model Model.User) (id int) {
	id = bll.dal.Add(model)
	return
}

func (bll *UserBLL) Update(model Model.User) {
	bll.dal.Update(model)
}

func (bll *UserBLL) Delete(id int) {
	bll.dal.Delete(id)
}

func (bll *UserBLL) DeleteModel(model *Model.User) {
	bll.dal.DeleteModel(model)
}

func (bll *UserBLL) GetById(id int) (model *Model.User) {
	m := bll.dal.GetById(id)
	model = m.(*Model.User)
	return
}

func (bll *UserBLL) GetByName(name string) (model *Model.User) {
	dal := bll.dal.(*DAL.UserDAL)
	model = dal.GetByName(name)

	return
}

func (bll *UserBLL) GetByEmail(address string) (model *Model.User) {
	dal := bll.dal.(*DAL.UserDAL)
	model = dal.GetByEmail(address)

	return
}
func (bll *UserBLL) GetByCode(validateCode string) (model *Model.User) {
	dal := bll.dal.(*DAL.UserDAL)
	model = dal.GetByCode(validateCode)

	return
}
func (bll *UserBLL) GetStatistic(userName string) (model *Model.Userstatistic) {
	dal := bll.dal.(*DAL.UserDAL)
	model = dal.GetStatistic(userName)

	return
}
func (bll *UserBLL) GetList(whereStr string) (models []Model.User) {
	ms := bll.dal.GetList(whereStr)
	models = ms.([]Model.User)

	return
}
