/*
分页
*/

package controler

import (
	//	"errors"
	"fmt"
	"html/template"
	"math"
	"strings"
)

// 分页结构体
type Pagination struct {
	//query   *mgo.Query
	count   int
	prePage int
	url     string
}

// 在页面显示分页信息, 内容为 上一页 当前页/下一页 下一页
func (p *Pagination) Html(number int) template.HTML {
	pageCount := int(math.Ceil(float64(p.count) / float64(p.prePage)))

	if pageCount <= 1 {
		return template.HTML("")
	}

	linkFlag := "?"

	if strings.Index(p.url, "?") > -1 {
		linkFlag = "&"
	}

	html := `<ul class="pager">`
	if number > 1 {
		html += fmt.Sprintf(`<li class="previous"><a href="%s%sp=%d">&larr; 上一页</a></li>`, p.url, linkFlag, number-1)
	}

	html += fmt.Sprintf(`<li class="number">%d/%d</li>`, number, pageCount)

	if number < pageCount {
		html += fmt.Sprintf(`<li class="next"><a href="%s%sp=%d">下一页 &rarr;</a></li>`, p.url, linkFlag, number+1)
	}

	return template.HTML(html)
}

// 返回第几页的查询
func (p *Pagination) Page(number int) error {

	return nil
}

// 内容总数
func (p *Pagination) Count() int {
	return p.count
}

// 创建一个分页结构体
func NewPagination(url string, prePage int) *Pagination {
	p := Pagination{}

	p.count, _ = 0, 0
	p.prePage = prePage
	p.url = url

	return &p
}
