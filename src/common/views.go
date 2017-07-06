/*
一些辅助方法
*/

package common

import (
	"DBBLL"
	"DBModel"
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	//"github.com/gorilla/mux"

	"github.com/gorilla/sessions"
	"github.com/jimmykuu/webhelpers"
	"github.com/jimmykuu/wtforms"
	"github.com/nfnt/resize"
)

const (
	PerPage       = 20
	TypeTopic     = 'T'
	TypeArticle   = 'A'
	TypeSite      = 'S'
	TypePackage   = 'P'
	DefaultAvatar = "gopher_teal.jpg"
)

var (
	Store       *sessions.CookieStore
	fileVersion map[string]string = make(map[string]string) // {path: version}
	utils       *Utils
)

var funcMaps = template.FuncMap{
	"gravatar": func(email string, size uint16) string {
		h := md5.New()
		io.WriteString(h, email)
		return fmt.Sprintf("http://www.gravatar.com/avatar/%x?s=%d", h.Sum(nil), size)
	},
}

type Utils struct {
}

// 没有http://开头的增加http://
func (u *Utils) Url(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	}

	return "http://" + url
}

func (u *Utils) StaticUrl(path string) string {
	version, ok := fileVersion[path]
	if ok {
		return "/static/" + path + "?v=" + version
	}

	file, err := os.Open("static/" + path)

	if err != nil {
		return "/static/" + path
	}

	h := md5.New()

	_, err = io.Copy(h, file)

	version = fmt.Sprintf("%x", h.Sum(nil))[:5]

	fileVersion[path] = version

	return "/static/" + path + "?v=" + version
}

func (u *Utils) Index(index int) int {
	return index + 1
}

func (u *Utils) Equal(src, dest string) bool {
	return src == dest
}

func (u *Utils) FormatTime(t time.Time) string {
	now := time.Now()
	duration := now.Sub(t)
	if duration.Seconds() < 60 {
		return fmt.Sprintf("刚刚")
	} else if duration.Minutes() < 60 {
		return fmt.Sprintf("%.0f 分钟前", duration.Minutes())
	} else if duration.Hours() < 24 {
		return fmt.Sprintf("%.0f 小时前", duration.Hours())
	}

	t = t.Add(time.Hour * time.Duration(Config.TimeZoneOffset))
	return t.Format("2006-01-02 15:04")
}

func (u *Utils) UserInfo(username string) template.HTML {
	bll := DBBLL.NewUserBLL(Config.DB, Config.DBConn)
	// 检查用户名
	user := bll.GetByName(username)

	format := `<div>
        <a href="/member/%s"><img class="gravatar img-rounded" src="%s-middle" style="float:left;"></a>
        <h3><a href="/member/%s">%s</a></h3>
        <div class="clearfix"></div>
    </div>`

	return template.HTML(fmt.Sprintf(format, username, user.PhotoImgSrc(), username, username))
}

func (u *Utils) Truncate(html template.HTML, length int) string {
	text := webhelpers.RemoveFormatting(string(html))
	return webhelpers.Truncate(text, length, "...")
}

func (u *Utils) HTML(str string) template.HTML {
	return template.HTML(str)
}

func (u *Utils) UserPhotoImgSrc(userName string) string {
	bll := DBBLL.NewUserBLL(Config.DB, Config.DBConn)
	user := bll.GetByName(userName)
	return user.PhotoImgSrc()
}

func (u *Utils) RenderInput(form wtforms.Form, fieldStr string, inputAttrs ...string) template.HTML {
	field, err := form.Field(fieldStr)
	if err != nil {
		panic(err)
	}

	errorClass := ""

	if field.HasErrors() {
		errorClass = " has-error"
	}

	format := `<div class="form-group%s">
        %s
        %s
        %s
    </div>`

	var inputAttrs2 []string = []string{`class="form-control"`}
	inputAttrs2 = append(inputAttrs2, inputAttrs...)

	return template.HTML(
		fmt.Sprintf(format,
			errorClass,
			field.RenderLabel(),
			field.RenderInput(inputAttrs2...),
			field.RenderErrors()))
}

func (u *Utils) RenderInputH(form wtforms.Form, fieldStr string, labelWidth, inputWidth int, inputAttrs ...string) template.HTML {
	field, err := form.Field(fieldStr)
	if err != nil {
		panic(err)
	}

	errorClass := ""

	if field.HasErrors() {
		errorClass = " has-error"
	}
	format := `<div class="form-group%s">
        %s
        <div class="col-lg-%d">
            %s%s
        </div>
    </div>`
	labelClass := fmt.Sprintf(`class="col-lg-%d control-label"`, labelWidth)

	var inputAttrs2 []string = []string{`class="form-control"`}
	inputAttrs2 = append(inputAttrs2, inputAttrs...)

	return template.HTML(
		fmt.Sprintf(format,
			errorClass,
			field.RenderLabel(labelClass),
			inputWidth,
			field.RenderInput(inputAttrs2...),
			field.RenderErrors(),
		))
}

