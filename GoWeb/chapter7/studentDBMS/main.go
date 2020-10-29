package main

import (
	"log"
	"net/http"

	"./controllers"
)

//多路复用器，负责分发请求到不同的controller
func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/request/", handleRequest)
	server.ListenAndServe()

}

//处理器函数 ---充当多路复用器
func handleRequest(w http.ResponseWriter, r *http.Request) {

	log.Println("=====  进入多路复用器了哦  =====")

	var err error

	switch r.Method {
	case "GET":
		err = controllers.HandleGET(w, r)
	case "POST":
		err = controllers.HandlePOST(w, r)
	case "PUT":
		err = controllers.HandlePUT(w, r)
	case "DELETE":
		err = controllers.HandleDel(w, r)
	}

	if err != nil {
		log.Println("request failed,err:" + err.Error())
		return
	}
}
