package main

import "fmt"

//Post 帖子数据体
type Post struct {
	ID      int
	Content string
	Author  string
}

//PostByID 根据ID存储帖子内容
var PostByID map[int]*Post

//PostByAuthor 根据作者存储帖子内容
var PostByAuthor map[string][]*Post

func storeMemory(post Post) {
	PostByID[post.ID] = &post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], &post)
}

//将数据存储再内存中，应用运行时的内存中
//这种方法的优点在于：从内存中进行数据的写入读取速度极快
//缺点是：无法做到数据的持久化
func learn6() {

	PostByID = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)

	post1 := Post{ID: 1, Content: "Hello World!", Author: "Sau Sheong"}
	post2 := Post{ID: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := Post{ID: 3, Content: "Hola Mundo!", Author: "Pedro"}
	post4 := Post{ID: 4, Content: "Greetings Earthlings", Author: "Sau Sheong"}

	storeMemory(post1)
	storeMemory(post2)
	storeMemory(post3)
	storeMemory(post4)

	fmt.Println(PostByID[1])
	fmt.Println(PostByID[2])
	for _, post := range PostByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}

	for _, post := range PostByAuthor["Pedro"] {
		fmt.Println(post)
	}
}
