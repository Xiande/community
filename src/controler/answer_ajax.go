package controler

import (
	"DBBLL"
	"DBModel"
	"bytes"
	"common"
	"encoding/json"
	//	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"viewmodel"
)

func AnswerAjaxHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == http.MethodGet {
		op := r.Form.Get("op")

		//fmt.Println(op)
		switch op {
		case "bindanswers":
			qid_p := r.Form.Get("qid")
			aid_p := r.Form.Get("aid")
			pageIndex_p := r.Form.Get("pageIndex")
			pageSize_p := r.Form.Get("pageSize")
			sortField_p := r.Form.Get("sortField")
			sortOrder_p := r.Form.Get("sortOrder")
			qid, err := strconv.Atoi(qid_p)
			if err != nil {
				panic(err)
			}

			aid, err := strconv.Atoi(aid_p)
			if err != nil {
				aid = 0
			}

			pageIndex, err := strconv.Atoi(pageIndex_p)
			if err != nil {
				panic(err)
			}

			pageSize, err := strconv.Atoi(pageSize_p)
			if err != nil {
				panic(err)
			}

			bindAnswers(w, r, qid, aid, pageIndex, pageSize, sortField_p, sortOrder_p)
		case "favoriteanswer":
			qid_p := r.Form.Get("qid")
			aid_p := r.Form.Get("aid")
			qid, err := strconv.Atoi(qid_p)
			if err != nil {
				panic(err)
			}

			aid, err := strconv.Atoi(aid_p)
			if err != nil {
				panic(err)
			}

			favoriteAnswer(w, r, qid, aid)
		case "removefa":
			faid_p := r.Form.Get("faid")
			faid, err := strconv.Atoi(faid_p)
			if err != nil {
				panic(err)
			}

			removeFavoriteAnswer(w, r, faid)
		case "voteanswer":
			qid_p := r.Form.Get("qid")
			aid_p := r.Form.Get("aid")
			qid, err := strconv.Atoi(qid_p)
			if err != nil {
				panic(err)
			}

			aid, err := strconv.Atoi(aid_p)
			if err != nil {
				panic(err)
			}

			voteAnswer(w, r, qid, aid)
		case "expertanswer":
			qid_p := r.Form.Get("qid")
			aid_p := r.Form.Get("aid")
			qid, err := strconv.Atoi(qid_p)
			if err != nil {
				panic(err)
			}

			aid, err := strconv.Atoi(aid_p)
			if err != nil {
				panic(err)
			}

			expertAnswer(w, r, qid, aid)
		case "bestanswer":
			qid_p := r.Form.Get("qid")
			aid_p := r.Form.Get("aid")
			qid, err := strconv.Atoi(qid_p)
			if err != nil {
				panic(err)
			}

			aid, err := strconv.Atoi(aid_p)
			if err != nil {
				panic(err)
			}

			bestAnswer(w, r, qid, aid)
		case "checkstatus":
			qid_p := r.Form.Get("qid")
			status := r.Form.Get("status")
			qid, err := strconv.Atoi(qid_p)
			if err != nil {
				panic(err)
			}

			if err != nil {
				panic(err)
			}

			checkAnswerStatus(w, r, qid, status)
		case "getvoteques":
			pageIndex_p := r.Form.Get("pIndex")
			pageSize_p := r.Form.Get("pSize")

			pageIndex, err := strconv.Atoi(pageIndex_p)
			if err != nil {
				panic(err)
			}

			pageSize, err := strconv.Atoi(pageSize_p)
			if err != nil {
				panic(err)
			}

			getQuestionByVotedAnswer(w, r, pageIndex, pageSize)
		case "getfavanswer":
			pageIndex_p := r.Form.Get("pIndex")
			pageSize_p := r.Form.Get("pSize")

			pageIndex, err := strconv.Atoi(pageIndex_p)
			if err != nil {
				panic(err)
			}

			pageSize, err := strconv.Atoi(pageSize_p)
			if err != nil {
				panic(err)
			}

			getQuestionByFavoriteAnswer(w, r, pageIndex, pageSize)
		}
	}

	if r.Method == http.MethodPost {
		bs, _ := ioutil.ReadAll(r.Body)
		js := bytes.NewBuffer(bs).String()
		js, _ = url.QueryUnescape(js)
		//fmt.Println(js)
		var postData ClientPostAnswer
		err := json.Unmarshal([]byte(js), &postData)
		if err != nil {
			common.WriteLog(err.Error())
			panic(err)
		}

		user, ok := common.CurrentUser(r)

		if ok {
			var a DBModel.Answer
			qid, _ := strconv.Atoi(postData.QuestionID)
			a.QuestionId = qid
			a.CreateBy = user.Username
			a.UserName = user.Username
			a.DisplayName = user.DisplayName
			a.CreateDate = time.Now()
			aBll := DBBLL.NewAnswerBLL(common.Config.DB, common.Config.DBConn)
			aBll.AddTx(a, postData.Content)
			qbll := DBBLL.NewQuestionBLL(common.Config.DB, common.Config.DBConn)
			q := qbll.GetById(qid)
			q.AnswersCount++
			qbll.Update(*q)
			//if q.UserName != a.UserName && len(aBll.GetListByUserName(user.Username)) == 1 {

			//}

			ret := "Success"
			w.Write([]byte(ret))
		} else {
			ret := "SystemError"
			w.Write([]byte(ret))
		}

		return

	}

	defer func() {
		if x := recover(); x != nil {
			common.WriteLog(x.(error).Error())
		}
	}()
}

