package main

import (
	"go-practice/logger/logger"
	"log"
	"net/http"
)

func main() {
	logger.InitLogger()
	//for {
	//	logger.Debug("user server is running")
	//	time.Sleep(time.Second)
	//}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/favicon.ico" {
		logger.Debug("hello http server")
	}
}
