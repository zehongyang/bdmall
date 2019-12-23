package main

import (
	"bdmall/app"
	"bdmall/app/conf"
	"bdmall/app/queue"
	"log"
	"net/http"
)

func main() {
	router := app.InitRouter()
	//监听发送邮件队列任务
	go queue.Sub()
	s := http.Server{
		Addr:    conf.ServerConfig.ServerName,
		Handler: router,
	}
	log.Fatal(s.ListenAndServe())
}
