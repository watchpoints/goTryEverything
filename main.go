package main

import (
	"fmt"
	"log"
	"net/http"

	"gitee.com/wang_cyi/goTryEverything/oneday"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>hello world</h1>")
}

//http://localhost:1984/
func start_http_server() {
	//构建http服务
	mux := http.NewServeMux()
	http.HandleFunc("/", IndexHandler)

	fmt.Println("TryEverthing start ")
	ip := "localhost:1984"
	log.Println(ip, "start")
	err := http.ListenAndServe(ip, mux)
	if err != nil {
		log.Println(err)
	}
}
func start_job_server() {
	fmt.Println("start_job_server")
	dailyAction := &oneday.DailyAction{
		Name: "daily aciton",
	}
	dailyAction.StartCron()

}

func main() {
	log.Println("start GoTryEverthing")
	go start_job_server()
	start_http_server() //forever
	log.Println("stop GoTryEverthing")

}
