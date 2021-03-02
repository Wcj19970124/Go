package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

const testPath = "./static/src/dynamic_detail.vue"

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

	reg := regexp.MustCompile("(?:require\\(('|\"))(.*)(?:('|\")\\))")
	if reg.MatchString(string(data)) {
		fmt.Println("匹配成功")
	} else {
		fmt.Println("匹配不成功!")
	}

	slice := reg.FindAllString(string(data), -1)
	fmt.Println(slice)
}
