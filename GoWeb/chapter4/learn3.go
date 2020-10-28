package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Teacher  数据体
type Teacher struct {
	ID     int
	Name   string
	Age    int
	Gender string
}

//总结：对于客户端请求的响应消息，即http应答报文，我们可以通过http.ResponseWriter进行修改
//Write()方法是将信息写入到应答消息的消息体中
//WriteHeader()方法是写入应答的状态码进入到应答消息，但是修改了这个之后，后面就无法对头部消息进行操作了，所以通常应该放在最后
//Header()用于对应答消息的头部进行操作，可以增删改查
//另外重定向的情况，需要将状态码设置为302
//输出json数据的情况，需要将头部中的Content-Type字段设置为application/json
func learn3() {

	server := http.Server{
		Addr: "127.0.0.1:8110",
	}

	http.HandleFunc("/get/request/write", writeExample)
	http.HandleFunc("/get/request/header", writeHeader)
	http.HandleFunc("/redirect", writeHeaderSet)
	http.HandleFunc("/get/stu/json", writeJSON)

	server.ListenAndServe()
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
    <head>
        <title>Go Web Programming</title>
    </head>
    <body>
        <h1>Hello Go</h1>
    </body>
	</html>`

	w.Write([]byte(str)) //responseWriter的write方法用于写信息到HTTP响应的主体中
}

func writeHeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501) // responseWriter的WriteHeader方法用于写请求的返回状态码

	fmt.Fprintf(w, "write header success")
}

func writeHeaderSet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "www.baidu.com") //Header是用来设置响应的头部信息，但是需要注意的是如果使用了WriteHeader()设置
	//了状态码的话，那么之后是不允许对头部信息进行修改的
	w.WriteHeader(302)
}

//测试JSON数据的输出
func writeJSON(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	t := Teacher{
		ID:     20976303,
		Name:   "wcj",
		Age:    24,
		Gender: "男",
	}

	resp, _ := json.Marshal(&t)

	w.Write(resp)
}
