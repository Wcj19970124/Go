package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)
	http.HandleFunc("/condition", condition)
	http.HandleFunc("/iteration", iteration)
	http.HandleFunc("/setting", setting)
	http.HandleFunc("/nest", nest)
	http.HandleFunc("/channel", channel)
	http.HandleFunc("/function", function)
	http.HandleFunc("/xss", xss)
	http.HandleFunc("/form", form)
	server.ListenAndServe()
}

//基本显示模板文件
func process(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("temp.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	file.Execute(w, "hello world")
}

//模板条件动作
func condition(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("temp.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	rand.Seed(time.Now().Unix())
	file.Execute(w, rand.Intn(10) > 5)
}

//迭代动作
func iteration(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("temp.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	dayOfWeeks := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	file.Execute(w, dayOfWeeks)
}

//设置动作
func setting(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("temp.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	file.Execute(w, "hello")
}

//嵌套模板
func nest(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("temp.html", "temp2.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	file.Execute(w, "hello")
}

//管道
func channel(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("temp.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	m := make(map[string]interface{})
	m["name"] = "wcj"
	m["age"] = 24
	m["gender"] = "mail"

	file.Execute(w, m)
}

//模板支持自定义函数
func function(w http.ResponseWriter, r *http.Request) {

	funcMap := template.FuncMap{"fdate": formdate}

	file := template.New("temp.html").Funcs(funcMap)
	file, _ = file.ParseFiles("temp.html")
	file.Execute(w, time.Now())
}

func formdate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

//测试xss攻击
func xss(w http.ResponseWriter, r *http.Request) {
	fmt.Println("进入攻击处理器函数~~~~~~~~~")
	fmt.Println(r.FormValue("comment"))
	file, _ := template.ParseFiles("temp.html")
	file.Execute(w, template.HTML(r.FormValue("comment")))
}

func form(w http.ResponseWriter, r *http.Request) {
	file, _ := template.ParseFiles("xss.html")
	file.Execute(w, nil)
}
