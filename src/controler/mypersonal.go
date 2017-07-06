package controler

import (
	"common"
	"net/http"
)

func MyPersonalHandler(w http.ResponseWriter, r *http.Request) {
	common.RenderTemplate(w, r, "mypersonal.html", map[string]interface{}{})
}
