package controler

import (
	"DBBLL"
	"DBModel"
	"common"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	//	"io"
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strings"
	"viewmodel"

	"code.google.com/p/go-uuid/uuid"
)

func QuestionAjaxHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		r.ParseForm()
		op := r.Form.Get("op")
		//fmt.Println(op)
		switch op {
		case "viewquestion":
			qid := r.Form.Get("qid")
			id, err := strconv.Atoi(qid)
			if err == nil {
				viewQuestion(w, r, id)
			}
		case "getques":
			pageIndex := r.Form.Get("pIndex")
			pageSize := r.Form.Get("pSize")
			keyOne := r.Form.Get("keyOne")
			keyTwo := r.Form.Get("keyTwo")
			sType := r.Form.Get("sType")
			sort := r.Form.Get("sort")
			lang := r.Form.Get("lang")
			pIndex, _ := strconv.Atoi(pageIndex)
			pSize, _ := strconv.Atoi(pageSize)

			if lang == "" {
				lang = "all"
			}
			getQuestionPage(w, r, pIndex, pSize, keyOne, keyTwo, sType, sort, lang)
		case "favoritequestion":
			id := r.Form.Get("qid")
			qid, _ := strconv.Atoi(id)
			favoriteQuestion(w, r, qid)
		case "removefq":
			id := r.Form.Get("fqid")
			fqid, _ := strconv.Atoi(id)
			removeFavoriteQuestion(w, r, fqid)
		case "getfavques":
			pageIndex := r.Form.Get("pIndex")
			pageSize := r.Form.Get("pSize")
			pIndex, _ := strconv.Atoi(pageIndex)
			pSize, _ := strconv.Atoi(pageSize)
			getFavoriteQuestionPage(w, r, pIndex, pSize)
		case "deletequestion":
			id := r.Form.Get("qid")
			qid, _ := strconv.Atoi(id)
			deleteQuestion(w, r, qid)
		case "topdynamic":
			top_p := r.Form.Get("top")
			top, err := strconv.Atoi(top_p)
			if err != nil {
				top = 5
			}
			getTopDynamic(w, r, top)
		}
	}

	if r.Method == http.MethodPost {
		r.ParseForm()

		bs, _ := ioutil.ReadAll(r.Body)
		js := bytes.NewBuffer(bs).String()
		js, _ = url.QueryUnescape(js)
		var postData ClientPostQuestion
		err := json.Unmarshal([]byte(js), &postData)
		if err != nil {
			panic(err)
		}

		user, ok := common.CurrentUser(r)

		if ok {
			var q DBModel.Question
			q.Title = postData.Title
			q.LanguageType = postData.LanguageType
			q.UserName = user.Username
			q.DisplayName = user.DisplayName

			q.CreateBy = user.Username
			q.CreateDate = time.Now()
			q.Tags = strings.TrimRight(postData.Tags, ",")
			q.EmailNotice = postData.NoticeEmail

			qbll := DBBLL.NewQuestionBLL(common.Config.DB, common.Config.DBConn)
			qbll.AddTx(q, postData.Content)

			ret := "AddQuestionSuccess"
			w.Write([]byte(ret))
		} else {
			ret := "SystemError"
			w.Write([]byte(ret))
		}
		return
	}
}

func ImageUploadAjaxHandler(w http.ResponseWriter, r *http.Request) {
	exts := "|.png|.jpg|.bmp|"
	dir := "uploadimages"

	if r.Method == "POST" {
		r.ParseForm()
		num := r.Form.Get("CKEditorFuncNum")

		_, err := os.Stat(dir)
		if err != nil {
			os.Mkdir(dir, os.ModePerm)
		}

		id := uuid.NewUUID().String()
		file, h, err := r.FormFile("upload")

		if err != nil {
			//Paste Image
			bs, _ := ioutil.ReadAll(r.Body)
			js := bytes.NewBuffer(bs).String()
			js, _ = url.QueryUnescape(js)
			var postData PasteImage
			err := json.Unmarshal([]byte(js), &postData)
			if err != nil {
				panic(err)
			}

			base64Str, _ := base64.StdEncoding.DecodeString(postData.Data)
			filename := fmt.Sprintf("%s%c%s%s", dir, os.PathSeparator, id, postData.Ext)
			err = ioutil.WriteFile(filename, base64Str, os.ModePerm)
			if err != nil {
				common.WriteLog(err.Error())
			}

			url := fmt.Sprintf("/%s/%s%s", dir, id, postData.Ext)
			w.Write([]byte(url))

		} else {
			defer file.Close()
			ext := path.Ext(h.Filename)
			if !strings.Contains(exts, "|"+ext+"|") {
				retStr := "<script type=\"text/javascript\"> window.parent.CKEDITOR.tools.callFunction(" + num + ",''," + "window.parent.dict['FileFormatTip']);</script>"
				w.Write([]byte(retStr))
				return
			}

			var size int64
			if statInterface, ok := file.(common.Stat); ok {
				fileInfo, _ := statInterface.Stat()
				size = fileInfo.Size()
			}

			if sizeInterface, ok := file.(common.Sizer); ok {
				size = sizeInterface.Size()
			}

			if size > 1024*600 {
				retStr := "<script type=\"text/javascript\"> window.parent.CKEDITOR.tools.callFunction(" + num + ",''," + "window.parent.dict['FileSizeTip']);</script>"
				w.Write([]byte(retStr))
				return
			}

			f, err := os.Create(fmt.Sprintf("%s%c%s%s", dir, os.PathSeparator, id, ext))
			if err != nil {
				common.WriteLog(err.Error())
				panic(err)
			}

			defer f.Close()
			common.ZoomAuto(file, f, 780, ext)
			url := fmt.Sprintf("/%s/%s%s", dir, id, ext)
			retStr := "<script type=\"text/javascript\"> window.parent.CKEDITOR.tools.callFunction(" + num + ",'" + url + "');</script>"
			//fmt.Println(retStr)
			w.Write([]byte(retStr))
		}

	}

}

