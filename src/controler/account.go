/*
处理用户相关的操作,注册,登录,验证,等等
*/
package controler

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strings"
	//"time"

	"DBBLL"
	"DBModel"
	"common"

	"code.google.com/p/go-uuid/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jimmykuu/wtforms"
	. "qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
)

var defaultPhotos = []string{
	"default_boy.png",
	"default_girl.png",
}

// 加密密码,转成md5
func encryptPassword(password string) string {
	h := md5.New()
	io.WriteString(h, password)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// URL: /signup
// 处理用户注册,要求输入用户名,密码和邮箱
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	form := wtforms.NewForm(
		wtforms.NewTextField("username", "用户名", "", wtforms.Required{}, wtforms.Regexp{Expr: `^[a-zA-Z0-9_]{3,16}$`, Message: "请使用a-z, A-Z, 0-9以及下划线, 长度3-16之间"}),
		wtforms.NewTextField("displayname", "显示名", "", wtforms.Required{}, wtforms.Regexp{Expr: `^[a-zA-Z0-9_]{3,16}$`, Message: "请使用a-z, A-Z, 0-9以及下划线, 长度3-16之间"}),
		wtforms.NewPasswordField("password", "密码", wtforms.Required{}),
		wtforms.NewTextField("email", "电子邮件", "", wtforms.Required{}, wtforms.Email{}),
	)

	if r.Method == "POST" {
		if form.Validate(r) {
			bll := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn)
			username := form.Value("username")
			pwd := form.Value("password")
			displayname := form.Value("displayname")
			email := form.Value("email")
			//fmt.Println(username)
			// 检查用户名
			u := bll.GetByName(username)
			if u != nil {
				//fmt.Println(u.DisplayName)
				form.AddError("username", "该用户名已经被注册")

				common.RenderTemplate(w, r, "account/signup.html", map[string]interface{}{"form": form})
				return
			}

			// 检查邮箱
			user_mail := bll.GetByName(email)

			if user_mail != nil {
				form.AddError("email", "电子邮件地址已经被注册")

				common.RenderTemplate(w, r, "account/signup.html", map[string]interface{}{"form": form})
				return
			}

			var user DBModel.User
			user.Username = username
			user.Password = encryptPassword(pwd)
			user.DisplayName = displayname
			user.Email = email
			user.IsActive = true
			user.IsModerator = false
			user.IsSuperuser = false
			bll.Add(user)

			// 注册成功后设成登录状态
			session, _ := common.Store.Get(r, "user")
			session.Values["username"] = username
			session.Save(r, w)

			// 跳到修改用户信息页面
			http.Redirect(w, r, "/profile", http.StatusFound)
			return
		}
	}

	common.RenderTemplate(w, r, "account/signup.html", map[string]interface{}{"form": form})
}

// URL: /activate/{code}
// 用户根据邮件中的链接进行验证,根据code找到是否有对应的用户,如果有,修改User.IsActive为true
func ActivateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	bll := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn)
	var user *DBModel.User
	user = bll.GetByCode(code)

	if user == nil {
		common.Message(w, r, "没有该验证码", "请检查连接是否正确", "error")
		return
	}

	user.IsActive = true
	user.ValidateCode = ""
	bll.Update(*user)

	common.Message(w, r, "通过验证", `恭喜你通过验证,请 <a href="/signin">登录</a>.`, "success")
}

// URL: /signin
// 处理用户登录,如果登录成功,设置Cookie
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	next := r.Form.Get("source")

	form := wtforms.NewForm(
		wtforms.NewHiddenField("next", next),
		wtforms.NewTextField("username", "用户名", "", &wtforms.Required{}),
		wtforms.NewPasswordField("password", "密码", &wtforms.Required{}),
	)

	if r.Method == "POST" {
		if form.Validate(r) {
			username := form.Value("username")
			bll := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn)
			var user *DBModel.User
			user = bll.GetByName(username)
			if user == nil {
				form.AddError("username", "该用户不存在")
				common.RenderTemplate(w, r, "account/signin.html", map[string]interface{}{"form": form})
				return
			}

			if !user.IsActive {
				form.AddError("username", "邮箱没有经过验证,如果没有收到邮件,请联系管理员")
				common.RenderTemplate(w, r, "account/signin.html", map[string]interface{}{"form": form})
				return
			}

			if user.Password != encryptPassword(form.Value("password")) {
				form.AddError("password", "密码和用户名不匹配")

				common.RenderTemplate(w, r, "account/signin.html", map[string]interface{}{"form": form})
				return
			}

			session, _ := common.Store.Get(r, "user")
			session.Values["username"] = user.Username
			session.Save(r, w)

			if form.Value("next") == "" {
				http.Redirect(w, r, "/", http.StatusFound)
			} else {
				http.Redirect(w, r, next, http.StatusFound)
			}

			return
		}
	}

	common.RenderTemplate(w, r, "account/signin.html", map[string]interface{}{"form": form})
}