// 在模板中渲染成表单控件
// func (u *Utils) Input(form wtforms.Form, fieldName string, attrs ...string) template.HTML {
// 	field, err := form.Field(fieldName)

// 	if err != nil {
// 		panic(err)
// 	}

// 	errClass := ""
// 	if field.HasErrors() {
// 		errClass = " has-error"
// 	}

// 	return template.HTML(fmt.Sprintf(`<div class="form-group%s">%s%s%s</div>`, errClass, field.RenderLabel(), field.RenderInput(attrs), field.RenderErrors()))
// }

// 在模板中渲染成表单控件, 水平排列
// func (u *Utils) InputH(form wtforms.Form, fieldName string, attrs ...string) template.HTML {
// 	field, err := form.Field(fieldName)

// 	if err != nil {
// 		panic(err)
// 	}

// 	errClass := ""
// 	if field.HasErrors() {
// 		errClass = " has-error"
// 	}

// 	return template.HTML(fmt.Sprintf(`<div class="form-group%s">%s%s%s</div>`, errClass, field.RenderLabel(`class="control-label"`), field.RenderInput(attrs), field.RenderErrors()))
// }

func Message(w http.ResponseWriter, r *http.Request, title string, message string, class string) {
	RenderTemplate(w, r, "message.html", map[string]interface{}{"title": title, "message": template.HTML(message), "class": class})
}

// 获取链接的页码，默认"?p=1"这种类型
func Page(r *http.Request) (int, error) {
	p := r.FormValue("p")
	page := 1

	if p != "" {
		var err error
		page, err = strconv.Atoi(p)

		if err != nil {
			return 0, err
		}
	}

	return page, nil
}

func SendMail(subject string, message string, to []string) {
	auth := smtp.PlainAuth(
		"",
		Config.SmtpUsername,
		Config.SmtpPassword,
		Config.SmtpHost,
	)
	msg := fmt.Sprintf("To: %s\r\nFrom: jimmykuu@126.com\r\nSubject: %s\r\nContent-Type: text/html\r\n\r\n%s", strings.Join(to, ";"), subject, message)
	err := smtp.SendMail(Config.SmtpAddr, auth, Config.FromEmail, to, []byte(msg))
	if err != nil {
		panic(err)
	}
}

// 检查一个string元素是否在数组里面
func stringInArray(a []string, x string) bool {
	sort.Strings(a)
	index := sort.SearchStrings(a, x)

	if index == 0 {
		if a[0] == x {
			return true
		}

		return false
	} else if index > len(a)-1 {
		return false
	}

	return true
}

func init() {
	if Config.DB == "" {
		fmt.Println("数据库地址还没有配置,请到config.json内配置db字段.")
		os.Exit(1)
	}

	/*
		session, err := mgo.Dial(Config.DB)
		if err != nil {
			fmt.Println("MongoDB连接失败:", err.Error())
			os.Exit(1)
		}

		session.SetMode(mgo.Monotonic, true)

		DB = session.DB("gopher")
	*/

	Store = sessions.NewCookieStore([]byte(Config.CookieSecret))

	utils = &Utils{}

	// 如果没有status,创建
	/*
		var status Status
		c := DB.C("status")
		err = c.Find(nil).One(&status)

		if err != nil {
			c.Insert(&Status{
				Id_:        bson.NewObjectId(),
				UserCount:  0,
				TopicCount: 0,
				ReplyCount: 0,
				UserIndex:  0,
			})
		}
	*/

	// 检查是否有超级账户设置
	var superusers []string
	for _, username := range strings.Split(Config.Superusers, ",") {
		username = strings.TrimSpace(username)
		if username != "" {
			superusers = append(superusers, username)
		}
	}

	if len(superusers) == 0 {
		fmt.Println("你没有设置超级账户,请在config.json中的superusers中设置,如有多个账户,用逗号分开")
	}

	bll := DBBLL.NewUserBLL(Config.DB, Config.DBConn)

	var users []DBModel.User
	users = bll.GetList("IsSuperuser = 1")

	// 如果mongodb中的超级用户不在配置文件中,取消超级用户
	for _, user := range users {
		if !stringInArray(superusers, user.Username) {
			user.IsSuperuser = false
			bll.Update(user)
		}
	}

	// 设置超级用户
	for _, username := range superusers {
		var dbuser *DBModel.User
		dbuser = bll.GetByName(username)
		if dbuser != nil {
			dbuser.IsSuperuser = true
			bll.Update(*dbuser)
		}
	}
}

