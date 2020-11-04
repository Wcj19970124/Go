package main

import (
	"bufio"
	"fmt"
	"os"
)

//谨记：一个汉字是3个字节
//bufio：实际上就是缓冲式io，对io.Reader和io.Writer进行了包装
//缓冲式io.Reader:将数据包装到缓冲中，然后可以将缓冲中的数据写入到文件中...
//缓冲式io.Writer：不需要将数据包装到缓冲中，直接将数据写入到缓冲中，然后可以直接flush刷新到文件中
func main() {

	// var str string = "\n一天一天，一年一年"
	// reader := strings.NewReader(str)
	// b1 := bufio.NewReader(reader) //包装一个io.Reader对象为一个自带默认缓冲大小的reader
	// b2 := bufio.NewReaderSize(reader, 10)     //包装一个io.Reader对象为一个指定缓冲大小的readers
	// fmt.Println(b1.Buffered(), b2.Buffered()) //读取缓冲中现有的字节数
	// data, err := b1.Peek(28)                  //读取reader中一定的字节数，返回一个byte切片，如果指定字节数大于缓冲中的字节数，返回错误EOF
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(string(data))
	// }
	// date := make([]byte, len(str))
	// b1.Read(date) //读取缓冲中的数据放入到指定的byte切片中
	// fmt.Println(string(date))
	// // b1.UnreadByte() //吐出最后一个读取的字节数
	// // b1.Read(date)
	// // fmt.Println(string(date))
	// r, i, err := b1.ReadRune() //读取一个utf-8编码的unicode,非utf-8编码返回EOF
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(r, i)

	// line, n, err := b1.ReadLine() //读取一行，但是是一个比较低水平的读取原语，基本不会用，一般会被ReadBytes()和ReadString()替代
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(string(line), n)
	// }

	// date, _ = b1.ReadBytes('\n') //读取在遇到指定字节的数据，返回一个包含已读取到的数据和指定字节的字节切片，ReadString()也是如此
	// fmt.Println(string(date))
	file, err := os.Create("bufio.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	// b1.WriteTo(file) //将缓冲中的数据写入到io.Writer中,这里是将缓冲中的数据写入到*os.File，因为其实现了io.Reader和io.Writer
	writer := bufio.NewWriter(file) //将io.Writer对象包装成自带默认缓冲大小的io.Writer
	fmt.Println(writer.Buffered())  //返回缓冲中已经使用的字节数
	fmt.Println(writer.Available()) //返回缓冲中未使用的字节数，即默认缓冲区大小为4096bytes
	date := []byte("hello world")
	writer.Write(date) //将指定byte切片中的内容写入缓冲区，剩余的还有一些WriteByte(),WriteBytes(),WriteString()等等
	writer.Flush()     //将缓冲中的数据刷新进入下层io.Writer,例如可以将缓冲的数据刷新进入到文件中
}
