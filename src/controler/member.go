package controler

import (
	"DBBLL"
	"common"
	"net/http"
	"viewmodel"

	"github.com/gorilla/mux"
)

func PersonalHandler(w http.ResponseWriter, r *http.Request) {
	uBll := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn)
	fuBll := DBBLL.NewFollowUserBLL(common.Config.DB, common.Config.DBConn)
	vars := mux.Vars(r)
	username := vars["username"]
	var userInfo viewmodel.UserInfo
	userInfo.Base = uBll.GetByName(username)
	userInfo.Statistic = uBll.GetStatistic(username)
	curUser, ok := common.CurrentUser(r)
	if ok {
		fur := fuBll.GetRecord(curUser.Username, username)
		userInfo.IsFollowed = fur != nil
	}

	common.RenderTemplate(w, r, "/account/info.html", map[string]interface{}{
		"userinfo": &userInfo,
		"active":   "members",
	})
}
func PersonalQuestionHandler(w http.ResponseWriter, r *http.Request) {
	uBll := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn)
	fuBll := DBBLL.NewFollowUserBLL(common.Config.DB, common.Config.DBConn)
	vars := mux.Vars(r)
	username := vars["username"]
	var userInfo viewmodel.UserInfo
	userInfo.Base = uBll.GetByName(username)
	userInfo.Statistic = uBll.GetStatistic(username)
	curUser, ok := common.CurrentUser(r)
	if ok {
		fur := fuBll.GetRecord(curUser.Username, username)
		userInfo.IsFollowed = fur != nil
	}

	common.RenderTemplate(w, r, "/account/userquestions.html", map[string]interface{}{
		"userinfo": &userInfo,
		"active":   "members",
	})
}
