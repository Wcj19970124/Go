package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type animal struct {
	ID   int
	Name string
	Age  int
}

//使用csv格式的文件进行数据的存储和读取
func learn8() {

	//第一步：创建一个csv格式的文件
	csvFile, err := os.Create("animal.csv")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer csvFile.Close()

	animals := []animal{
		{ID: 1, Name: "cat", Age: 8},
		{ID: 2, Name: "dog", Age: 20},
		{ID: 3, Name: "bird", Age: 5},
	}
	//第二步：创建一个写入器，对文件进行写入
	writer := csv.NewWriter(csvFile)
	for _, animal := range animals {
		line := []string{strconv.Itoa(animal.ID), animal.Name, strconv.Itoa(animal.Age)}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}

	writer.Flush() //将缓冲区中的内容刷新到csv文件中

	//第三步：先打开csv文件
	file, err := os.Open("animal.csv")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	//第四步：创建读取器，对文件进行读取
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 //文件中缺失某个字段中，不会造成程序崩溃
	record, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
	}

	var animalList []animal
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		age, _ := strconv.ParseInt(item[2], 0, 0)
		animal := animal{int(id), item[1], int(age)}
		animalList = append(animalList, animal)
	}

	fmt.Println(animalList)

}
