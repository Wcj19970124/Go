package main

import (
	"fmt"
	"os/exec"

	"./util"
)

const (
	// weexURL       = "git@git.code.oa.com:IED/daojuWeex.git" //工蜂Weex项目地址
	weexURL       = "git@git.code.oa.com:chengjunwu/Test.git" //测试URL
	localWeexPath = "daojuWeex"                               //Weex项目本地存放地址
)

//GitWeex 拉取工蜂Weex项目
func GitWeex() error {
	if !util.IsFileExists(localWeexPath) { //first git clone
		cmd := exec.Command("git", "clone", weexURL, localWeexPath)
		if err := cmd.Run(); err != nil { //TODD:使用阻塞命令，保证项目拉取完毕之后进行解析
			return err
		}
		return nil
	}
	// cmd := exec.Command("cd", localWeexPath)
	// if err := cmd.Run(); err != nil {
	// 	fmt.Println("切换执行发生错误:" + err.Error())
	// }
	cmd := exec.Command("git", "-C", "./daojuWeex", "pull")
	if err := cmd.Run(); err != nil {
		fmt.Println("发生了错误:" + err.Error())
		return err
	}
	return nil
}