// URL: /signout
// 用户登出,清除Cookie
func SignoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := common.Store.Get(r, "user")
	session.Options = &sessions.Options{MaxAge: -1}
	session.Save(r, w)
	//common.RenderTemplate(w, r, "account/signout.html", map[string]interface{}{"signout": true})
	http.Redirect(w, r, "/signin", http.StatusFound)
}

// URL /profile
// 用户设置页面,显示用户设置,用户头像,密码修改
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := common.CurrentUser(r)

	profileForm := wtforms.NewForm(
		wtforms.NewTextField("displayname", "显示名", user.DisplayName, wtforms.Required{}, wtforms.Regexp{Expr: `^[a-zA-Z0-9_]{3,16}$`, Message: "请使用a-z, A-Z, 0-9以及下划线, 长度3-16之间"}),
		wtforms.NewTextField("website", "个人网站", user.Website),
		wtforms.NewTextField("location", "所在地", user.Location),
		wtforms.NewTextField("tagline", "签名", user.Tagline),
		wtforms.NewTextArea("bio", "个人简介", user.Bio),
		wtforms.NewTextField("github_username", "GitHub用户名", user.GitHubUsername),
		wtforms.NewTextField("weibo", "新浪微博", user.Weibo),
	)

	if r.Method == "POST" {
		if profileForm.Validate(r) {
			bll := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn)
			user.DisplayName = profileForm.Value("displayname")
			user.Website = profileForm.Value("website")
			user.Location = profileForm.Value("location")
			user.Tagline = profileForm.Value("tagline")
			user.Bio = profileForm.Value("bio")
			user.GitHubUsername = profileForm.Value("github_username")
			user.Weibo = profileForm.Value("weibo")
			bll.Update(*user)
			http.Redirect(w, r, "/profile", http.StatusFound)
			return
		}
	}

	common.RenderTemplate(w, r, "account/profile.html", map[string]interface{}{
		"user":          user,
		"profileForm":   profileForm,
		"defaultPhotos": defaultPhotos,
	})
}

