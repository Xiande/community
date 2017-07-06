package DBDAL

import (
	"DBModel"
	"database/sql"
	"fmt"

	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

type UserDAL struct {
	BaseDAL
}

func NewUserDAL(driverName, connStr string) (dal *UserDAL) {
	dal = &UserDAL{}
	dal.ConnectionString = connStr
	dal.DriverName = driverName
	return
}
func (dal UserDAL) Add(model interface{}) (id int) {
	m, ok := model.(DBModel.User)
	if !ok {
		panic("Can not convert model to User type")
	}

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)
	orm := beedb.New(db)

	//Insert
	err = orm.Save(&m)
	if err != nil {
		panic(err)
	}
	id = m.Id
	return
}

func (dal UserDAL) Update(model interface{}) {
	m, ok := model.(DBModel.User)
	if !ok {
		panic("Can not convert model to User type")
	}

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Update
	orm := beedb.New(db)
	orm.Save(&m)
}

func (dal UserDAL) GetById(id int) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var model DBModel.User
	orm.Where(id).Find(&model)
	return &model
}

func (dal UserDAL) GetByName(name string) *DBModel.User {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	//beedb.OnDebug = true
	var models []DBModel.User
	s := fmt.Sprintf("%s = ?", "Username")
	orm.Where(s, name).FindAll(&models)
	if len(models) > 0 {
		return &models[0]
	}

	return nil
}

func (dal UserDAL) GetByEmail(address string) *DBModel.User {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var models []DBModel.User
	orm.Where(fmt.Sprintf("%s = ?", "Email"), address).Find(&models)
	if len(models) > 0 {
		return &models[0]
	}

	return nil
}

func (dal UserDAL) GetByCode(validateCode string) *DBModel.User {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var models []DBModel.User
	orm.Where(fmt.Sprintf("Where %s = ?", "ValidateCode"), validateCode).Find(&models)
	if len(models) > 0 {
		return &models[0]
	}

	return nil
}

func (dal UserDAL) GetStatistic(userName string) *DBModel.Userstatistic {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	//beedb.OnDebug = true
	var model DBModel.Userstatistic
	model.UserName = userName
	//favorite count
	count, err := orm.SetTable("favoritequestion").Where(fmt.Sprintf("UserName = '%s'", userName)).Select("count(1) as Rows").RowsCount()
	dal.checkError(err)
	model.FavoriteCount = count
	//view count
	count, err = orm.SetTable("question").Where(fmt.Sprintf("UserName = '%s'", userName)).Select("ROUND(SUM(IFNULL(ViewedCount, 0))) as Rows").RowsCount()
	dal.checkError(err)
	model.ViewCount = count
	//vote count
	count, err = orm.SetTable("answer").Where(fmt.Sprintf("UserName = '%s'", userName)).Select("ROUND(SUM(IFNULL(VotedCount, 0))) as Rows").RowsCount()
	dal.checkError(err)
	model.VotedCount = count
	//best count
	count, err = orm.SetTable("answer").Where(fmt.Sprintf("UserName = '%s' And IsBest=1", userName)).Select("count(1) as Rows").RowsCount()
	dal.checkError(err)
	model.BestCount = count
	//expert count
	count, err = orm.SetTable("answer").Where(fmt.Sprintf("UserName = '%s' And IsExpert=1", userName)).Select("count(1) as Rows").RowsCount()
	dal.checkError(err)
	model.ExpertCount = count
	//favorite tag
	count, err = orm.SetTable("favoritetag").Where(fmt.Sprintf("UserName = '%s'", userName)).Select("count(1) as Rows").RowsCount()
	dal.checkError(err)
	model.TagsCount = count

	return &model
}

func (dal UserDAL) GetList(whereStr string) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var users []DBModel.User
	if whereStr != "" {
		orm.Where(whereStr).FindAll(&users)
	} else {
		orm.FindAll(&users)
	}

	return users
}

func (dal UserDAL) Delete(id int) {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Delete
	orm := beedb.New(db)
	var model DBModel.User
	err = orm.Where(id).Find(&model)
	if err != nil {
		panic(err)
	}

	rows, err := orm.Delete(model)
	if rows > 0 && err != nil {
		panic(err)
	}
}

func (dal UserDAL) DeleteModel(model interface{}) {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//m := model.(DBModel.User)
	//Delete
	orm := beedb.New(db)
	rows, err := orm.Delete(model)
	if rows > 0 && err != nil {
		panic(err)
	}
}
