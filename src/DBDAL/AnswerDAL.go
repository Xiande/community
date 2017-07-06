package DBDAL

import (
	"DBModel"
	"database/sql"
	"fmt"
	//	"strings"
	"time"

	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

type AnswerDAL struct {
	BaseDAL
}

func NewAnswerDAL(driverName, connStr string) (dal *AnswerDAL) {
	dal = &AnswerDAL{}
	dal.ConnectionString = connStr
	dal.DriverName = driverName
	return
}
func (dal AnswerDAL) Add(model interface{}) (id int) {
	m, ok := model.(DBModel.Answer)
	if !ok {
		panic("Can not convert model to Answer type")
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

func (dal AnswerDAL) AddTx(model DBModel.Answer, content string) (qid int) {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)
	orm := beedb.New(db)
	tx, err := orm.Db.Begin()
	dal.checkError(err)

	defer tx.Rollback()
	//Insert Answer
	err = orm.SaveTx(&model, tx)
	dal.checkError(err)

	qid = model.Id
	//fmt.Println(qid)
	//Insert content
	var c DBModel.Answercontent
	c.AnswerId = qid
	c.AnswerContent = content
	err = orm.SaveTx(&c, tx)
	dal.checkError(err)

	tx.Commit()

	return
}

func (dal AnswerDAL) Update(model interface{}) {
	m, ok := model.(DBModel.Answer)
	if !ok {
		panic("Can not convert model to Answer type")
	}

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Update
	orm := beedb.New(db)
	orm.Save(&m)
}

func (dal AnswerDAL) GetById(id int) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var model DBModel.Answer
	orm.Where(id).Find(&model)
	return &model
}

func (dal AnswerDAL) GetList(whereStr string) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var list []DBModel.Answer
	if whereStr != "" {
		orm.Where(whereStr).FindAll(&list)
	} else {
		orm.FindAll(&list)
	}

	return list
}

func (dal AnswerDAL) GetLastReply(userName string, top int) (answers []DBModel.Answer) {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	orm.Where(fmt.Sprintf("UserName = '%s'", userName)).OrderBy("CreateDate desc").Limit(top, 0).FindAll(&answers)

	return
}

func (dal AnswerDAL) GetPagingByQuestionId(qid, aid, pageIndex, pageSize int, sortField, sortOrder string) (interface{}, int) {
	count := 0
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	//beedb.OnDebug = true
	var list []DBModel.Answer
	if aid == 0 {
		offset := (pageIndex - 1) * pageSize //from 0
		where := fmt.Sprintf("QuestionId = %d", qid)
		count, err = orm.SetTable("answer").Where(where).Select("").RowsCount()
		orm.Where(where).Limit(pageSize, offset).OrderBy(fmt.Sprintf("%s %s", sortField, sortOrder)).FindAll(&list)
	} else {
		var a DBModel.Answer
		orm.Where(aid).Find(&a)
		list = append(list, a)
		count = 1
	}

	return list, count
}

func (dal AnswerDAL) Delete(id int) {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Delete
	orm := beedb.New(db)
	var model DBModel.Answer
	err = orm.Where(id).Find(&model)
	if err != nil {
		panic(err)
	}

	rows, err := orm.Delete(model)
	if rows > 0 && err != nil {
		panic(err)
	}
}

func (dal AnswerDAL) DeleteModel(model interface{}) {
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

func (dal AnswerDAL) ExpertAnswer(qid, aid int) {
	var (
		experted, answer DBModel.Answer
		question         DBModel.Question
	)

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)
	orm := beedb.New(db)
	tx, err := orm.Db.Begin()
	dal.checkError(err)

	defer tx.Rollback()

	err = orm.Where(fmt.Sprintf("IsExpert = 1 And QuestionId = %d", qid)).Find(&experted)
	if err == nil && experted.Id > 0 && experted.Id != aid {
		experted.IsExpert = false
		experted.ExpertTime = time.Time{}
		orm.SaveTx(experted, tx)
	}

	err = orm.Where(aid).Find(&answer)
	dal.checkError(err)
	answer.IsExpert = true
	answer.ExpertTime = time.Now()
	orm.SaveTx(&answer, tx)

	err = orm.Where(qid).Find(&question)
	dal.checkError(err)
	question.IsExpert = true
	question.ExpertTime = time.Now()
	orm.SaveTx(&question, tx)

	tx.Commit()

	return
}

func (dal AnswerDAL) BestAnswer(qid, aid int) {
	var (
		bested, answer DBModel.Answer
	)

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)
	orm := beedb.New(db)
	tx, err := orm.Db.Begin()
	dal.checkError(err)

	defer tx.Rollback()

	err = orm.Where(fmt.Sprintf("IsBest = 1 And QuestionId = %d", qid)).Find(&bested)
	if err == nil && bested.Id > 0 && bested.Id != aid {
		bested.IsBest = false
		bested.BestTime = time.Time{}
		orm.SaveTx(bested, tx)
	}

	err = orm.Where(aid).Find(&answer)
	dal.checkError(err)
	answer.IsBest = true
	answer.BestTime = time.Now()
	orm.SaveTx(&answer, tx)

	tx.Commit()

	return
}
