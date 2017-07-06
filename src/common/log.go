package common

import (
	"fmt"
	"os"
	"time"

	l4g "github.com/alecthomas/log4go"
)

func WriteLog(msg string) {
	l4g.Debug(msg)

	//defer l4g.Close()
}

func init() {
	dir, fileName := "log", time.Now().Format("20060102")+".log"
	_, err := os.Stat(dir)
	if err != nil {
		os.Mkdir(dir, os.ModePerm)
	}

	l4g.AddFilter("stdout", l4g.DEBUG, l4g.NewConsoleLogWriter())                                                         //输出到控制台,级别为DEBUG
	l4g.AddFilter("file", l4g.DEBUG, l4g.NewFileLogWriter(fmt.Sprintf("%s%c%s", dir, os.PathSeparator, fileName), false)) //输出到文件,级别为DEBUG,文件名为test.log,每次追加该原文件
}
