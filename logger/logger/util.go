package logger

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

type LogData struct {
	Message  string
	TimeStr  string
	LevelStr string
	FileName string
	FuncName string
	LineNo   int
}

func GetLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}

	return
}

// WriteLog 封装每条日志信息
func WriteLog(LogLevel int, args ...interface{}) *LogData {
	now := time.Now()
	nowTime := now.Format("2006-01-02 15:04:05")

	levelStr := ""
	switch LogLevel {
	case DEBUG:
		levelStr = "[DEBUG] "
	case INFO:
		levelStr = "[INFO] "
	case ERROR:
		levelStr = "[ERROR] "
	default:
		levelStr = "[DEBUG] "
	}

	nowTime = nowTime + " "
	fileName, funcName, lineNo := GetLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	msg := fmt.Sprint(args...)

	logData := &LogData{
		Message: msg,
		TimeStr:  nowTime,
		LevelStr: levelStr,
		FileName: fileName,
		FuncName: funcName,
		LineNo:   lineNo,
	}

	return logData
}
