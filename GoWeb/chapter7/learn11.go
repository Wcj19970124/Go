package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Article 用于解析XML的article标签的结构体
type Article struct {
	XMLName xml.Name `xml:"article"`
	ID      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	XML     string   `xml:",innerxml"`
}

//Author 作者结构体
type Author struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func learn11() {
	file, err := os.Open("article.xml")
	if err != nil {
		log.Println(err.Error())
	}

	defer file.Close() //最后需要关闭文件句柄

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err.Error())
	}

	var article Article
	err = xml.Unmarshal(data, &article)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(article)
}