func bindAnswers(w http.ResponseWriter, r *http.Request, qid, aid, pageIndex, pageSize int, sortField, sortOrder string) {
	qBll := DBBLL.NewQuestionBLL(common.Config.DB, common.Config.DBConn)
	aBll := DBBLL.NewAnswerBLL(common.Config.DB, common.Config.DBConn)
	userBll := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn)
	acBll := DBBLL.NewAnswerContentBLL(common.Config.DB, common.Config.DBConn)
	fuBll := DBBLL.NewFollowUserBLL(common.Config.DB, common.Config.DBConn)

	q := qBll.GetById(qid)
	curUser, ok := common.CurrentUser(r)
	aList, count := aBll.GetPagingListByQuestionId(qid, aid, pageIndex, pageSize, sortField, sortOrder)
	var vmList []viewmodel.AnswerVM
	for _, a := range aList {
		var vm viewmodel.AnswerVM
		//fmt.Print(a.Id)
		vm.Id = a.Id
		vm.BestTime = a.BestTime
		vm.CreateBy = a.CreateBy
		vm.CreateDate = a.CreateDate
		vm.DisplayName = a.DisplayName
		vm.ExpertTime = a.ExpertTime
		vm.IsBest = a.IsBest
		vm.IsExpert = a.IsExpert
		vm.QuestionId = a.QuestionId
		vm.UserName = a.UserName
		vm.VotedCount = a.VotedCount
		//TBD:

		if ok {
			vm.IsFollowed = fuBll.GetRecord(curUser.Username, a.UserName) != nil
			vm.IsAuthor = curUser.Username == q.UserName
			vm.IsAdmin = curUser.IsSuperuser
			vm.IsModerator = curUser.IsModerator
		} else {
			vm.IsAuthor = false
			vm.IsAdmin = false
			vm.IsModerator = false
		}

		user := userBll.GetByName(a.UserName)
		vm.AuthorEmail = user.Email
		vm.PhotoImgSrc = user.PhotoImgSrc()
		vm.AnswerContent = acBll.GetByAnswerId(a.Id).AnswerContent
		vm.CanBest = aBll.GetBestByQuestionId(qid) == nil
		vm.CanExpert = aBll.GetExpertByQuestionId(qid) == nil

		vmList = append(vmList, vm)
	}

	var ret viewmodel.PagingList
	ret.ResultList = vmList
	ret.TotalRows = count

	js, err := json.Marshal(ret)
	if err == nil {
		//fmt.Println(string(js))
		w.Write(js)
	} else {
		ret := "SystemError"
		w.Write([]byte(ret))
	}
}
func getQuestionByVotedAnswer(w http.ResponseWriter, r *http.Request, pageIndex, pageSize int) {
	qBll := DBBLL.NewQuestionBLL(common.Config.DB, common.Config.DBConn)
	userBll := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn)
	fuBll := DBBLL.NewFollowUserBLL(common.Config.DB, common.Config.DBConn)
	curUser, ok := common.CurrentUser(r)
	if ok {
		qList, count := qBll.GetQuestionByVotedAnswer(pageIndex, pageSize, curUser.Username)
		var vmList []viewmodel.FavoriteQuestionVM
		for _, q := range qList {
			var vm viewmodel.FavoriteQuestionVM
			//fmt.Print(a.Id)
			vm.Favoritequestionext = q
			vm.PhotoImgSrc = userBll.GetByName(q.UserName).PhotoImgSrc()
			vm.IsFollowed = fuBll.GetRecord(curUser.Username, q.UserName) != nil
			vmList = append(vmList, vm)
		}

		var ret viewmodel.PagingList
		ret.ResultList = vmList
		ret.TotalRows = count

		js, err := json.Marshal(ret)
		if err == nil {
			//fmt.Println(string(js))
			w.Write(js)
		} else {
			ret := "SystemError"
			w.Write([]byte(ret))
		}
	}
}
func getQuestionByFavoriteAnswer(w http.ResponseWriter, r *http.Request, pageIndex, pageSize int) {
	fuBll := DBBLL.NewFollowUserBLL(common.Config.DB, common.Config.DBConn)
	qBll := DBBLL.NewQuestionBLL(common.Config.DB, common.Config.DBConn)
	userBll := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn)
	curUser, ok := common.CurrentUser(r)
	if ok {
		qList, count := qBll.GetQuestionByFavoriteAnswer(pageIndex, pageSize, curUser.Username)
		var vmList []viewmodel.FavoriteQuestionVM
		for _, q := range qList {
			var vm viewmodel.FavoriteQuestionVM
			//fmt.Print(a.Id)
			vm.Favoritequestionext = q
			vm.PhotoImgSrc = userBll.GetByName(q.UserName).PhotoImgSrc()
			vm.IsFollowed = fuBll.GetRecord(curUser.Username, q.UserName) != nil
			vmList = append(vmList, vm)
		}

		var ret viewmodel.PagingList
		ret.ResultList = vmList
		ret.TotalRows = count

		js, err := json.Marshal(ret)
		if err == nil {
			//fmt.Println(string(js))
			w.Write(js)
		} else {
			ret := "SystemError"
			w.Write([]byte(ret))
		}
	}
}

