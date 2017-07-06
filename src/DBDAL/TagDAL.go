package DBDAL

import (
	"DBModel"
	"database/sql"
	"fmt"

	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

type TagDAL struct {
	BaseDAL
}

func NewTagDAL(driverName, connStr string) (dal *TagDAL) {
	dal = &TagDAL{}
	dal.ConnectionString = connStr
	dal.DriverName = driverName
	return
}
func (dal TagDAL) Add(model interface{}) (id int) {
	m, ok := model.(DBModel.Tag)
	if !ok {
		panic("Can not convert model to Tag type")
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

func (dal TagDAL) Update(model interface{}) {
	m, ok := model.(DBModel.Tag)
	if !ok {
		panic("Can not convert model to Tag type")
	}

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Update
	orm := beedb.New(db)
	orm.Save(&m)
}

func (dal TagDAL) GetById(id int) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var model DBModel.Tag
	orm.Where(id).Find(&model)
	return &model
}

func (dal TagDAL) GetByName(name string) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	//beedb.OnDebug = true
	var model DBModel.Tag
	s := fmt.Sprintf("%s = ?", "TagName")
	orm.Where(s, name).Find(&model)
	return &model
}

func (dal TagDAL) FilterTag(key string) (names string) {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var tags []DBModel.Tag
	whereStr := "TagName like '" + key + "%'"
	orm.Where(whereStr).FindAll(&tags)

	for _, m := range tags {
		names = names + m.TagName + "\r\n"
	}

	return
}
func (dal TagDAL) GetTagsPaging(tagName, userName string, pageIndex, pageSize int) (tags []DBModel.TagExt, count int) {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	//beedb.OnDebug = true
	whereStr := ""
	if tagName != "" {
		whereStr += fmt.Sprintf("v.TagName = '%s'", tagName)
	}

	offset := (pageIndex - 1) * pageSize //from 0
	count, err = orm.SetTable("vtagquestiontotal v").Join("Left", fmt.Sprintf("(Select * From favoritetag Where UserName = '%s') as t", userName), "v.id = t.TagId").Where(whereStr).Select("").RowsCount()
	orm.SetTable("vtagquestiontotal v").Join("Left", fmt.Sprintf("(Select * From favoritetag Where UserName = '%s') as t", userName), "v.id = t.TagId").Where(whereStr).OrderBy("v.TagQuestionCount Desc").Select("v.*, (t.Id is not null) as IsMyFavorite").Limit(pageSize, offset).FindAll(&tags)
	return
}

func (dal TagDAL) GetList(whereStr string) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var tags []DBModel.Tag
	if whereStr != "" {
		orm.Where(whereStr).FindAll(&tags)

	} else {
		orm.FindAll(&tags)
	}

	return tags
}

func (dal TagDAL) Delete(id int) {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Delete
	orm := beedb.New(db)
	var model DBModel.Tag
	err = orm.Where(id).Find(&model)
	if err != nil {
		panic(err)
	}

	rows, err := orm.Delete(model)
	if rows > 0 && err != nil {
		panic(err)
	}
}

func (dal TagDAL) DeleteModel(model interface{}) {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//m := model.(DBModel.Tag)
	//Delete
	orm := beedb.New(db)
	rows, err := orm.Delete(model)
	if rows > 0 && err != nil {
		panic(err)
	}
}
