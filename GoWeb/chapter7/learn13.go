package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

//Person 主
type Person struct {
	XMLName xml.Name `xml:"person"`
	ID      string   `xml:"id,attr"`
	Name    string   `xml:"name"`
	Age     string   `xml:"age"`
	Pets    []Pet    `xml:"pets"`
}

//Pet 从
type Pet struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name"`
	Age  string `xml:"age"`
}

func learn13() {
	// marshalXML()
	encodeXML()
}

//marshal方式生成xml
func marshalXML() {
	person := Person{
		ID:   "1",
		Name: "wcj",
		Age:  "24",
		Pets: []Pet{
			{
				ID:   "1",
				Name: "cat",
				Age:  "3",
			},
			{
				ID:   "2",
				Name: "dog",
				Age:  "4",
			},
		},
	}

	data, err := xml.MarshalIndent(&person, "", "\t")
	if err != nil {
		log.Println(err.Error())
	}

	err = ioutil.WriteFile("person.xml", data, 0644)
	if err != nil {
		log.Println(err.Error())
	}
}

//Encode方式生成xml
func encodeXML() {

	person := Person{
		ID:   "1",
		Name: "wcj",
		Age:  "24",
		Pets: []Pet{
			{
				ID:   "1",
				Name: "cat",
				Age:  "3",
			},
			{
				ID:   "2",
				Name: "dog",
				Age:  "4",
			},
		},
	}

	file, err := os.Create("person.xml")
	if err != nil {
		log.Println(err.Error())
	}

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "\t")
	err = encoder.Encode(&person)
	if err != nil {
		log.Println(err.Error())
	}
}
