package controler

import (
	//"DBBLL"
	//"DBModel"
	"common"
	//	"fmt"
	"net/http"
	//"strconv"

	//"github.com/gorilla/mux"
	//"github.com/jimmykuu/wtforms"
)

func TagQuestionHandler(w http.ResponseWriter, r *http.Request) {
	common.RenderTemplate(w, r, "tag/tagquestion.html", map[string]interface{}{})
}
