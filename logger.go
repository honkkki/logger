package logger

import (
	"fmt"
	"strconv"
	"time"
)

var logger LoggerInterface
var config = make(map[string]string)

// 初始化基本配置信息
func init() {
	now := time.Now()
	config["type"] = "file"
	config["fp"] = "./log_file"
	config["fn"] = fmt.Sprintf("%04d%02d%02d", now.Year(), now.Month(), now.Day())
	config["log_chan_size"] = "1000"
}

func InitLogger() {
	switch config["type"] {
	case "file":
		chanSize, err := strconv.Atoi(config["log_chan_size"])
		if err != nil {
			chanSize = 1000
		}
		logger = NewFileLogger(config["fp"], config["fn"], chanSize)
	case "console":
		logger = NewConsoleLogger()
	}
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Close() {
	logger.Close()
}
