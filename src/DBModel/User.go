package DBModel

import (
	"time"
)

const (
	DefaultPhoto = "default_boy.png"
)

// 用户
type User struct {
	Id             int `beedb:"PK"`
	Username       string
	DisplayName    string
	Password       string
	Email          string
	PhotoUrl       string
	Website        string
	Location       string
	Tagline        string
	Bio            string
	Twitter        string
	Weibo          string
	GitHubUsername string
	JoinedAt       time.Time
	Follow         string
	Fans           string
	IsSuperuser    bool
	IsActive       bool
	ValidateCode   string
	ResetCode      string
	UIndex         int
	IsModerator    bool
}

// 是否是默认头像
func (u *User) IsDefaultPhoto(photo string) bool {
	filename := u.PhotoUrl
	if filename == "" {
		filename = DefaultPhoto
	}

	return filename == photo
}

func (u *User) PhotoImgSrc() string {
	// 如果没有设置头像，用默认头像
	filename := u.PhotoUrl
	if filename == "" {
		filename = DefaultPhoto
	}

	return "http://og3qrxo6x.bkt.clouddn.com/community/photos/" + filename
}
