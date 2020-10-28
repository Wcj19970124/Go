package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//文件的写入和读取的两种方法
func learn7() {

	//第一种：利用ioutil中提供的写和读的方法
	data := []byte("hello world!")
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	read1, _ := ioutil.ReadFile("data1")
	fmt.Println(string(read1))

	//第二种方法，利用os提供的方法
	file, err := os.Create("data2")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	bytes, err := file.Write(data)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(bytes)
	}

	file2, err := os.Open("data2")
	if err != nil {
		fmt.Println(err.Error())
	}
	read2 := make([]byte, len(data))
	bytes, err = file2.Read(read2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(read2))
}
