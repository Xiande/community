package DBDAL

import (
	"DBModel"
	"database/sql"
	"fmt"

	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

type FavoriteQuestionDAL struct {
	BaseDAL
}

func NewFavoriteQuestionDAL(driverName, connStr string) (dal *FavoriteQuestionDAL) {
	dal = &FavoriteQuestionDAL{}
	dal.ConnectionString = connStr
	dal.DriverName = driverName
	return
}
func (dal FavoriteQuestionDAL) Add(model interface{}) (id int) {
	m, ok := model.(DBModel.Favoritequestion)
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

func (dal FavoriteQuestionDAL) Update(model interface{}) {
	m, ok := model.(DBModel.Favoritequestion)
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

func (dal FavoriteQuestionDAL) GetById(id int) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var model DBModel.Favoritequestion
	orm.Where(id).Find(&model)
	return &model
}

func (dal FavoriteQuestionDAL) GetFavoriteQuestionList(pIndex, pSize int, curUserName string) (models []DBModel.Favoritequestionext, count int) {

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)
	//beedb.OnDebug = true
	//Select
	orm := beedb.New(db)
	//beedb.OnDebug = true
	offset := (pIndex - 1) * pSize
	where := fmt.Sprintf("favoritequestion.username = '%s'", curUserName)
	count, err = orm.SetTable("favoritequestion").Join("", "question", "favoritequestion.questionid = question.id").Where(where).Select("").RowsCount()

	err = orm.SetTable("favoritequestion").Join("", "question", "favoritequestion.questionid = question.id").Select("question.*, favoritequestion.id as FavoriteQuestionId").Where(where).OrderBy("favoritequestion.createdate desc").Limit(pSize, offset).FindAll(&models)

	if err != nil {
		panic(err)
	}

	return
}

func (dal FavoriteQuestionDAL) GetList(whereStr string) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	//beedb.OnDebug = true
	var list []DBModel.Favoritequestion
	if whereStr != "" {
		orm.Where(whereStr).FindAll(&list)
	} else {
		orm.FindAll(&list)
	}

	return list
}

func (dal FavoriteQuestionDAL) Delete(id int) {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Delete
	orm := beedb.New(db)
	var model DBModel.Favoritequestion
	err = orm.Where(id).Find(&model)
	if err != nil {
		panic(err)
	}

	rows, err := orm.Delete(model)
	if rows > 0 && err != nil {
		panic(err)
	}
}

func (dal FavoriteQuestionDAL) DeleteModel(model interface{}) {
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
