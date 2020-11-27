package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	sex  string
}

//此包提供json数据的解码和编码，主要是四个方法
//Unmarshal()/Marshal():这种方法通常用于数据量不大，且是[]byte的情况下比较方便
//Decode()/Encode():这种方法通常用于数据量比较大的情况，且因为其和读取文件一样
//是创建编码器和解码器，所以对于数据形式的要求并不固定，只要能通过io.Reader/io.Writer读取和写入即可
func main() {

	str := `{
		"id":100,
		"name":"wcj",
		"price":125.58
	}`

	//Unmarshal解码
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(str), &m); err == nil {
		fmt.Println(m)
	} else {
		fmt.Println("解析json出错,err:" + err.Error())
	}

	for _, v := range m {
		if _, ok := v.(float64); ok {
			fmt.Println(v.(float64))
		}
	}

	//Marshal：编码/需要注意的是结构体必须为可导出的，且字段也要为可导出，不可导出的字段不会解析为json字符串
	// stu := Student{Id: 1, Name: "wcj", sex: "男"}
	// if data, err := json.Marshal(&stu); err == nil {
	// 	fmt.Println(string(data))
	// } else {
	// 	fmt.Println("编码出错,err:" + err.Error())
	// }

	//Decode：解码
	// tm := make(map[string]interface{})
	// dec := json.NewDecoder(strings.NewReader(str))
	// dec.UseNumber()
	// if dec.Decode(&tm) == nil {
	// 	fmt.Println(tm)
	// } else {
	// 	fmt.Println("解码出错")
	// }

	//Encode：编码
	// buf := bytes.Buffer{}
	// encoder := json.NewEncoder(&buf)
	// if encoder.Encode(stu) == nil {
	// 	fmt.Println(string(buf.Bytes()))
	// } else {
	// 	fmt.Println("编码失败")
	// }
}
