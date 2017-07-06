package DBDAL

import (
	"DBModel"
	"database/sql"
	//	"fmt"

	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

type FavoriteAnswerDAL struct {
	BaseDAL
}

func NewFavoriteAnswerDAL(driverName, connStr string) (dal *FavoriteAnswerDAL) {
	dal = &FavoriteAnswerDAL{}
	dal.ConnectionString = connStr
	dal.DriverName = driverName
	return
}
func (dal FavoriteAnswerDAL) Add(model interface{}) (id int) {
	m, ok := model.(DBModel.Favoriteanswer)
	if !ok {
		panic("Can not convert model to AnswerContent type")
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

func (dal FavoriteAnswerDAL) Update(model interface{}) {
	m, ok := model.(DBModel.Favoriteanswer)
	if !ok {
		panic("Can not convert model to Favoriteanswer type")
	}

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Update
	orm := beedb.New(db)
	orm.Save(&m)
}

func (dal FavoriteAnswerDAL) GetById(id int) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var model DBModel.Favoriteanswer
	orm.Where(id).Find(&model)
	return &model
}

func (dal FavoriteAnswerDAL) GetList(whereStr string) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	//beedb.OnDebug = true
	var list []DBModel.Favoriteanswer
	if whereStr != "" {
		orm.Where(whereStr).FindAll(&list)
	} else {
		orm.FindAll(&list)
	}

	return list
}

func (dal FavoriteAnswerDAL) Delete(id int) {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Delete
	orm := beedb.New(db)
	var model DBModel.Favoriteanswer
	err = orm.Where(id).Find(&model)
	if err != nil {
		panic(err)
	}

	rows, err := orm.Delete(model)
	if rows > 0 && err != nil {
		panic(err)
	}
}

func (dal FavoriteAnswerDAL) DeleteModel(model interface{}) {
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
