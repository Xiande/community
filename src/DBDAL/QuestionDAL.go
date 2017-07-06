package DBDAL

import (
	"DBModel"
	"database/sql"
	"fmt"
	"strings"

	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

type QuestionDAL struct {
	BaseDAL
}

func NewQuestionDAL(driverName, connStr string) (dal *QuestionDAL) {
	dal = &QuestionDAL{}
	dal.ConnectionString = connStr
	dal.DriverName = driverName
	return
}
func (dal QuestionDAL) Add(model interface{}) (id int) {
	m, ok := model.(DBModel.Question)
	if !ok {
		panic("Can not convert model to Question type")
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

func (dal QuestionDAL) AddTx(model DBModel.Question, content string) (qid int) {

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)
	orm := beedb.New(db)
	tx, err := orm.Db.Begin()
	dal.checkError(err)

	//Insert question
	err = orm.SaveTx(&model, tx)
	dal.checkError(err)

	qid = model.Id
	//fmt.Println(qid)
	//Insert content
	var c DBModel.Questioncontent
	c.QuestionId = qid
	c.QuestionContent = content
	err = orm.SaveTx(&c, tx)
	dal.checkError(err)
	//Insert tag, questiontag
	var allTags []DBModel.Tag
	orm.FindAll(&allTags)
	inputTags := strings.Split(model.Tags, ",")
	for _, name := range inputTags {
		exist := false
		var dbTagId int
		for _, tag := range allTags {
			if strings.ToLower(tag.TagName) == strings.ToLower(name) {
				exist = true
				dbTagId = tag.Id
				break
			}
		}

		if !exist {
			var newTag DBModel.Tag
			newTag.TagName = name
			err = orm.SaveTx(&newTag, tx)
			dal.checkError(err)
			dbTagId = newTag.Id
		}

		var qt DBModel.Questiontag
		qt.QuestionId = qid
		qt.TagId = dbTagId
		err = orm.SaveTx(&qt, tx)
		dal.checkError(err)
	}

	defer tx.Rollback()

	tx.Commit()

	return
}

func (dal QuestionDAL) Update(model interface{}) {
	m, ok := model.(DBModel.Question)
	if !ok {
		panic("Can not convert model to Question type")
	}

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Update
	orm := beedb.New(db)
	orm.Save(&m)
}

func (dal QuestionDAL) GetById(id int) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var model DBModel.Question
	orm.Where(id).Find(&model)
	return &model
}

func (dal QuestionDAL) GetList(whereStr string) interface{} {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	var list []DBModel.Question
	if whereStr != "" {
		orm.Where(whereStr).FindAll(&list)
	} else {
		orm.FindAll(&list)
	}

	return list
}

func (dal QuestionDAL) GetPagingList(pIndex, pSize int, keyOne, keyTwo, sType, sortField, lang, curUserName string) (models []DBModel.Question, count int) {

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	//beedb.OnDebug = true
	offset := (pIndex - 1) * pSize
	where := ""
	if lang != "all" {
		where += fmt.Sprintf(" LanguageType = '%s' And", lang)
	}

	if keyOne != "" {
		where += fmt.Sprintf(" Title like '%%%s%%' And", keyOne)
	}

	if keyTwo != "" {
		where += fmt.Sprintf(" Title like '%%%s%%' And", keyTwo)
	}

	sort := "CreateDate Desc"
	if sortField != "" {
		switch sortField {
		case "Expert":
			where += " IsExpert = 1 And"
		case "NoneAnswer":
			where += " AnswersCount = 0 And"
		default:
			sort = sortField + " Desc"
		}
	}

	switch sType {
	case "my":
		where += fmt.Sprintf(" UserName = '%s'", curUserName)
		count, err = orm.SetTable("question").Where(where).Select("").RowsCount()
		err = orm.Where(where).OrderBy(sort).Limit(pSize, offset).FindAll(&models)
	case "join":
		where += fmt.Sprintf(" answer.UserName = '%s'", curUserName)
		count, err = orm.SetTable("question").Join("", "answer", "question.Id = answer.QuestionId").Where(where).Select("count(distinct question.Id) as RowsCount").RowsCount()
		err = orm.SetTable("question").Join("", "answer", "question.Id = answer.QuestionId").Where(where).Select("distinct question.*").OrderBy(sort).FindAll(&models)
	case "follow":
		where += fmt.Sprintf(" followuser.UserName = '%s'", curUserName)
		count, err = orm.SetTable("question").Join("", "followuser", "question.UserName = followuser.FollowedUserName").Where(where).Select("").RowsCount()
		err = orm.SetTable("question").Join("", "followuser", "question.UserName = followuser.FollowedUserName").Where(where).Select("question.*").OrderBy(sort).FindAll(&models)
	default:
		if where != "" {
			where = strings.Trim(where, "And")
			orm.WhereStr = where
		}
		//fmt.Println(where)
		count, err = orm.SetTable("question").Where(where).Select("").RowsCount()
		err = orm.Where(where).OrderBy(sort).Limit(pSize, offset).FindAll(&models)
	}

	dal.checkError(err)

	return
}
func (dal QuestionDAL) GetPagingListByTag(pIndex, pSize int, keyOne, keyTwo, tagName, sortField, lang, curUserName string) (models []DBModel.Question, count int) {

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	//beedb.OnDebug = true
	offset := (pIndex - 1) * pSize

	var tag DBModel.Tag
	orm.Where("TagName = ?", tagName).Find(&tag)
	if tag.Id <= 0 {
		panic("not found tag: " + tagName)
	}

	where := fmt.Sprintf(" questiontag.TagId = %d And", tag.Id)
	if lang != "all" {
		where += fmt.Sprintf(" question.LanguageType = '%s' And", lang)
	}

	if keyOne != "" {
		where += fmt.Sprintf(" question.Title like '%%%s%%' And", keyOne)
	}

	if keyTwo != "" {
		where += fmt.Sprintf(" question.Title like '%%%s%%' And", keyTwo)
	}

	sort := "question.CreateDate Desc"
	if sortField != "" {
		switch sortField {
		case "Expert":
			where += " question.IsExpert = 1 And"
		case "NoneAnswer":
			where += " question.AnswersCount = 0 And"
		default:
			sort = sortField + " Desc"
		}
	}

	if where != "" {
		where = strings.Trim(where, "And")
		orm.WhereStr = where
	}
	//fmt.Println(where)

	count, err = orm.SetTable("question").Join("", "questiontag", "question.Id = questiontag.QuestionId").Where(where).Select("").RowsCount()
	dal.checkError(err)

	err = orm.SetTable("question").Join("", "questiontag", "question.Id = questiontag.QuestionId").Where(where).Select("question.*").OrderBy(sort).Limit(pSize, offset).FindAll(&models)

	dal.checkError(err)

	return
}

func (dal QuestionDAL) GetUserQuestionTop(userName string, top int) (models []DBModel.Question) {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Select
	orm := beedb.New(db)
	//beedb.OnDebug = true

	sort := "CreateDate Desc"

	err = orm.SetTable("question").Join("", "followuser", "question.UserName = followuser.FollowedUserName").Where("followuser.UserName = ? ", userName).Limit(top).Select("question.*").OrderBy(sort).FindAll(&models)

	if err != nil {
		panic(err)
	}

	return
}

func (dal QuestionDAL) GetQuestionByVotedAnswer(pIndex, pSize int, curUserName string) (models []DBModel.Favoritequestionext, count int) {

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)
	//beedb.OnDebug = true
	//Select
	orm := beedb.New(db)
	//beedb.OnDebug = true
	offset := (pIndex - 1) * pSize
	where := fmt.Sprintf("avr.username = '%s'", curUserName)
	count, err = orm.SetTable("question q").Join("", "answer a", "q.Id = a.QuestionId").Join("", "answervotedrecord avr", "a.Id = avr.AnswerId").Where(where).Select("").RowsCount()

	err = orm.SetTable("question q").Join("", "answer a", "q.Id = a.QuestionId").Join("", "answervotedrecord avr", "a.Id = avr.AnswerId").Where(where).OrderBy("q.createdate desc").Select("q.*, a.Id as FavoriteAnswerId").Limit(pSize, offset).FindAll(&models)

	if err != nil {
		panic(err)
	}

	return
}

