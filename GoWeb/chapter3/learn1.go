package main

import (
	"fmt"
	"net/http"
)

type student struct{}

func learn1() {

	stu := student{}
	server := http.Server{
		Addr: "127.0.0.1:8090",
	}

	http.Handle("/get/student", &stu)
	http.HandleFunc("/go/web", getGoWeb)

	server.ListenAndServe()
}

//ServeHTTP  实现此方法的即为处理器
//调用处理器，使用http.Handle()
func (s *student) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "get student info success!")
}

//处理器函数，即只要是和处理器拥有同样的签名的函数就是处理器函数
//Go中有一种处理器函数的类型，即handlerFun,当一个方法是处理器函数的时候
//会在使用的时候，默认进行强制类型转化内，将其转换为可执行的处理器函数
//调用处理器函数，使用http.HandleFunc()
func getGoWeb(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Go Web is good! But i don't want to study it")
}
