package logger

import (
	"fmt"
	"os"
)

type ConsoleLogger struct {
	level int
}

func NewConsoleLogger() *ConsoleLogger {
	consoleLogger := &ConsoleLogger{}
	return consoleLogger
}

func (f *ConsoleLogger) SetLevel(level int) {
	if level < DEBUG || level > ERROR {
		level = DEBUG
	}
	f.level = level
}

func (f *ConsoleLogger) Debug(args ...interface{}) {
	f.SetLevel(DEBUG)
	logData := WriteLog(f.level, args...)
	fmt.Fprintf(os.Stdout, "%s %s %s:%d %s %s \n", logData.LevelStr,
		logData.TimeStr, logData.FileName, logData.LineNo,
		logData.FuncName, logData.Message)
}

func (f *ConsoleLogger) Info(args ...interface{}) {
	f.SetLevel(INFO)
	logData := WriteLog(f.level, args...)
	fmt.Fprintf(os.Stdout, "%s %s %s:%d %s %s \n", logData.LevelStr,
		logData.TimeStr, logData.FileName, logData.LineNo,
		logData.FuncName, logData.Message)
}

func (f *ConsoleLogger) Error(args ...interface{}) {
	f.SetLevel(ERROR)
	logData := WriteLog(f.level, args...)
	fmt.Fprintf(os.Stdout, "%s %s %s:%d %s %s \n", logData.LevelStr,
		logData.TimeStr, logData.FileName, logData.LineNo,
		logData.FuncName, logData.Message)
}

func (f *ConsoleLogger) Close() {
}
