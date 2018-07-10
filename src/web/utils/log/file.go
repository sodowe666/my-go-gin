package log

import (
	"fmt"
	"os"
	"time"
	"log"
)

var (
	logSavePath = "runtime/logs/"
	logFileName = "app"
	logFileExt  = "log"
	timeFormat  = "20060102"
)

//获取日志目录
func getLogFileDir() string {
	return fmt.Sprintf("%s", logSavePath)
}

//获取日志名称
func getLogFileFullPath() string {
	relativeDir := getLogFileDir()
	fileName := fmt.Sprintf("%s_%s.%s", logFileName, time.Now().Format(timeFormat), logFileExt)
	return fmt.Sprintf("%s%s", relativeDir, fileName)// runtime/log/app_xxxx.log
}

//创建日志目录
func mkDir() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = os.MkdirAll(fmt.Sprintf("%s/%s", dir, getLogFileDir()), os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
}

//打开日志文件
func openLogFile() *os.File {
	_, err := os.Stat(getLogFileFullPath())
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission : %s", err.Error())
	}
	file, err := os.OpenFile(getLogFileFullPath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to Open logFlie:%v", err)
	}
	return file
}
