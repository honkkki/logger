# logger

🐱‍🏍🐱‍🏍🐱‍🏍

a simple web logger

## Install
```shell
go get -u github.com/honkkki/logger
```

## Usage
```go
package main

import (
	"github.com/honkkki/logger"
	"time"
)

func main() {
	logger.InitLogger()
	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	done := make(chan struct{})
	go func() {
		<-time.After(10*time.Second)
		done <- struct{}{}
	}()

	for {
		select {
		case <-tick.C:
			logger.Info("logging")
		case <-done:
			logger.Info("finish")
			time.Sleep(time.Second)
			return
		}
	}
}
```