// URL: /profile/avatar
// 修改头像,提交到七牛云存储
func ChangePhotoHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := common.CurrentUser(r)

	if r.Method == "POST" {
		formFile, formHeader, err := r.FormFile("file")
		if err != nil {
			fmt.Println("changePhotoHandler:", err.Error())
			common.RenderTemplate(w, r, "account/Photo.html", map[string]interface{}{
				"user":  user,
				"error": "请选择图片上传",
			})
			return
		}

		defer formFile.Close()

		// 检查是否是jpg或png文件
		uploadFileType := formHeader.Header["Content-Type"][0]

		isValidateType := false
		for _, imgType := range []string{"image/png", "image/jpeg"} {
			if imgType == uploadFileType {
				isValidateType = true
				break
			}
		}

		if !isValidateType {
			fmt.Println("upload image type error:", uploadFileType)
			// 提示错误
			common.RenderTemplate(w, r, "account/Photo.html", map[string]interface{}{
				"user":  user,
				"error": "文件类型错误，请选择jpg/png图片上传。",
			})
			return
		}

		// 检查文件尺寸是否在500K以内
		fileSize := formFile.(common.Sizer).Size()

		if fileSize > 500*1024 {
			// > 500K
			fmt.Printf("upload image size > 500K: %dK\n", fileSize/1024)
			common.RenderTemplate(w, r, "account/Photo.html", map[string]interface{}{
				"user":  user,
				"error": "图片大小大于500K，请选择500K以内图片上传。",
			})
			return
		}

		filenameExtension := ".jpg"
		if uploadFileType == "image/png" {
			filenameExtension = ".png"
		}

		// 文件名：32位uuid，不带减号和后缀组成
		filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + filenameExtension

		bucket := "xiande"
		key := "community/photos/" + filename

		ACCESS_KEY = common.Config.QiniuAccessKey
		SECRET_KEY = common.Config.QiniuSecretKey

		c := kodo.New(0, nil)
		policy := &kodo.PutPolicy{
			Scope: bucket + ":" + key,
			//设置Token过期时间
			Expires: 3600,
		}

		//生成一个上传token
		token := c.MakeUptoken(policy)
		//构建一个uploader
		zone := 0
		uploader := kodocli.NewUploader(zone, nil)

		var ret common.PutRet
		//调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名
		//res := uploader.PutFile(nil, &ret, token, key, "a.png", nil)
		res := uploader.Put(nil, &ret, token, key, formFile, fileSize, nil)
		//res := uploader.PutWithoutKey(nil, &ret, token, formFile, fileSize, nil)
		//打印返回的信息
		fmt.Println(ret)
		//打印出错信息

		if res != nil {
			fmt.Println("upload to qiniu failed:", res.Error())
			common.RenderTemplate(w, r, "account/Photo.html", map[string]interface{}{
				"user":  user,
				"error": "上传失败，请反馈错误",
			})
			return
		}

		// 存储远程文件名
		user.PhotoUrl = filename
		bll := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn)
		bll.Update(*user)

		http.Redirect(w, r, "/profile#photo", http.StatusFound)
		return
	}

	common.RenderTemplate(w, r, "account/Photo.html", map[string]interface{}{"user": user})
}

// URL: /profile/choose_default_Photo
// 选择默认头像
func ChooseDefaultPhoto(w http.ResponseWriter, r *http.Request) {
	user, _ := common.CurrentUser(r)

	if r.Method == "POST" {
		photo := r.FormValue("defaultPhotos")

		if photo != "" {
			user.PhotoUrl = photo
			bll := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn)
			bll.Update(*user)
		}

		http.Redirect(w, r, "/profile#photo", http.StatusFound)
	}
}

