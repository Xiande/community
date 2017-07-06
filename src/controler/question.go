package controler

import (
	"DBBLL"
	//"DBModel"
	"common"
	//	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	//"github.com/jimmykuu/wtforms"
)

func QuestionAddHandler(w http.ResponseWriter, r *http.Request) {
	common.RenderTemplate(w, r, "question/add.html", map[string]interface{}{})
}

func QuestionViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	qid := vars["id"]
	id, err := strconv.Atoi(qid)
	if err == nil {
		qbll := DBBLL.NewQuestionBLL(common.Config.DB, common.Config.DBConn)
		q := qbll.GetById(id)
		q.ViewedCount++
		qbll.Update(*q)
	} else {
		common.WriteLog(err.Error())
	}

	//fmt.Println(qid)
	common.RenderTemplate(w, r, "question/view.html", map[string]interface{}{})
}

func QuestionSearchResultHandler(w http.ResponseWriter, r *http.Request) {
	common.RenderTemplate(w, r, "question/searchresult.html", map[string]interface{}{})
}
