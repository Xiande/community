// Community project main.go
package main

import (
	//"DBBLL"
	"common"
	"community"
	//"fmt"
	//"qiniupkg.com/api.v7/conf"
	//"qiniupkg.com/api.v7/kodo"
	//"qiniupkg.com/api.v7/kodocli"
	//"database/sql"
)

func main() {
	//bll := DBBLL.NewUserBLL(common.Config.DB, common.Config.DBConn)
	//s := bll.GetStatic("guxiande")
	//m, c := bll.GetPagingList(1, 10, "", "", "", "", "all", "")
	//fmt.Println(s)
	community.StartServer()
	/*
		bucket := "xiande"
		key := "community/photos/a.png"

		conf.ACCESS_KEY = common.Config.QiniuAccessKey
		conf.SECRET_KEY = common.Config.QiniuSecretKey
		c := kodo.New(0, nil)
		policy := &kodo.PutPolicy{
			Scope: bucket + ":" + key,
			//设置Token过期时间
			Expires: 3600,
		}
		//生成一个上传token
		//"UY1crWGZD6u94QGSqnd6tDLF_v7gYdZIdIISEjzZ:1Mp0LneIgSILV-iUiLKCZxlBuWk=:eyJzY29wZSI6InhpYW5kZSIsImRlYWRsaW5lIjoxNDc4MjQ4MTc5fQ=="
		token := c.MakeUptoken(policy)
		fmt.Println(token)
		//构建一个uploader
		zone := 0
		uploader := kodocli.NewUploader(zone, nil)

		var ret common.PutRet
		//调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名
		filepath := "a.png"

		res := uploader.PutFile(nil, &ret, token, key, filepath, nil)
		//res := uploader.PutWithoutKey(nil, &ret, token, formFile, fileSize, nil)
		//打印返回的信息
		fmt.Println(ret)
		//打印出错信息

		if res != nil {
			fmt.Println("upload to qiniu failed:", res.Error())
		}
	*/
	defer func() {
		if x := recover(); x != nil {
			common.WriteLog(x.(error).Error())

		}
	}()
}
