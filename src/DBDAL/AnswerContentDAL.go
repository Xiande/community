package DBDAL

import (
	"DBModel"
	"database/sql"

	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

type AnswerContentDAL struct {
	BaseDAL
}

func NewAnswerContentDAL(driverName, connStr string) (dal *AnswerContentDAL) {
	dal = &AnswerContentDAL{}
	dal.ConnectionString = connStr
	dal.DriverName = driverName
	return
}
func (dal AnswerContentDAL) Add(model interface{}) (id int) {
	m, ok := model.(DBModel.Answercontent)
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

func (dal AnswerContentDAL) Update(model interface{}) {
	m, ok := model.(DBModel.Answercontent)
	if !ok {
		panic("Can not convert model to AnswerContent type")
	}

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Update
	orm := beedb.New(db)
	orm.Save(&m)
}

func (dal AnswerContentDAL) GetById(id int) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var model DBModel.Answercontent
	orm.Where(id).Find(&model)
	return &model
}

func (dal AnswerContentDAL) GetList(whereStr string) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var list []DBModel.Answercontent
	if whereStr != "" {
		orm.Where(whereStr).FindAll(&list)
	} else {
		orm.FindAll(&list)
	}

	return list
}

func (dal AnswerContentDAL) Delete(id int) {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Delete
	orm := beedb.New(db)
	var model DBModel.Answercontent
	err = orm.Where(id).Find(&model)
	if err != nil {
		panic(err)
	}

	rows, err := orm.Delete(model)
	if rows > 0 && err != nil {
		panic(err)
	}
}

func (dal AnswerContentDAL) DeleteModel(model interface{}) {
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
