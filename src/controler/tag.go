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

func TagListHandler(w http.ResponseWriter, r *http.Request) {
	common.RenderTemplate(w, r, "tag/tags.html", map[string]interface{}{})
}
