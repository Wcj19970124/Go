package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type student struct {
	ID   int
	Name string
	Age  int
}

//使用gob包，将数据以二进制流的形式存储和读取
func learn9() {
	stu := student{ID: 1, Name: "HHY", Age: 24}
	store(stu, "gob")
	var s student
	load(&s, "gob")
	fmt.Println(s)
}

//二进制形式读取数据
func load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename) //将编码的数据读取出来
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw) //创建缓冲区
	dec := gob.NewDecoder(buffer)  //创建解码器
	err = dec.Decode(data)         //将数据解码到传入的结构体或者其他存放的容器中
	if err != nil {
		panic(err)
	}
}

//二进制形式存储文件
func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)      //创建缓冲区
	encode := gob.NewEncoder(buffer) //创建编码器
	err := encode.Encode(data)       //对原始数据进行编码
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600) //将编码之后的数据写入文件
	if err != nil {
		panic(err)
	}
}
