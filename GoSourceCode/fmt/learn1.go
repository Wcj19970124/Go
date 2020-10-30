package main

import (
	"fmt"
	"strings"
)

//fmt包下的scan()函数：
//Scan(),Scanln(),Scanf()：都是聪标准输入读取内容，并且写入到变量中
//Sscan(),Sscanln(),Sscanf():都是从指定的字符串中读取内容，写入到指定变量中
//Fscan(),Fscanln(),Fscanf():都是从io.Reader中读取内容，写入到指定变量中
//scan()这种类型：都是以空白作为分隔符，换行也默认为分隔符
//scanln()这种类型：以换行作为结束符
//scanf()这种类型：以某种格式对读取的内容进行格式化之后再写入到变量中
func learn1() {

	//fmt.Scan()函数：作用就是从标准输入依据空格(换行也算空格)分割数据，将数据写入到指定的变量中，输入的数据和参数个数一样时终止
	var a string
	var b int
	fmt.Scan(&a, &b)
	fmt.Println(a, b)

	//fmt.Scanln()函数：作用和fmt.Scan()函数一样，只不过当换行时终止
	var c string
	var d string
	fmt.Scanln(&c, &d)
	fmt.Println(c)
	fmt.Println(d)

	//fmt.Scanf()函数：作用和上面两者类似，只不过是将标准输入中的内容按照指定的参数进行格式化并且写入指定参数中
	var e bool
	fmt.Scanf("%t", &e)
	fmt.Println(e)

	//fmt.Sscan()函数：将指定的字符串写入到指定的参数变量中,以空白作为分割(换行也作为空白处理)
	var f string
	var g string
	fmt.Sscan("哈哈 呵呵", &f, &g) //如果没有空白，会将所有的信息写入到第一个参数变量中，第二个自动为""
	fmt.Println(f, g)
	fmt.Println(g)

	//fmt.Sscanln()函数：作用和Sscan()函数类似，只不过以换行作为结束符
	var h string
	var i string
	fmt.Sscanln("哈哈\n 你真是天才", &h, &i)
	fmt.Println(h, i)

	//fmt.Sscanf()函数：作用就是将字符串通过指定的format进行格式化，然后写入指定的参数变量中
	var j string
	fmt.Sscanf("猪吧你 呵呵", "%s", &j)
	fmt.Println(j)

	//fmt.Fscan()函数：从指定的io.Reader中读取文本，依据空白作为分割写入到参数变量中
	str := "一只羊两只羊\n 三只羊"
	var k string
	fmt.Fscan(strings.NewReader(str), &k)
	fmt.Println(k)

	//fmt.Fscanln()和散户：从指定的io.Reader()中读取文本，依据换行作为结束符，将读取的内容写入到变量中
	var l string
	fmt.Fscanln(strings.NewReader(str), &l)
	fmt.Println(l)

	//fmt.Fscanf()函数：从指定的io.Reader()中读取文本，依据空白作为分隔符(换行默认为空白)，将读取到的内容按照指定参数格式化写入变量
	var m string
	fmt.Fscanf(strings.NewReader(str), "%s", &m)
	fmt.Println(m)
}
