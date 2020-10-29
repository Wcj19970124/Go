package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//Hero 数据体
type Hero struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Creator  Creator `json:"creator"`
	Comments []Topic `json:"comments"`
}

//Creator 数据体
type Creator struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Topic 数据体
type Topic struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	User    string `json:"user"`
}

//总结：Marshal()和Unmarshal()方法，需要打开文件，读取其中的数据，然后进行序列化或者反序列即可
//Decode()和Encode()方法，需要打开文件之后，根据文件创建一个编码器或者解码器，然后再对文件进行解码或者编码即可
func main() {
	// unmarshalJSON()
	// decodeJSON()
	// marshalJSON()
	encodeJSON()
}

//解码JSON数据到结构体中，使用Unmarshal()方法，此方法适用于解析[]byte数据
func unmarshalJSON() {

	jsonFile, err := os.Open("article.json")
	if err != nil {
		log.Println(err.Error())
	}

	defer jsonFile.Close()
	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println(err.Error())
	}

	var hero Hero
	if json.Unmarshal(data, &hero) == nil {
		fmt.Println(hero)
	}
}

//解码JSON数据，使用decode方法,此方法适用于解析request.Body和流数据
func decodeJSON() {

	jsonFile, err := os.Open("article.json")
	if err != nil {
		log.Println(err.Error())
	}

	defer jsonFile.Close()
	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println(err.Error())
	}

	var hero Hero
	if json.NewDecoder(strings.NewReader(string(data))).Decode(&hero) == nil {
		fmt.Println(hero)
	}
}

//使用marshal()方法，编码json数据
func marshalJSON() {

	hero := Hero{
		ID:   1,
		Name: "亚索",
		Creator: Creator{
			ID:   1,
			Name: "Micale",
		},
		Comments: []Topic{
			{
				ID:      1,
				Content: "快乐风男",
				User:    "ysj",
			},
			{
				ID:      1,
				Content: "简自豪0-10搞我",
				User:    "zjq",
			},
		},
	}

	data, err := json.MarshalIndent(&hero, "", "\t")
	if err != nil {
		log.Println(err.Error())
	}

	err = ioutil.WriteFile("hero.json", data, 0644)
	if err != nil {
		log.Println(err.Error())
	}
}

//使用Encode()方法，编码json数据
func encodeJSON() {

	hero := Hero{
		ID:   1,
		Name: "亚索",
		Creator: Creator{
			ID:   1,
			Name: "Micale",
		},
		Comments: []Topic{
			{
				ID:      1,
				Content: "快乐风男",
				User:    "ysj",
			},
			{
				ID:      1,
				Content: "简自豪0-10搞我",
				User:    "zjq",
			},
		},
	}

	jsonFile, err := os.Create("hero.json")
	if err != nil {
		log.Println(err.Error())
	}

	err = json.NewEncoder(jsonFile).Encode(&hero)
	if err != nil {
		log.Println(err.Error())
	}
}
