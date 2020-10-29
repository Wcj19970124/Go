package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Hero 英雄数据体
type Hero struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Creator  Creator   `json:"creator"`
	Comments []Comment `json:"comments"`
}

//Creator 创造者结构体
type Creator struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Comment 评论结构体
type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	User    string `json:"user"`
}

//解码
func unmarshal(filename string) (*Hero, error) {

	var hero Hero
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Println("open file failed,err:" + err.Error())
		return &hero, err
	}

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println("read file failed,err:" + err.Error())
		return &hero, err
	}

	err = json.Unmarshal(data, &hero)
	if err != nil {
		log.Println("unmarshal json failed,err:" + err.Error())
		return &hero, err
	}

	return &hero, nil
}

//解码
func decode(filename string) (*Hero, error) {
	var hero Hero
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Println("open file failed,err:" + err.Error())
		return &hero, err
	}

	err = json.NewDecoder(jsonFile).Decode(&hero)
	if err != nil {
		log.Println("decode json failed,err:" + err.Error())
		return &hero, err
	}

	return &hero, nil

}

func main() {
	hero, err := decode("hero.json")
	if err != nil {
		log.Println("decode json failed,err:" + err.Error())
		return
	}
	fmt.Println(hero)
}
