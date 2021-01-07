package main

import (
	"fmt"
	"os/exec"
)

type person struct {
	Name string
	Age  int
}

func main() {

	// cmd := "git"                                      //命令
	// param := "clone"                                  //命令参数
	// proUrl := "git@git.code.oa.com:IED/daojuWeex.git" //命令参数
	// proPath := "./static"                             //文件存放路径

	// cmd := exec.Command("tr", "a-z", "A-Z")
	// cmd.Stdin = strings.NewReader("some input")
	// var out bytes.Buffer
	// cmd.Stdout = &out
	// err := cmd.Run()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(out.String())

	// cmd := exec.Command("echo", `{"Name":"wcj","Age":24}`)
	// stdout, err := cmd.StdoutPipe()
	// if err != nil {
	// 	fmt.Println("发生了错误:" + err.Error())
	// }
	// if err := cmd.Start(); err != nil {
	// 	fmt.Println("命令执行失败:" + err.Error())
	// }

	// var p person
	// if err := json.NewDecoder(stdout).Decode(&p); err != nil {
	// 	fmt.Println("解码失败:" + err.Error())
	// }

	// if err := cmd.Wait(); err != nil {
	// 	fmt.Println("命令结束关闭管道失败:" + err.Error())
	// }

	// fmt.Printf("%s is %d year", p.Name, p.Age)

	// data, err := exec.Command("data").Output()
	// if err != nil {
	// 	fmt.Println("执行失败:" + err.Error())
	// } else {
	// 	fmt.Println(string(data))
	// }

	cmd := exec.Command("git", "clone", "git@git.code.oa.com:IED/daojuWeex.git", "./static")
	_, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("start 执行，获取标准输出管道失败:" + err.Error())
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("命令执行失败：" + err.Error())
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("命令执行完毕之后，关闭管道失败:" + err.Error())
	}
	fmt.Println("执行命令成功!")
}