func viewQuestion(w http.ResponseWriter, r *http.Request, qid int) {
	qBll := DBBLL.NewQuestionBLL(common.Config.DB, common.Config.DBConn)
	qcBll := DBBLL.NewQuestionContentBLL(common.Config.DB, common.Config.DBConn)
	userBll := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn)
	aBll := DBBLL.NewAnswerBLL(common.Config.DB, common.Config.DBConn)
	acBll := DBBLL.NewAnswerContentBLL(common.Config.DB, common.Config.DBConn)
	fuBll := DBBLL.NewFollowUserBLL(common.Config.DB, common.Config.DBConn)
	q := qBll.GetById(qid)
	qc := qcBll.GetByQuestionId(qid)
	author := userBll.GetByName(q.UserName)
	curUser, ok := common.CurrentUser(r)
	aList := aBll.GetListByQuestionId(qid)
	var vm viewmodel.QuestionVM
	vm.Id = q.Id
	vm.QuestionContent = qc.QuestionContent
	vm.UserName = q.UserName
	if ok {
		vm.IsAuthor = q.UserName == curUser.Username
		vm.IsAdmin = curUser.IsSuperuser
		vm.IsModerator = curUser.IsModerator
		vm.IsFollowed = fuBll.GetRecord(curUser.Username, q.UserName) != nil
	}

	vm.DisplayName = q.DisplayName
	vm.Title = q.Title
	vm.Tags = q.Tags
	vm.CreateBy = q.CreateBy
	vm.CreateDate = q.CreateDate
	vm.EmailNotice = q.EmailNotice
	vm.ViewedCount = q.ViewedCount
	vm.AuthorEmail = author.Email
	vm.PhotoImgSrc = userBll.GetByName(q.UserName).PhotoImgSrc()
	//fmt.Println(vm.PhotoImgSrc)
	vm.AnswersCount = len(aList)
	ba := aBll.GetBestByQuestionId(qid)
	if ba != nil {
		var avm viewmodel.AnswerVM
		avm.Answer = *ba

		if ok {
			avm.IsAuthor = curUser.Username == q.UserName
			avm.IsAdmin = curUser.IsSuperuser
			avm.IsModerator = curUser.IsModerator
			avm.IsFollowed = fuBll.GetRecord(curUser.Username, ba.UserName) != nil
		} else {
			avm.IsAuthor = false
			avm.IsAdmin = false
			avm.IsModerator = false
		}

		user := userBll.GetByName(ba.UserName)
		avm.AuthorEmail = user.Email
		avm.PhotoImgSrc = user.PhotoImgSrc()
		avm.AnswerContent = acBll.GetByAnswerId(ba.Id).AnswerContent
		avm.CanBest = aBll.GetBestByQuestionId(qid) == nil
		avm.CanExpert = aBll.GetExpertByQuestionId(qid) == nil

		vm.BestAnswer = avm
	}

	ea := aBll.GetExpertByQuestionId(qid)
	if ea != nil {
		var avm viewmodel.AnswerVM
		avm.Answer = *ea

		if ok {
			avm.IsAuthor = curUser.Username == q.UserName
			avm.IsAdmin = curUser.IsSuperuser
			avm.IsModerator = curUser.IsModerator
			avm.IsFollowed = fuBll.GetRecord(curUser.Username, ea.UserName) != nil
		} else {
			avm.IsAuthor = false
			avm.IsAdmin = false
			avm.IsModerator = false
		}

		user := userBll.GetByName(ea.UserName)
		avm.AuthorEmail = user.Email
		avm.PhotoImgSrc = user.PhotoImgSrc()
		avm.AnswerContent = acBll.GetByAnswerId(ea.Id).AnswerContent
		avm.CanBest = aBll.GetBestByQuestionId(qid) == nil
		avm.CanExpert = aBll.GetExpertByQuestionId(qid) == nil

		vm.ExpertAnswer = avm
	}

	js, err := json.Marshal(vm)
	if err == nil {
		w.Write(js)
	} else {
		ret := "SystemError"
		w.Write([]byte(ret))
	}
}