func ParseTemplate(file string, data map[string]interface{}) []byte {
	var buf bytes.Buffer

	t, err := template.ParseFiles("templates/base.html", "templates/"+file)
	if err != nil {
		panic(err)
	}
	t = t.Funcs(funcMaps)
	err = t.Execute(&buf, data)

	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

// 返回当前用户
func CurrentUser(r *http.Request) (*DBModel.User, bool) {
	session, _ := Store.Get(r, "user")
	username, ok := session.Values["username"]
	//	fmt.Println(session.Values["username"])
	if !ok {
		return nil, false
	}

	bll := DBBLL.NewUserBLL(Config.DB, Config.DBConn)

	// 检查用户名
	user := bll.GetByName(username.(string))

	if user == nil {
		return nil, false
	}

	return user, true
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, file string, data map[string]interface{}) {
	_, isPresent := data["signout"]

	// 如果isPresent==true，说明在执行登出操作
	if !isPresent {
		// 加入用户信息
		user, ok := CurrentUser(r)

		if ok {
			data["username"] = user.Username
			data["displayname"] = user.DisplayName
			data["isSuperUser"] = user.IsSuperuser
			data["email"] = user.Email
			data["photoImgSrc"] = user.PhotoImgSrc()
			data["fansCount"] = len(user.Fans)
			data["followCount"] = len(user.Follow)
			data["IsModerator"] = user.IsModerator
		}
	}

	data["utils"] = utils

	data["analyticsCode"] = analyticsCode
	data["staticFileVersion"] = Config.StaticFileVersion
	flash, _ := Store.Get(r, "flash")
	data["flash"] = flash
	data["goVersion"] = goVersion

	_, ok := data["active"]
	if !ok {
		data["active"] = ""
	}

	page := ParseTemplate(file, data)
	w.Write(page)
}

func StaticHandler(templateFile string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, r, templateFile, map[string]interface{}{})
	}
}

func getPage(r *http.Request) (page int, err error) {
	p := r.FormValue("p")
	page = 1

	if p != "" {
		page, err = strconv.Atoi(p)

		if err != nil {
			return
		}
	}

	return
}

