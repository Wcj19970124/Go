package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const testPath = "./static/src/activity_listview.vue"

func main() {

	file, err := os.Open(testPath)
	if err != nil {
		fmt.Println("文件打开失败:" + err.Error())
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("文件读取失败:" + err.Error())
	}

	reg := regexp.MustCompile("require\\('.*'\\)")
	if reg.MatchString(string(data)) {
		fmt.Println("匹配成功")
	} else {
		fmt.Println("匹配不成功!")
	}

	slice := reg.FindAllString(string(data), -1)

	res := []string{}
	for _, v := range slice {
		begin := strings.Index(v, "'")
		end := strings.LastIndex(v, "'")

		res = append(res, v[begin+1:end])
	}

	fmt.Println(filepath.Base(res[1]))
}
