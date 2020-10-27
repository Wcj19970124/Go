package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func learn2() {
	server := http.Server{
		Addr: "127.0.0.1:8100",
	}

	http.HandleFunc("/process", getRequestInfo)

	server.ListenAndServe()
}

//获取request信息查看
//对Post请求的表单和Url中的参数信息进行解析
//request.Form解析获得的是一个map，并且其中的值包含url的参数和表单中的参数
//request.PostForm则只会解析出来表单中的参数
func getRequestInfo(w http.ResponseWriter, r *http.Request) {

	// r.ParseForm()              //此方法会将Url和表单中的参数全部解析到Form字段中,Form字段是一个map
	// fmt.Println(w, r.PostForm) ///PostForm字段中只有解析的表单中的字段属性，没有url中的字段属性

	// r.ParseMultipartForm(1024)      //此方法用于解析表单提交数据为multipart-form类型的数据
	// fmt.Println(w, r.MultipartForm) //MultipartForm字段不仅包含有表单中的属性，同时内部还有一个map用来存放上传的文件

	// fmt.Println(r.FormValue("hello"))     //不需要先对数据进行解析，直接可以获取url或者表单中的数据，不过指挥获取map中的第一个，但是map中的数据顺序并不固定
	// fmt.Println(r.PostFormValue("hello")) //只能获取表单中的数据，无法获取url中的参数

	// r.ParseMultipartForm(1024) //获取表单上传的文件，首先对文件进行解析，然后将解析的文件打开，在读取即可
	// fileHeader := r.MultipartForm.File["upload"][0]
	// file, err := fileHeader.Open()
	// if err == nil {
	// 	data, err := ioutil.ReadAll(file)
	// 	if err == nil {
	// 		fmt.Fprintf(w, string(data))
	// 	}
	// }

	file, _, err := r.FormFile("upload") //简便的获取上传文件的方法，不需要对数据进行解析操作
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintf(w, string(data))
		}
	}
}
