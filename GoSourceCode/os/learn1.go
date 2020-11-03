package main

import (
	"fmt"
	"os"
)

//os包主要提供的就是对于底层操作系统的操作的封装
//1，获取底层os的环境等各种信息
//2,FileMode对象，表示文件的模式（目录或者文件）和权限
//3,FileInfo对象，表示文件对象，可以进行基本的例如权限修改，信息获取的操作，但是不能进行读写
//4,File对象，表示已打开的文件，可以使用其进行io操作
func learn1() {

	hostName, _ := os.Hostname() //获取系统提供的主机名
	fmt.Println(hostName)
	pageSize := os.Getpagesize()
	fmt.Println(pageSize) //获取底层系统内存页的尺寸
	// env := os.Environ()   //获取系统的环境变量参数，返回map
	// fmt.Println(env)
	goRoot := os.Getenv("GOROOT") //获取指定环境变量的值
	fmt.Println(goRoot)
	// os.Setenv("GOROOT", "D:/Go") //设置环境变量
	// os.Clearenv()                //清空环境变量
	// os.Exit(-1)                                                                                                  //系统退出，0表示成功，其他表示失败
	fmt.Println(os.Getuid(), os.Geteuid(), os.Getgid(), os.Getegid(), os.Getpid(), os.Getppid()) //获取用户id，组id，进程id

	//FileMode对象，用于表示文件的模式（目录/文件）和权限位
	var f os.FileMode = 24
	fmt.Println(f.IsDir(), f.IsRegular(), f.Perm(), f.String()) //检查FileMode是否为一个目录/普通文件，输出权限位，字符串形式输出

	//FileInfo对象，用于表示一个文件对象
	// fileInfo, _ := os.Stat("active.txt") //返回一个FileInfo对象，不存在则返回PathError,存在则跳转该链接,和Lstat有一定区别
	// fmt.Println(fileInfo.Name(), fileInfo.Size(), fileInfo.Mode(), fileInfo.IsDir(), fileInfo.ModTime(), fileInfo.Sys())
	fmt.Println(os.IsPathSeparator('/')) //判断指定字符是否是一个路径分隔符
	//判断错误类型是否为目录或者文件不存在/权限不够
	fmt.Println(os.IsExist(os.ErrExist), os.IsNotExist(os.ErrNotExist), os.IsPermission(os.ErrPermission))
	fmt.Println(os.Getwd()) //获取当前工作目录
	// os.Chdir() //修改当前目录为dir
	// os.Chmod()  //修改文件权限
	// os.Chown()  //修改文件拥有者
	// os.Chtimes()  //修改文件时间信息
	// os.Mkdir()     //创建目录按照指定权限
	// os.MkdirAll()  //递归创建目录
	// os.Remove()    //删除目录
	// os.RemoveAll() ///递归删除目录

	//File对象，一个已打开的文件，可以直接用于io操作,比较常用的几个方法是Open(),Create(),Write(),WriteString(),Close()
	//其他的还有一些涉及基本操作的，例如权限，拥有者等等，不再阐述
	file, _ := os.Create("active.txt") //创建一个可供操作的文件
	fmt.Println(file)
	// activeFile, _ := os.Open("active.txt") //打开的文件只能读不能写
	// defer activeFile.Close()
	ret, err := file.WriteString("你是一头猪吗?")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(ret)
	}

}
