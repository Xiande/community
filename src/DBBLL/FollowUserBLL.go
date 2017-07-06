// DBBLL project DBBLL.go
package DBBLL

import (
	DAL "DBDAL"
	Model "DBModel"
	"fmt"
)

type FollowUserBLL struct {
	BaseBLL
}

func NewFollowUserBLL(driverName, conn string) (bll *FollowUserBLL) {
	bll = new(FollowUserBLL)
	//	bll.driverName = driverName
	//	bll.connectionString = conn
	bll.dal = DAL.NewFollowUserDAL(driverName, conn)
	return
}

func (bll *FollowUserBLL) Add(model Model.Followuser) (id int) {
	id = bll.dal.Add(model)
	return
}

func (bll *FollowUserBLL) Update(model Model.Followuser) {
	bll.dal.Update(model)
}

func (bll *FollowUserBLL) Delete(id int) {
	bll.dal.Delete(id)
}

func (bll *FollowUserBLL) DeleteModel(model *Model.Followuser) {
	bll.dal.DeleteModel(model)
}

func (bll *FollowUserBLL) GetById(id int) (model *Model.Followuser) {
	m := bll.dal.GetById(id)
	model = m.(*Model.Followuser)
	return
}

func (bll *FollowUserBLL) GetRecord(userName, followedUserName string) (model *Model.Followuser) {
	whereStr := fmt.Sprintf("UserName = '%s' And FollowedUserName = '%s'", userName, followedUserName)
	ms := bll.GetList(whereStr)
	if len(ms) > 0 {
		model = &ms[0]
	} else {
		model = nil
	}

	return
}

func (bll *FollowUserBLL) GetList(whereStr string) (models []Model.Followuser) {
	ms := bll.dal.GetList(whereStr)
	models = ms.([]Model.Followuser)

	return
}