func favoriteQuestion(w http.ResponseWriter, r *http.Request, qid int) {
	fqBll := DBBLL.NewFavoriteQuestionBLL(common.Config.DB, common.Config.DBConn)

	curUser, ok := common.CurrentUser(r)
	if ok {
		fq := fqBll.GetFavoriteQuestion(curUser.Username, qid)
		if fq == nil {
			var model DBModel.Favoritequestion
			model.QuestionId = qid
			model.UserName = curUser.Username
			model.CreateBy = curUser.Username
			model.CreateDate = time.Now()
			fqBll.Add(model)
			w.Write([]byte("Success"))
		} else {
			w.Write([]byte("HadFavoriteQuestion"))
		}
	}

	return
}

func removeFavoriteQuestion(w http.ResponseWriter, r *http.Request, fqid int) {
	fqBll := DBBLL.NewFavoriteQuestionBLL(common.Config.DB, common.Config.DBConn)

	_, ok := common.CurrentUser(r)
	if ok {
		fq := fqBll.GetById(fqid)
		if fq != nil {
			fqBll.DeleteModel(fq)
			w.Write([]byte("Success"))
		} else {
			w.Write([]byte("SystemError"))
		}
	}

	return
}

func getQuestionPage(w http.ResponseWriter, r *http.Request, pIndex, pSize int, keyOne, keyTwo, sType, sort, lang string) {
	qBll := DBBLL.NewQuestionBLL(common.Config.DB, common.Config.DBConn)
	uBll := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn)
	fuBll := DBBLL.NewFollowUserBLL(common.Config.DB, common.Config.DBConn)
	curUser, ok := common.CurrentUser(r)
	userName := ""
	if ok {
		userName = curUser.Username
	}

	list, count := qBll.GetPagingList(pIndex, pSize, keyOne, keyTwo, sType, sort, lang, userName)
	var vmList []viewmodel.QuestionVM
	for _, q := range list {
		var vm viewmodel.QuestionVM
		vm.Question = q
		vm.PhotoImgSrc = uBll.GetByName(q.UserName).PhotoImgSrc()
		vm.IsFollowed = fuBll.GetRecord(userName, q.UserName) != nil
		vmList = append(vmList, vm)
	}
	var ret viewmodel.PagingList
	ret.ResultList = vmList
	ret.TotalRows = count
	js, err := json.Marshal(ret)
	if err == nil {
		w.Write(js)
	} else {
		ret := "SystemError"
		w.Write([]byte(ret))
	}
}

func getFavoriteQuestionPage(w http.ResponseWriter, r *http.Request, pIndex, pSize int) {
	fqBll := DBBLL.NewFavoriteQuestionBLL(common.Config.DB, common.Config.DBConn)
	fuBll := DBBLL.NewFollowUserBLL(common.Config.DB, common.Config.DBConn)
	uBll := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn)
	curUser, ok := common.CurrentUser(r)
	if ok {
		ms, count := fqBll.GetListByUesrName(pIndex, pSize, curUser.Username)
		var vmList []viewmodel.FavoriteQuestionVM
		for _, q := range ms {
			var vm viewmodel.FavoriteQuestionVM
			vm.Favoritequestionext = q
			vm.IsFollowed = fuBll.GetRecord(curUser.Username, q.UserName) != nil
			vm.PhotoImgSrc = uBll.GetByName(q.UserName).PhotoImgSrc()
			vmList = append(vmList, vm)
		}

		var ret viewmodel.PagingList
		ret.ResultList = vmList
		ret.TotalRows = count
		js, err := json.Marshal(ret)
		if err == nil {
			w.Write(js)
		} else {
			ret := ""
			w.Write([]byte(ret))
		}
	}
}
func deleteQuestion(w http.ResponseWriter, r *http.Request, qid int) {
	qBll := DBBLL.NewQuestionBLL(common.Config.DB, common.Config.DBConn)
	qBll.Delete(qid)
	ret := "Success"
	w.Write([]byte(ret))
}

func getTopDynamic(w http.ResponseWriter, r *http.Request, top int) {
	qBll := DBBLL.NewQuestionBLL(common.Config.DB, common.Config.DBConn)
	//uBll := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn)
	curUser, ok := common.CurrentUser(r)
	if ok {
		var vms []viewmodel.UserQuestionTop
		qs := qBll.GetUserQuestionTop(curUser.Username, top)
		for _, q := range qs {
			var vm viewmodel.UserQuestionTop
			vm.Id = q.Id
			vm.UserName = q.UserName
			vm.DisplayName = q.DisplayName
			vm.Title = q.Title
			vm.Photo = DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn).GetByName(q.UserName).PhotoImgSrc()
			vm.CreateDate = q.CreateDate

			vms = append(vms, vm)
		}

		js, err := json.Marshal(vms)
		if err == nil {
			w.Write(js)
		} else {
			w.Write([]byte(js))
		}
	}
}

type PasteImage struct {
	Data string `json:"data"`
	Op   string `json:"op"`
	Ext  string `json:ext`
}

type ClientPostQuestion struct {
	Title        string `json:"Title"`
	Content      string `json:"Content"`
	Tags         string `json:"Tags"`
	NoticeEmail  bool   `json:"NoticeEmail"`
	LanguageType string `json:"LanguageType"`
}
