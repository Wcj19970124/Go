package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"./util"
)

const path = "./static/src"

func test() {

	comm := util.Command{CommandName: "git", Params: []string{"clone", "git@git.code.oa.com:IED/daojuWeex.git", "./static"}}

	if err := util.ExecCommand(comm); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("拉取Weex执行成功!")
	}

	res, err := parseWeexPro(path)
	if err != nil {
		fmt.Println("解析Weex项目出错:" + err.Error())
	}

	data, err := json.Marshal(&res)
	if err != nil {
		fmt.Println("转成json发生错误:" + err.Error())
	}
	fmt.Println(string(data))

}

//parseWeexPro 解析Weex项目，获取所需文件
func parseWeexPro(rootDir string) (map[string]interface{}, error) {

	root := make(map[string]interface{})
	js := []string{}
	weex := []string{}
	component := []string{}
	filepath.Walk(rootDir, func(subPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Dir(subPath) == "static\\src" { //一级目录下Weex页面获取
			if !info.IsDir() && filepath.Ext(info.Name()) == ".vue" {
				weex = append(weex, info.Name())
			}
		} else {
			if !info.IsDir() && filepath.Ext(info.Name()) == ".vue" {
				component = append(component, info.Name())
			}
			if !info.IsDir() && filepath.Ext(info.Name()) == ".js" {
				js = append(js, info.Name())
			}
		}
		return nil
	})

	root["weex"] = weex
	root["component"] = component
	root["js"] = js
	return root, nil
}
