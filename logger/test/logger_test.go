package test

import (
	"fmt"
	"go-practice/logger/logger"
	"testing"
	"time"
)

func TestFileLogger(t *testing.T) {
	fmt.Println("test file logger")
	log := logger.NewFileLogger("../log_file", "test", 1000)
	defer log.Close()
	log.Debug("log debug")
	log.Info("log info")
	log.Error("log error")
	log.Info([]string{
		"hello", "golang",
	})

	// 等待协程执行完任务
	time.Sleep(time.Second * 3)
}

func TestConsoleLogger(t *testing.T) {
	log := logger.NewConsoleLogger()
	log.Debug("log console debug")
}