func (dal QuestionDAL) GetQuestionByFavoriteAnswer(pIndex, pSize int, curUserName string) (models []DBModel.Favoritequestionext, count int) {

	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)
	//beedb.OnDebug = true
	//Select
	orm := beedb.New(db)
	//beedb.OnDebug = true
	offset := (pIndex - 1) * pSize
	where := fmt.Sprintf("fa.username = '%s'", curUserName)
	count, err = orm.SetTable("question q").Join("", "answer a", "q.Id = a.QuestionId").Join("", "favoriteanswer fa", "a.Id = fa.AnswerId").Where(where).Select("").RowsCount()

	err = orm.SetTable("question q").Join("", "answer a", "q.Id = a.QuestionId").Join("", "favoriteanswer fa", "a.Id = fa.AnswerId").Where(where).OrderBy("q.createdate desc").Select("q.*, a.Id as AnswerId, fa.Id as FavoriteAnswerId").Limit(pSize, offset).FindAll(&models)

	if err != nil {
		panic(err)
	}

	return
}

func (dal QuestionDAL) Delete(id int) {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//Delete
	orm := beedb.New(db)

	tx, err := orm.Db.Begin()
	dal.checkError(err)

	defer tx.Rollback()
	//beedb.OnDebug = true
	var model DBModel.Question
	err = orm.Where(id).Find(&model)
	dal.checkError(err)

	//delete question tags
	var qts []DBModel.Questiontag
	err = orm.Where(fmt.Sprintf("QuestionId = %d", id)).FindAll(&qts)
	dal.checkError(err)

	for _, qt := range qts {
		rows, err := orm.DeleteTx(qt, tx)
		if rows > 0 && err != nil {
			panic(err)
		}
	}

	//delete question content
	var qc DBModel.Questioncontent
	err = orm.Where(fmt.Sprintf("QuestionId = %d", id)).Find(&qc)
	dal.checkError(err)
	rows, err := orm.DeleteTx(qc, tx)
	if rows > 0 && err != nil {
		panic(err)
	}

	//delete favorite questions
	var fqs []DBModel.Favoritequestion
	err = orm.Where(fmt.Sprintf("QuestionId = %d", id)).FindAll(&fqs)
	dal.checkError(err)
	for _, fq := range fqs {
		rows, err = orm.DeleteTx(fq, tx)
		if rows > 0 && err != nil {
			panic(err)
		}
	}

	//delete answers
	var as []DBModel.Answer
	err = orm.Where(fmt.Sprintf("QuestionId = %d", id)).FindAll(&as)
	dal.checkError(err)
	for _, a := range as {
		//delete answer content
		var ac DBModel.Answercontent
		err = orm.Where(fmt.Sprintf("AnswerId = %d", a.Id)).Find(&ac)
		dal.checkError(err)
		rows, err = orm.DeleteTx(ac, tx)
		if rows > 0 && err != nil {
			panic(err)
		}

		//delete favorite answer
		var fas []DBModel.Favoriteanswer
		err = orm.Where(fmt.Sprintf("AnswerId = %d", a.Id)).FindAll(&fas)
		dal.checkError(err)
		for _, fa := range fas {
			rows, err = orm.DeleteTx(fa, tx)
			if rows > 0 && err != nil {
				panic(err)
			}
		}

		//delete answer voted record
		var avrs []DBModel.Answervotedrecord
		err = orm.Where(fmt.Sprintf("AnswerId = %d", a.Id)).FindAll(&avrs)
		dal.checkError(err)
		for _, avr := range avrs {
			rows, err = orm.DeleteTx(avr, tx)
			if rows > 0 && err != nil {
				panic(err)
			}
		}

		rows, err = orm.DeleteTx(a, tx)
		if rows > 0 && err != nil {
			panic(err)
		}
	}

	rows, err = orm.DeleteTx(model, tx)
	if rows > 0 && err != nil {
		panic(err)
	}

	tx.Commit()

}

func (dal QuestionDAL) GetFixedCount() int {
	db, err := sql.Open(dal.DriverName, dal.ConnectionString)
	defer db.Close()
	dal.checkError(err)

	//m := model.(DBModel.User)
	//Delete
	orm := beedb.New(db)
	count, err := orm.SetTable("question").Join("", "answer", "question.id = answer.QuestionId").Select("count(distinct question.Id) as Rows").RowsCount()
	if err != nil {
		panic(err)
	}

	return count
}

func (dal QuestionDAL) DeleteModel(model interface{}) {
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