/*
// URL: /comment/{contentId}
// 评论，不同内容共用一个评论方法
func CommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	user, _ := CurrentUser(r)

	vars := mux.Vars(r)
	contentId := vars["contentId"]

	var temp map[string]interface{}
	c := DB.C("contents")
	c.Find(bson.M{"_id": bson.ObjectIdHex(contentId)}).One(&temp)

	temp2 := temp["content"].(map[string]interface{})

	type_ := temp2["type"].(int)

	var url string
	switch type_ {
	case TypeArticle:
		url = "/a/" + contentId
	case TypeTopic:
		url = "/t/" + contentId
	case TypePackage:
		url = "/p/" + contentId
	}

	c.Update(bson.M{"_id": bson.ObjectIdHex(contentId)}, bson.M{"$inc": bson.M{"content.commentcount": 1}})

	content := r.FormValue("content")

	html := r.FormValue("html")
	html = strings.Replace(html, "<pre>", `<pre class="prettyprint linenums">`, -1)

	Id_ := bson.NewObjectId()
	now := time.Now()

	c = DB.C("comments")
	c.Insert(&Comment{
		Id_:       Id_,
		Type:      type_,
		ContentId: bson.ObjectIdHex(contentId),
		Markdown:  content,
		Html:      template.HTML(html),
		CreatedBy: user.Id_,
		CreatedAt: now,
	})

	if type_ == TypeTopic {
		// 修改最后回复用户Id和时间
		c = DB.C("contents")
		c.Update(bson.M{"_id": bson.ObjectIdHex(contentId)}, bson.M{"$set": bson.M{"latestreplierid": user.Id_.Hex(), "latestrepliedat": now}})
	}

	http.Redirect(w, r, url, http.StatusFound)
}

// URL: /comment/{commentId}/delete
// 删除评论
func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var commentId string = vars["commentId"]

	c := DB.C("comments")
	var comment Comment
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(commentId)}).One(&comment)

	if err != nil {
		Message(w, r, "评论不存在", "该评论不存在", "error")
		return
	}

	c.Remove(bson.M{"_id": comment.Id_})

	c = DB.C("contents")
	c.Update(bson.M{"_id": comment.ContentId}, bson.M{"$inc": bson.M{"content.commentcount": -1}})

	if comment.Type == TypeTopic {
		var topic Topic
		c.Find(bson.M{"_id": comment.ContentId}).One(&topic)
		if topic.LatestReplierId == comment.CreatedBy.Hex() {
			if topic.CommentCount == 0 {
				// 如果删除后没有回复，设置最后回复id为空，最后回复时间为创建时间
				c.Update(bson.M{"_id": topic.Id_}, bson.M{"$set": bson.M{"latestreplierid": "", "latestrepliedat": topic.CreatedAt}})
			} else {
				// 如果删除的是该主题最后一个回复，设置主题的最新回复id，和时间
				var latestComment Comment
				c = DB.C("comments")
				c.Find(bson.M{"contentid": topic.Id_}).Sort("-createdat").Limit(1).One(&latestComment)

				c = DB.C("contents")
				c.Update(bson.M{"_id": topic.Id_}, bson.M{"$set": bson.M{"latestreplierid": latestComment.CreatedBy.Hex(), "latestrepliedat": latestComment.CreatedAt}})
			}
		}
	}

	var url string
	switch comment.Type {
	case TypeArticle:
		url = "/a/" + comment.ContentId.Hex()
	case TypeTopic:
		url = "/t/" + comment.ContentId.Hex()
	case TypePackage:
		url = "/p/" + comment.ContentId.Hex()
	}

	http.Redirect(w, r, url, http.StatusFound)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	p := r.FormValue("p")
	page := 1

	if p != "" {
		var err error
		page, err = strconv.Atoi(p)

		if err != nil {
			Message(w, r, "页码错误", "页码错误", "error")
			return
		}
	}

	q := r.FormValue("q")

	keywords := strings.Split(q, " ")

	var noSpaceKeywords []string

	for _, keyword := range keywords {
		temp := strings.TrimSpace(keyword)
		if temp != "" {
			noSpaceKeywords = append(noSpaceKeywords, temp)
		}
	}

	fmt.Println(noSpaceKeywords, len(noSpaceKeywords))

	conditions := []bson.M{bson.M{"content.type": TypeTopic}}

	for _, keyword := range noSpaceKeywords {
		conditions = append(conditions, bson.M{"content.markdown": bson.M{"$regex": bson.RegEx{keyword, "i"}}})
	}

	c := DB.C("contents")

	var pagination *Pagination

	if len(noSpaceKeywords) == 0 {
		pagination = NewPagination(c.Find(bson.M{"content.type": TypeTopic}), "/search?"+q, PerPage)
	} else {
		pagination = NewPagination(c.Find(bson.M{"$and": conditions}), "/search?q="+q, PerPage)
	}

	var topics []Topic

	query, err := pagination.Page(page)
	if err != nil {
		message(w, r, "页码错误", "页码错误", "error")
		return
	}

	query.All(&topics)

	if err != nil {
		println(err.Error())
	}

	RenderTemplate(w, r, "search.html", map[string]interface{}{
		"q":          q,
		"topics":     topics,
		"pagination": pagination,
		"page":       page,
		"active":     "topic",
	})
}
*/
// 获取文件大小的接口
type Sizer interface {
	Size() int64
}

// 获取文件信息的接口
type Stat interface {
	Stat() (os.FileInfo, error)
}

type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

func ZoomAuto(file multipart.File, target *os.File, targetWidth int, ext string) {
	var img image.Image
	var err error
	switch ext {
	case ".gif":
		img, err = gif.Decode(file)
	case ".jpg":
		img, err = jpeg.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	}

	if err != nil {
		panic(err)
	}

	bounds := img.Bounds()

	var newWidth, newHight float32
	newWidth = float32(bounds.Dx())
	newHight = float32(bounds.Dy())
	//如果宽大于模版
	if newWidth > float32(targetWidth) {
		//宽按模版，高按比例缩放
		newWidth = float32(targetWidth)
		newHight = newHight * (float32(targetWidth) / float32(bounds.Dx()))
		m := resize.Resize(uint(newWidth), uint(newHight), img, resize.Lanczos3)

		switch ext {
		case ".gif":
			err = gif.Encode(target, m, &gif.Options{})
		case ".jpg":
			err = jpeg.Encode(target, m, &jpeg.Options{})
		case ".png":
			err = png.Encode(target, m)
		}
	} else {
		io.Copy(target, file)
	}

}
