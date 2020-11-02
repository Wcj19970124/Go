package main

import (
	"fmt"
	"strings"
)

//strings包提供了一些基础的字符串操作方法：并且提供了将字符串转换为Reader,以及Replace替换字符串中
//字符的方法
func main() {
	var a = "烛龙"
	var b = "烛龙"
	fmt.Println(strings.EqualFold(a, b))                              //判断两个字符串是否相等，不过这里特定了是utf-8编码的
	fmt.Println(strings.HasPrefix(a, "烛"), strings.HasSuffix(a, "龙")) //判定前后缀
	fmt.Println(strings.Contains(a, "龙"))                             //判断字符串中是否包含子串
	var c rune = 32
	fmt.Println(strings.ContainsRune(a, c)) //containsRune()：看字符串是否包含utf-8值
	var d = "你是一头猪吗一头猪?"
	fmt.Println(strings.ContainsAny(d, "一头猪")) //看字符串中是否包含另一个字符串的任意一个字符
	fmt.Println(strings.Count(d, "一头猪"))       //看字符串中包含几个指定的重复子串
	fmt.Println(strings.Index(d, "一头猪"))       //字符串第一次出现的位置
	fmt.Println(strings.LastIndex(d, "一头猪"))   //字符串最后一次出现的位置

	var e = "the coffee"
	s := strings.Title(e) //将字符串中的每个单词的首字母都置为大写
	fmt.Println(s)
	fmt.Println(strings.ToLower(e), strings.ToUpper(e))
	fmt.Println(strings.Repeat(e, 2)) //返回指定个字符串的拼接形式
	fmt.Println(strings.Trim(e, "e")) //去掉字符串的指定utf-8码点
	fmt.Println(strings.TrimSpace(e)) //将首尾的空格去掉
	slice := strings.Fields(e)        //将字符串以空格分割，获取字符串切片
	fmt.Println(slice)

	sli2 := strings.Split(e, "c") //以去掉某个字符的形式将字符串分割为字符串切片，另外还有各种切割的方法
	fmt.Println(sli2)

	str3 := strings.Join(sli2, " ") //将字符串切片以某个符号连接
	fmt.Println(str3)

	fmt.Println(strings.NewReader(d).Len()) //创建一个Reader,然后使用结构体的方法
}
