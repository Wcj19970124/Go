package main

//ioutil包：比较重要的方法，大概就是下面这些
//比较重要的就是io.Reader和io.Writer:这两个接口是整个io包的祖先
func main() {
	// ioutil.ReadAll()   //从io.Reader中读取信息直到结尾,返回byte的切片和err
	// ioutil.ReadFile()  //读取指定文件名的文件中的信息，返回byte切片和err
	// ioutil.ReadDir()   //读取目录中的信息，返回目录下的文件切片和err
	// ioutil.TempDir()   //创建临时目录,返回目录路径
	// ioutil.TempFile()  //创建临时文件，返回文件指针
	// ioutil.WriteFile() //将指定的文件中写入信息，如果文件不存在，按照指定权限进行创建
}
