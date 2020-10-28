package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

//Blog 博客数据体
type Blog struct {
	XMLName  xml.Name  `xml:"blog"`
	ID       int       `xml:"id,attr"`
	Content  string    `xml:"content"`
	Writer   Writer    `xml:"writer"`
	XML      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

//Writer 数据体
type Writer struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

//Comment 数据体
type Comment struct {
	ID      string `xml:"id,attr"`
	Content string `xml:"content"`
	Writer  Writer `xml:"writer"`
}

func learn12() {

	unMarshalXML()
	decodeXML()
}

//Unmarshal和Marshal这种方法只适用于解析byte切片
//且体积较小的内容，因此对于体积大的或者流形式的无法做到
//高效的解析了，因此xml和json包提供了另外两种方法，即Decode()和Encode()
func unMarshalXML() {
	xmlFile, err := os.Open("comment.xml")
	if err != nil {
		log.Println(err.Error())
	}

	defer xmlFile.Close()
	data, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Println(err.Error())
	}

	var blog Blog
	err = xml.Unmarshal(data, &blog)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(blog)
}

//Decode()和Encode()方法能够高效的解析体积庞大，流式的数据
func decodeXML() {
	xmlFile, err := os.Open("comment.xml")
	if err != nil {
		log.Println(err.Error())
	}

	defer xmlFile.Close()
	decoder := xml.NewDecoder(xmlFile) //创建解码器
	for {
		t, err := decoder.Token()
		if t == io.EOF {
			break
		}

		if err != nil {
			log.Println(err.Error())
			return
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se)
				fmt.Println(comment)
			}
		}
	}
}
