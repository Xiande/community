package controler

import (
	"DBBLL"
	"DBModel"
	//"bytes"
	"common"
	"encoding/json"
	//	"fmt"
	//"io/ioutil"
	"net/http"
	//"net/url"
	//"strconv"
	//"time"
	"viewmodel"
)

func UserAjaxHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == http.MethodGet {
		op := r.Form.Get("op")

		//fmt.Println(op)
		switch op {
		case "followuser":
			userName := r.Form.Get("itcode")
			followUser(w, r, userName)
		case "stopfollowuser":
			userName := r.Form.Get("itcode")
			stopFollowUser(w, r, userName)
		case "statistics":
			getUserstatistic(w, r)
		}
	}

	defer func() {
		if x := recover(); x != nil {
			common.WriteLog(x.(error).Error())
		}
	}()
}

func followUser(w http.ResponseWriter, r *http.Request, userName string) {
	fuBll := DBBLL.NewFollowUserBLL(common.Config.DB, common.Config.DBConn)
	ret := "FollowUserSuccess"
	curUser, ok := common.CurrentUser(r)
	if ok {
		if curUser.Username == userName {
			ret = "CannotFollowYourSelf"
		} else {
			fa := fuBll.GetRecord(curUser.Username, userName)

			if fa != nil {
				ret = "HadFollowedUser"
			} else {
				var model DBModel.Followuser
				model.UserName = curUser.Username
				model.FollowedUserName = userName
				fuBll.Add(model)
			}
		}
	} else {
		ret = "SystemError"
	}

	w.Write([]byte(ret))
}
func stopFollowUser(w http.ResponseWriter, r *http.Request, userName string) {
	fuBll := DBBLL.NewFollowUserBLL(common.Config.DB, common.Config.DBConn)
	ret := "StopFollowSuccess"
	curUser, ok := common.CurrentUser(r)
	if ok {
		fa := fuBll.GetRecord(curUser.Username, userName)

		if fa != nil {
			fuBll.DeleteModel(fa)
			ret = "StopFollowSuccess"
		} else {
			ret = "NotFollowing"
		}
	} else {
		ret = "StopfollowUserError"
	}

	w.Write([]byte(ret))
}

func getUserstatistic(w http.ResponseWriter, r *http.Request) {
	uBll := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn)

	var vm viewmodel.UserStatistic
	curUser, ok := common.CurrentUser(r)
	if ok {
		m := uBll.GetStatistic(curUser.Username)
		vm.Userstatistic = *m
		vm.PhotoStr = curUser.PhotoImgSrc()
		vm.BadgesCount = 0
	}

	ret, _ := json.Marshal(vm)
	w.Write([]byte(ret))
}
