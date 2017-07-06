package controler

import (
	"encoding/json"
	//	"fmt"
	"net/http"
	"strconv"
	//"strings"
	"DBBLL"
	"DBModel"
	"common"
	"time"
	"viewmodel"
	//	"io"
	//"bytes"
	//"encoding/base64"
	//"io/ioutil"
	//"net/url"
	//"os"
	//"path"
	//"strings"

	//"code.google.com/p/go-uuid/uuid"
)

func TagAjaxHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		r.ParseForm()
		op := r.Form.Get("op")
		//fmt.Println(op)
		switch op {
		case "gettagbyname":
			tName := r.Form.Get("tName")
			pIndex, _ := strconv.Atoi(r.Form.Get("pIndex"))
			pSize, _ := strconv.Atoi(r.Form.Get("pSize"))
			js := getTagsByName(r, tName, pIndex, pSize)
			w.Write(js)

		case "filtertag":
			key := r.Form.Get("q")
			//fmt.Println(key)
			bll := DBBLL.NewTagBLL(common.Config.DB, common.Config.DBConn)
			names := bll.FilterTag(key)
			//js := objToJson(names)
			w.Write([]byte(names))
		case "followtag":
			id := r.Form.Get("tagID")
			tagID, _ := strconv.Atoi(id)
			//fmt.Println(key)
			ret := followTag(r, tagID)
			//js := objToJson(names)
			w.Write([]byte(ret))
		case "removetag":
			id := r.Form.Get("favtagID")
			favtagID, _ := strconv.Atoi(id)
			//fmt.Println(key)
			ret := removeFollowTag(r, favtagID)
			//js := objToJson(names)
			w.Write([]byte(ret))
		case "tagques":
			pageIndex := r.Form.Get("pIndex")
			pageSize := r.Form.Get("pSize")
			keyOne := r.Form.Get("keyOne")
			keyTwo := r.Form.Get("keyTwo")
			tag := r.Form.Get("title")
			sort := r.Form.Get("sort")
			lang := r.Form.Get("lang")
			pIndex, _ := strconv.Atoi(pageIndex)
			pSize, _ := strconv.Atoi(pageSize)
			getQuestionPageByTag(w, r, pIndex, pSize, keyOne, keyTwo, tag, sort, lang)
		case "followedtag":
			pageIndex := r.Form.Get("pIndex")
			pageSize := r.Form.Get("pSize")
			pIndex, _ := strconv.Atoi(pageIndex)
			pSize, _ := strconv.Atoi(pageSize)
			ret := getFollowedTag(r, pIndex, pSize)
			w.Write([]byte(ret))
		}
	}
}

func getTagsByName(r *http.Request, name string, pageIndex, pageSize int) []byte {
	tbll := DBBLL.NewTagBLL(common.Config.DB, common.Config.DBConn)
	var exts []DBModel.TagExt
	curUser, ok := common.CurrentUser(r)
	userName := ""
	if ok {
		userName = curUser.Username
	}

	exts, count := tbll.GetTagsPaging(name, userName, pageIndex, pageSize)

	var ret viewmodel.PagingList
	ret.ResultList = exts
	ret.TotalRows = count
	js, err := json.Marshal(ret)
	//fmt.Println(string(js))
	//fmt.Println(ret.TotalRows)
	if err != nil {
		common.WriteLog(err.Error())
	}

	return js
}

func getFollowedTag(r *http.Request, pageIndex, pageSize int) []byte {
	ftBll := DBBLL.NewFavoriteTagBLL(common.Config.DB, common.Config.DBConn)
	curUser, ok := common.CurrentUser(r)
	if ok {
		list, count := ftBll.GetFavoriteTagsAndUseCounts(pageIndex, pageSize, curUser.Username)
		var ret viewmodel.PagingList
		ret.ResultList = list
		ret.TotalRows = count
		js, err := json.Marshal(ret)
		//fmt.Println(string(js))
		//fmt.Println(ret.TotalRows)
		if err != nil {
			common.WriteLog(err.Error())
		}

		return js
	}

	return []byte("[]")
}

func followTag(r *http.Request, tagID int) []byte {
	bll := DBBLL.NewFavoriteTagBLL(common.Config.DB, common.Config.DBConn)
	curUser, ok := common.CurrentUser(r)
	ret := "FollowTagSuccess"
	if ok {
		exist := bll.GetFavoriteTag(curUser.Username, tagID)
		if exist != nil {
			ret = "HadFollowTag"
		} else {
			var ft DBModel.Favoritetag
			ft.TagId = tagID
			ft.UserName = curUser.Username
			ft.CreateBy = curUser.Username
			ft.CreateDate = time.Now()
			bll.Add(ft)
		}

	} else {
		ret = "FollowTagError"
		common.WriteLog("not logon, cannot favorite tag")
	}

	return []byte(ret)
}

func removeFollowTag(r *http.Request, favtagID int) []byte {
	bll := DBBLL.NewFavoriteTagBLL(common.Config.DB, common.Config.DBConn)
	_, ok := common.CurrentUser(r)
	ret := "RemoveFavTagSuccess"
	if ok {
		exist := bll.GetById(favtagID)
		if exist == nil {
			ret = "NotFollowTag"
		} else {
			bll.DeleteModel(exist)
		}

	} else {
		ret = "RemoveFavTagError"
		common.WriteLog("not logon, cannot favorite tag")
	}

	return []byte(ret)
}

func getQuestionPageByTag(w http.ResponseWriter, r *http.Request, pIndex, pSize int, keyOne, keyTwo, tagName, sort, lang string) {
	qBll := DBBLL.NewQuestionBLL(common.Config.DB, common.Config.DBConn)
	uBll := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn)
	tBll := DBBLL.NewTagBLL(common.Config.DB, common.Config.DBConn)
	ftBll := DBBLL.NewFavoriteTagBLL(common.Config.DB, common.Config.DBConn)
	curUser, ok := common.CurrentUser(r)
	userName := ""
	if ok {
		userName = curUser.Username
	}

	list, count := qBll.GetPagingListByTag(pIndex, pSize, keyOne, keyTwo, tagName, sort, lang, userName)
	tag := tBll.GetByName(tagName)
	var vm viewmodel.PagingListForTQ
	var vmList []viewmodel.QuestionVM
	for _, q := range list {
		var vm viewmodel.QuestionVM
		vm.Question = q
		vm.PhotoImgSrc = uBll.GetByName(q.UserName).PhotoImgSrc()
		vmList = append(vmList, vm)
	}

	vm.ResultList = vmList
	vm.TotalRows = count
	vm.TagID = tag.Id
	vm.TagTitle = tag.TagName
	vm.IsMyFavorite = ftBll.GetFavoriteTag(userName, tag.Id) != nil

	js, err := json.Marshal(vm)
	if err == nil {
		w.Write(js)
	} else {
		ret := "SystemError"
		w.Write([]byte(ret))
	}
}
