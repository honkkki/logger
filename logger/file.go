package logger

import (
	"fmt"
	"os"
	"time"
)

type FileLogger struct {
	level       int
	filePath    string
	fileName    string
	file        *os.File
	LogDataChan chan *LogData
}

func NewFileLogger(fp, fn string, logChanSize int) *FileLogger {
	fileLogger := &FileLogger{
		filePath:    fp,
		fileName:    fn,
		LogDataChan: make(chan *LogData, logChanSize),
	}

	fileLogger.initFile()
	go fileLogger.writeLogBg()
	return fileLogger
}

func (f *FileLogger) initFile() {
	os.MkdirAll("./log_file", 0755)
	fileAllName := fmt.Sprintf("%s/%s.log", f.filePath, f.fileName)
	file, err := os.OpenFile(fileAllName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed, err : %s", fileAllName, err.Error()))
	}

	f.file = file
}

// 根据日期创建日志文件
func (f *FileLogger) checkSplitDate()  {
	now := time.Now()
	nowDate := fmt.Sprintf("%04d%02d%02d", now.Year(), now.Month(), now.Day())

	if nowDate == f.fileName {
		return
	} else {
		// 新的日期创建新的文件
		f.file.Close()
		fileAllName := fmt.Sprintf("%s/%s.log", f.filePath, nowDate)
		file, err := os.OpenFile(fileAllName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
		if err != nil {
			panic(fmt.Sprintf("open file %s failed, err : %s", fileAllName, err.Error()))
		}

		f.file = file
	}
}

// 写日志goroutine
func (f *FileLogger) writeLogBg() {
	for logData := range f.LogDataChan {
		fmt.Fprintf(f.file, "%s %s %s:%d %s %s \n", logData.LevelStr,
			logData.TimeStr, logData.FileName, logData.LineNo,
			logData.FuncName, logData.Message)
	}
}

func (f *FileLogger) SetLevel(level int) {
	if level < DEBUG || level > ERROR {
		level = DEBUG
	}
	f.level = level
}

func (f *FileLogger) Debug(args ...interface{}) {
	f.SetLevel(DEBUG)
	logData := WriteLog(f.level, args...)
	f.checkSplitDate()
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Info(args ...interface{}) {
	f.SetLevel(INFO)
	logData := WriteLog(f.level, args...)
	f.checkSplitDate()
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Error(args ...interface{}) {
	f.SetLevel(ERROR)
	logData := WriteLog(f.level, args...)
	f.checkSplitDate()
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Close() {
	f.file.Close()
}
