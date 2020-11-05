package main

import (
	"fmt"
	"net/url"
)

//url包主要是针对url的一些操作
//例如：对查询的残数进行转码，解码
//将url解析为一个结构体，便于获取url中的具体信息
//提供url.Values{},z专门用来针对url中的queryString,可以将queryString转为map类型，也可以将map拼接为queryString形式
func main() {
	str := url.QueryEscape("name=wcj") //将字符串转码为可以在URL中安全使用的
	fmt.Println(str)
	s, _ := url.QueryUnescape(str) //将由QueryEscape()转码的字符串解码为原字符串
	fmt.Println(s)

	urlStr := "http://www.wcj.cn/add/user?username=wcj&password=123"
	urlStruct, _ := url.Parse(urlStr) //将一个url解析为一个url.URL结构体，通过结构体可以获取其属性
	fmt.Println(urlStruct.Scheme, urlStruct.Opaque, urlStruct.User, urlStruct.Host, urlStruct.Path, urlStruct.RawQuery, urlStruct.RawFragment)
	fmt.Println(urlStruct.IsAbs()) //判断url是否是一个绝对url
	fmt.Println(urlStruct.Query()) //解析url.URL结构体中的RawQuery字段为一个键值对形式的url.Values{},就是解析queryString

	params, _ := url.ParseQuery(urlStr) //直接解析url中的queryString为一个url.Values{}
	fmt.Println(params)
}
