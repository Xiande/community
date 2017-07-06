package controler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	//"strings"
	"DBBLL"
	"viewmodel"
	//"DBModel"
	"common"
)

func CommonAjaxHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		r.ParseForm()
		op := r.Form.Get("op")
		//fmt.Println(op)
		switch op {
		case "getfixedcount":
			count := DBBLL.NewQuestionBLL(common.Config.DB, common.Config.DBConn).GetFixedCount()
			w.Write([]byte(strconv.Itoa(count)))
		case "staruser":
			js, err := getStarUser()
			if err != nil {
				panic(err)
			} else {
				w.Write(js)
			}
		}
	}
}

func getStarUser() (js []byte, err error) {
	var vm viewmodel.StarUser
	file, err := os.Open(fmt.Sprintf("config%cstaruser.json", os.PathSeparator))
	if err != nil {
		err = errors.New("Star User文件读取失败:" + err.Error())
		return
	}

	defer file.Close()

	dec := json.NewDecoder(file)
	err = dec.Decode(&vm)
	if err != nil {
		err = errors.New("Star User文件解析失败:" + err.Error())
		return
	}

	if vm.IsMember && vm.PhotoUrl == "" {
		user := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn).GetByName(vm.UserName)
		if user != nil {
			vm.PhotoUrl = user.PhotoImgSrc()
		}
	}

	js, err = json.Marshal(vm)
	return
}
