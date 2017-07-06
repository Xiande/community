package controler

import (
	"common"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//topicsHandler(w, r)
	common.RenderTemplate(w, r, "index.html", map[string]interface{}{})
}