/*
func FollowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	currUser, _ := common.CurrentUser(r)

	//不能关注自己
	if currUser.Username == username {
		common.Message(w, r, "提示", "不能关注自己", "error")
		return
	}

	user := User{}
	c := DB.C("users")
	err := c.Find(bson.M{"username": username}).One(&user)

	if err != nil {
		common.Message(w, r, "关注的会员未找到", "关注的会员未找到", "error")
		return
	}

	if user.IsFollowedBy(currUser.Username) {
		common.Message(w, r, "你已经关注该会员", "你已经关注该会员", "error")
		return
	}
	c.Update(bson.M{"_id": user.Id_}, bson.M{"$push": bson.M{"fans": currUser.Username}})
	c.Update(bson.M{"_id": currUser.Id_}, bson.M{"$push": bson.M{"follow": user.Username}})

	http.Redirect(w, r, "/member/"+user.Username, http.StatusFound)
}

func UnfollowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	currUser, _ := CurrentUser(r)

	//不能取消关注自己
	if currUser.Username == username {
		common.Message(w, r, "提示", "不能对自己进行操作", "error")
		return
	}

	user := User{}
	c := DB.C("users")
	err := c.Find(bson.M{"username": username}).One(&user)

	if err != nil {
		common.Message(w, r, "没有该会员", "没有该会员", "error")
		return
	}

	if !user.IsFollowedBy(currUser.Username) {
		common.Message(w, r, "不能取消关注", "该会员不是你的粉丝,不能取消关注", "error")
		return
	}

	c.Update(bson.M{"_id": user.Id_}, bson.M{"$pull": bson.M{"fans": currUser.Username}})
	c.Update(bson.M{"_id": currUser.Id_}, bson.M{"$pull": bson.M{"follow": user.Username}})

	http.Redirect(w, r, "/member/"+user.Username, http.StatusFound)
}

*/
/*
// URL: /forgot_password
// 忘记密码,输入用户名和邮箱,如果匹配,发出邮件
func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	form := wtforms.NewForm(
		wtforms.NewTextField("username", "用户名", "", wtforms.Required{}),
		wtforms.NewTextField("email", "电子邮件", "", wtforms.Email{}),
	)

	if r.Method == "POST" {
		if form.Validate(r) {
			var user User
			c := DB.C("users")
			err := c.Find(bson.M{"username": form.Value("username")}).One(&user)
			if err != nil {
				form.AddError("username", "没有该用户")
			} else if user.Email != form.Value("email") {
				form.AddError("username", "用户名和邮件不匹配")
			} else {
				message2 := `Hi %s,
我们的系统收到一个请求，说你希望通过电子邮件重新设置你在 Golang中国 的密码。你可以点击下面的链接开始重设密码：

%s/reset/%s

如果这个请求不是由你发起的，那没问题，你不用担心，你可以安全地忽略这封邮件。

如果你有任何疑问，可以回复这封邮件向我提问。`
				code := strings.Replace(uuid.NewUUID().String(), "-", "", -1)
				c.Update(bson.M{"_id": user.Id_}, bson.M{"$set": bson.M{"resetcode": code}})
				message2 = fmt.Sprintf(message2, user.Username, Config.Host, code)
				sendMail("[Golang中国]重设密码", message2, []string{user.Email})
				message(w, r, "通过电子邮件重设密码", "一封包含了重设密码指令的邮件已经发送到你的注册邮箱，按照邮件中的提示，即可重设你的密码。", "success")
				return
			}
		}
	}

	renderTemplate(w, r, "account/forgot_password.html", map[string]interface{}{"form": form})
}

// URL: /reset/{code}
// 用户点击邮件中的链接,根据code找到对应的用户,设置新密码,修改完成后清除code
func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	var user User
	c := DB.C("users")
	err := c.Find(bson.M{"resetcode": code}).One(&user)

	if err != nil {
		message(w, r, "重设密码", `无效的重设密码标记,可能你已经重新设置过了或者链接已经失效,请通过<a href="/forgot_password">忘记密码</a>进行重设密码`, "error")
		return
	}

	form := wtforms.NewForm(
		wtforms.NewPasswordField("new_password", "新密码", wtforms.Required{}),
		wtforms.NewPasswordField("confirm_password", "确认新密码", wtforms.Required{}),
	)

	if r.Method == "POST" && form.Validate(r) {
		if form.Value("new_password") == form.Value("confirm_password") {
			c.Update(
				bson.M{"_id": user.Id_},
				bson.M{
					"$set": bson.M{
						"password":  encryptPassword(form.Value("new_password")),
						"resetcode": "",
					},
				},
			)
			message(w, r, "重设密码成功", `密码重设成功,你现在可以 <a href="/signin" class="btn btn-primary">登录</a> 了`, "success")
			return
		} else {
			form.AddError("confirm_password", "密码不匹配")
		}
	}

	renderTemplate(w, r, "account/reset_password.html", map[string]interface{}{"form": form, "code": code, "account": user.Username})
}

type Sizer interface {
	Size() int64
}



// URL: /change_password
// 修改密码
func changePasswordHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := CurrentUser(r)

	form := wtforms.NewForm(
		wtforms.NewPasswordField("current_password", "当前密码", wtforms.Required{}),
		wtforms.NewPasswordField("new_password", "新密码", wtforms.Required{}),
		wtforms.NewPasswordField("confirm_password", "新密码确认", wtforms.Required{}),
	)

	if r.Method == "POST" && form.Validate(r) {
		if form.Value("new_password") == form.Value("confirm_password") {
			currentPassword := encryptPassword(form.Value("current_password"))
			if currentPassword == user.Password {
				c := DB.C("users")
				c.Update(bson.M{"_id": user.Id_}, bson.M{"$set": bson.M{"password": encryptPassword(form.Value("new_password"))}})
				message(w, r, "密码修改成功", `密码修改成功`, "success")
				return
			} else {
				form.AddError("current_password", "当前密码错误")
			}
		} else {
			form.AddError("confirm_password", "密码不匹配")
		}
	}

	renderTemplate(w, r, "account/change_password.html", map[string]interface{}{"form": form})
}
*/