func favoriteAnswer(w http.ResponseWriter, r *http.Request, qid, aid int) {
	faBll := DBBLL.NewFavoriteAnswerBLL(common.Config.DB, common.Config.DBConn)
	ret := "Success"
	curUser, ok := common.CurrentUser(r)
	if ok {
		fa := faBll.GetFavoriteAnswer(curUser.Username, aid)

		if fa != nil {
			ret = "HadFavoriteAnswer"
		} else {
			var model DBModel.Favoriteanswer
			model.QuestionId = qid
			model.AnswerId = aid
			model.CreateBy = curUser.Username
			model.UserName = curUser.Username
			model.CreateDate = time.Now()
			faBll.Add(model)
		}
	} else {
		ret = "SystemError"
	}

	w.Write([]byte(ret))
}
func removeFavoriteAnswer(w http.ResponseWriter, r *http.Request, faid int) {
	faBll := DBBLL.NewFavoriteAnswerBLL(common.Config.DB, common.Config.DBConn)
	ret := "Success"
	faBll.Delete(faid)
	w.Write([]byte(ret))
}
func voteAnswer(w http.ResponseWriter, r *http.Request, qid, aid int) {
	avrBll := DBBLL.NewAnswerVotedRecordBLL(common.Config.DB, common.Config.DBConn)
	ret := "Success"
	curUser, ok := common.CurrentUser(r)
	if ok {
		avr := avrBll.GetAnswerVotedRcord(curUser.Username, aid)
		if avr == nil {
			var model DBModel.Answervotedrecord
			model.QuestionId = qid
			model.AnswerId = aid
			model.UserName = curUser.Username
			model.CreateBy = curUser.Username
			model.CreateDate = time.Now()
			avrBll.Add(model)
		} else {
			ret = "HadVoted"
		}
	} else {
		ret = "SystemError"
	}

	w.Write([]byte(ret))
}

func expertAnswer(w http.ResponseWriter, r *http.Request, qid, aid int) {
	aBll := DBBLL.NewAnswerBLL(common.Config.DB, common.Config.DBConn)
	ret := "Success"
	_, ok := common.CurrentUser(r)
	if ok {
		aBll.ExpertAnswer(qid, aid)
	} else {
		ret = "SystemError"
	}

	w.Write([]byte(ret))
}
func bestAnswer(w http.ResponseWriter, r *http.Request, qid, aid int) {
	aBll := DBBLL.NewAnswerBLL(common.Config.DB, common.Config.DBConn)
	ret := "Success"
	_, ok := common.CurrentUser(r)
	if ok {
		aBll.BestAnswer(qid, aid)
	} else {
		ret = "SystemError"
	}

	w.Write([]byte(ret))
}
func checkAnswerStatus(w http.ResponseWriter, r *http.Request, qid int, status string) {
	ret := "False"
	aBll := DBBLL.NewAnswerBLL(common.Config.DB, common.Config.DBConn)
	if status == "best" {
		best := aBll.GetBestByQuestionId(qid)
		if best == nil {
			ret = "True"
		}
	} else if status == "expert" {
		expert := aBll.GetExpertByQuestionId(qid)
		if expert == nil {
			ret = "True"
		}
	} else {
		ret = "SystemError"
	}

	w.Write([]byte(ret))
}

type ClientPostAnswer struct {
	QuestionID string `json:"QuestionID"`
	Content    string `json:"Content"`
}
