/*
读取配置文件,设置URL,启动服务器
*/

package community

import (
	"common"
	"controler"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerFun(handler controler.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if handler.Permission == common.Everyone {
			handler.HandlerFunc(w, r)
		} else if handler.Permission == common.Authenticated {
			_, ok := common.CurrentUser(r)

			if !ok {
				http.Redirect(w, r, "/signin?source="+r.RequestURI, http.StatusFound)
				return
			}

			handler.HandlerFunc(w, r)
		} else if handler.Permission == common.Administrator {
			user, ok := common.CurrentUser(r)

			if !ok {
				http.Redirect(w, r, "/signin?source="+r.RequestURI, http.StatusFound)

				return
			}

			if !user.IsSuperuser {
				common.Message(w, r, "没有权限", "对不起，你没有权限进行该操作", "error")
				return
			}

			handler.HandlerFunc(w, r)
		}
	}
}

func StartServer() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.Handle("/uploadimages/", http.FileServer(http.Dir(".")))
	r := mux.NewRouter()
	for _, handler := range controler.Handlers {
		r.HandleFunc(handler.URL, handlerFun(handler))
	}

	http.Handle("/", r)

	fmt.Println("Server start on:", common.Config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", common.Config.Port), nil))
}
