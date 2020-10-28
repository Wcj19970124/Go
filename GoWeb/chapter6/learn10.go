package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/Go-SQL-Driver/MySQL"
)

//Article 文章结构体
type Article struct {
	ID      int
	Content string
	Author  string
}

var db *sql.DB //数据库连接池

func init() {
	var err error
	db, err = sql.Open("mysql", "root:520wcj@/goweb")
	if err != nil {
		panic(err)
	}
}

//Insert 插入单条文章信息进入数据库
func (article *Article) Insert() error {

	statement := "insert into article(content,author) values(?,?)"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(article.Content, article.Author).Scan(&article.ID)
	if err != nil {
		return err
	}
	return nil
}

//GetArticleByID 根据文章id获取文章信息
func GetArticleByID(id int) (Article, error) {
	var article Article
	sql := "select id,content,author from article where id = " + strconv.Itoa(id)
	err := db.QueryRow(sql).Scan(&article.ID, &article.Content, &article.Author)
	if err != nil {
		return Article{}, err
	}
	return article, nil
}

//UpdateArticle 更新帖子信息
func (article *Article) UpdateArticle() error {
	sql := "update article set content = ?,author=? where id = ?"
	log.Println(sql)
	_, err := db.Exec(sql, article.Content, article.Author, article.ID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteArticle 删除帖子
func DeleteArticle(id int) error {
	sql := "delete from article where id = ?"
	_, err := db.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	//insert article
	article := Article{Content: "今天的天气不错啊", Author: "wcj"}
	fmt.Println(article)
	err := article.Insert()
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(article)

	//query article
	article, err = GetArticleByID(2)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(article)

	//update article
	article = Article{ID: 2, Content: "你想要加班到十点吗？", Author: "hhy"}
	err = article.UpdateArticle()
	if err != nil {
		log.Println(err.Error())
	}

	err = DeleteArticle(2)
	if err != nil {
		log.Println(err.Error())
	}
}
