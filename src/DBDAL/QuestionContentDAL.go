package DBDAL

import (
	"DBModel"
	"database/sql"

	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

type QuestionContentDAL struct {
	BaseDAL
}

func NewQuestionContentDAL(driverName, connStr string) (dal *QuestionContentDAL) {
	dal = &QuestionContentDAL{}
	dal.ConnectionString = connStr
	dal.DriverName = driverName
	return
}
func (dal QuestionContentDAL) Add(model interface{}) (id int) {
	m, ok := model.(DBModel.Questioncontent)
	if !ok {
		panic("Can not convert model to QuestionContent type")
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

func (dal QuestionContentDAL) Update(model interface{}) {
	m, ok := model.(DBModel.Questioncontent)
	if !ok {
		panic("Can not convert model to QuestionContent type")
	}

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Update
	orm := beedb.New(db)
	orm.Save(&m)
}

func (dal QuestionContentDAL) GetById(id int) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var model DBModel.Questioncontent
	orm.Where(id).Find(&model)
	return &model
}

func (dal QuestionContentDAL) GetList(whereStr string) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var list []DBModel.Questioncontent
	if whereStr != "" {
		orm.Where(whereStr).FindAll(&list)
	} else {
		orm.FindAll(&list)
	}

	return list
}

func (dal QuestionContentDAL) Delete(id int) {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Delete
	orm := beedb.New(db)
	var model DBModel.Questioncontent
	err = orm.Where(id).Find(&model)
	if err != nil {
		panic(err)
	}

	rows, err := orm.Delete(model)
	if rows > 0 && err != nil {
		panic(err)
	}
}

func (dal QuestionContentDAL) DeleteModel(model interface{}) {
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
