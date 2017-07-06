package DBDAL

import (
	"DBModel"
	"database/sql"
	"fmt"

	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

type FavoriteTagDAL struct {
	BaseDAL
}

func NewFavoriteTagDAL(driverName, connStr string) (dal *FavoriteTagDAL) {
	dal = &FavoriteTagDAL{}
	dal.ConnectionString = connStr
	dal.DriverName = driverName
	return
}
func (dal FavoriteTagDAL) Add(model interface{}) (id int) {
	m, ok := model.(DBModel.Favoritetag)
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

func (dal FavoriteTagDAL) Update(model interface{}) {
	m, ok := model.(DBModel.Favoritetag)
	if !ok {
		panic("Can not convert model to Favoritetag type")
	}

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Update
	orm := beedb.New(db)
	orm.Save(&m)
}

func (dal FavoriteTagDAL) GetById(id int) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var model DBModel.Favoritetag
	orm.Where(id).Find(&model)
	return &model
}

func (dal FavoriteTagDAL) GetList(whereStr string) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	//beedb.OnDebug = true
	var list []DBModel.Favoritetag
	if whereStr != "" {
		orm.Where(whereStr).FindAll(&list)
	} else {
		orm.FindAll(&list)
	}

	return list
}

func (dal FavoriteTagDAL) GetFavoriteTagsAndUseCounts(pageIndex, pageSize int, userName string) (models []DBModel.Favoritetagcount, count int) {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	//beedb.OnDebug = true
	offset := (pageIndex - 1) * pageSize //from 0
	where := fmt.Sprintf("ft.UserName = '%s'", userName)
	count, err = orm.SetTable("favoritetag ft").Join("", "tag t", "ft.TagId = t.Id").Join("Left", "(Select Tagid, count(0) as TagCounts From questiontag Group by TagId) qt", "ft.TagId = qt.TagId").Where(where).Select("").RowsCount()
	dal.checkError(err)
	err = orm.SetTable("favoritetag ft").Join("", "tag t", "ft.TagId = t.Id").Join("Left", "(Select Tagid, count(0) as TagCounts From questiontag Group by TagId) qt", "ft.TagId = qt.TagId").Where(where).Select("ft.TagId as TagId, ft.Id as FavoriteTagId, t.TagName, qt.TagCounts").Limit(pageSize, offset).FindAll(&models)
	dal.checkError(err)
	return
}

func (dal FavoriteTagDAL) Delete(id int) {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Delete
	orm := beedb.New(db)
	var model DBModel.Favoritetag
	err = orm.Where(id).Find(&model)
	if err != nil {
		panic(err)
	}

	rows, err := orm.Delete(model)
	if rows > 0 && err != nil {
		panic(err)
	}
}

func (dal FavoriteTagDAL) DeleteModel(model interface{}) {
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
