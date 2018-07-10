package log

import (
	"log"
	"runtime"
	"fmt"
	"path/filepath"
)

type Level int

const (
	DEBUG   Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var callerDepth = 2
var levelFlag = [5]string{"DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}
var logger *log.Logger

func init() {
	file := openLogFile()
	logger = log.New(file, "", log.LstdFlags)
}

//设置日志信息前缀
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(callerDepth)
	var prex string
	if ok {
		prex = fmt.Sprintf("[%s] [%s:%d]", levelFlag[level], filepath.Base(file), line)
	} else {
		prex = fmt.Sprintf("[%s]", levelFlag[level])
	}
	logger.SetPrefix(prex)
}
//Debug
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}
//Info
func Info(v ...interface{})  {
	setPrefix(INFO)
	logger.Println(v)
}
//Warning
func Warning(v ...interface{})  {
	setPrefix(WARNING)
	logger.Println(v)
}
//Error
func Error(v ...interface{})  {
	setPrefix(ERROR)
	logger.Println(v)
}
//Fatal
func Fatal(v ...interface{})  {
	setPrefix(FATAL)
	logger.Println(v)
}
