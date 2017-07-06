package DBDAL

import (
	"DBModel"
	"database/sql"
	//	"fmt"

	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

type AnswerVotedRecordDAL struct {
	BaseDAL
}

func NewAnswerVotedRecordDAL(driverName, connStr string) (dal *AnswerVotedRecordDAL) {
	dal = &AnswerVotedRecordDAL{}
	dal.ConnectionString = connStr
	dal.DriverName = driverName
	return
}
func (dal AnswerVotedRecordDAL) Add(model interface{}) (id int) {
	m, ok := model.(DBModel.Answervotedrecord)
	if !ok {
		panic("Can not convert model to AnswerVotedRecord type")
	}

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)
	orm := beedb.New(db)
	tx, err := orm.Db.Begin()
	dal.checkError(err)
	defer tx.Rollback()

	//Insert
	err = orm.SaveTx(&m, tx)
	dal.checkError(err)
	id = m.Id

	var q DBModel.Question
	err = orm.Where(m.QuestionId).Find(&q)
	dal.checkError(err)
	q.VotedCount += 1
	orm.SaveTx(&q, tx)

	var a DBModel.Answer
	err = orm.Where(m.AnswerId).Find(&a)
	dal.checkError(err)
	a.VotedCount += 1
	orm.SaveTx(&a, tx)

	tx.Commit()

	return
}

func (dal AnswerVotedRecordDAL) Update(model interface{}) {
	m, ok := model.(DBModel.Answervotedrecord)
	if !ok {
		panic("Can not convert model to Answervotedrecord type")
	}

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Update
	orm := beedb.New(db)
	orm.Save(&m)
}

func (dal AnswerVotedRecordDAL) GetById(id int) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var model DBModel.Answervotedrecord
	orm.Where(id).Find(&model)
	return &model
}

func (dal AnswerVotedRecordDAL) GetList(whereStr string) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	//beedb.OnDebug = true
	var list []DBModel.Answervotedrecord
	if whereStr != "" {
		orm.Where(whereStr).FindAll(&list)
	} else {
		orm.FindAll(&list)
	}

	return list
}

func (dal AnswerVotedRecordDAL) Delete(id int) {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Delete
	orm := beedb.New(db)
	var model DBModel.Answervotedrecord
	err = orm.Where(id).Find(&model)
	if err != nil {
		panic(err)
	}

	rows, err := orm.Delete(model)
	if rows > 0 && err != nil {
		panic(err)
	}
}

func (dal AnswerVotedRecordDAL) DeleteModel(model interface{}) {
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
