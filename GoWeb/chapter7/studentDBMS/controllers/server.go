package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"strconv"

	"../model"
)

//HandleGET 处理get请求的处理器函数
func HandleGET(w http.ResponseWriter, r *http.Request) error {

	log.Println("===== 进入处理器函数了哦 =====")

	//第一步：获取url上的id
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	log.Println("=== student id :" + strconv.Itoa(id))
	if err != nil {
		log.Println("get param id failed,err:" + err.Error())
		return err
	}
	//请求model
	stu, err := model.GetStudentByID(id)
	if err != nil {
		return err
	}

	//将获取到的student信息转化成json数据
	data, err := json.Marshal(stu)
	if err != nil {
		log.Println("marshal student info failed,err:" + err.Error())
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	return nil
}

//HandlePOST 处理post请求的处理器函数
func HandlePOST(w http.ResponseWriter, r *http.Request) error {

	log.Println("===== 进入处理器函数了哦 =====")

	//第一步：获取request.Body的长度
	len := r.ContentLength
	data := make([]byte, len)

	//将body中的数据读入切片
	r.Body.Read(data)

	log.Println(string(data))

	//构建对象用于接收切片中的信息
	var s model.Student
	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println("unmarshal student info failed,err:" + err.Error())
	}

	//调model接口更新
	_, err = model.InsertStudent(s)
	if err != nil {
		return err
	}
	return nil

}

//HandlePUT 处理put请求的处理器函数
func HandlePUT(w http.ResponseWriter, r *http.Request) error {

	log.Println("===== 进入处理器函数了哦 =====")

	//第一步：获取request.Body的长度
	len := r.ContentLength
	data := make([]byte, len)

	//将body中的数据读入切片
	r.Body.Read(data)

	//构建对象用于接收切片中的信息
	var s model.Student
	err := json.Unmarshal(data, &s)
	if err != nil {
		log.Println("unmarshal student info failed,err:" + err.Error())
	}

	//调model接口更新
	_, err = model.UpdateStudent(s)
	if err != nil {
		return err
	}

	return nil
}

//HandleDel 处理del请求的处理器函数
func HandleDel(w http.ResponseWriter, r *http.Request) error {

	log.Println("===== 进入处理器函数了哦 =====")

	//第一步：获取参数
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	log.Println("===== student id:" + strconv.Itoa(id))
	if err != nil {
		log.Println("get param id failed,err:" + err.Error())
		return err
	}

	//调model删除
	_, err = model.DeleteStudentByID(id)
	if err != nil {
		return err
	}

	//写响应头状态码
	w.WriteHeader(200)
	return nil
}
