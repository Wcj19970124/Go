package model

import (
	"database/sql"
	"log"

	_ "github.com/Go-SQL-Driver/MySQL"
)

//DB 数据源
var DB *sql.DB

//Student 学生结构体
type Student struct {
	ID      int
	Name    string
	Age     int
	Gender  string
	Address string
}

//初始化数据源
func init() {
	var err error
	DB, err = sql.Open("mysql", "root:520wcj@/goweb")
	if err != nil {
		log.Println(err.Error())
		return
	}
}

//GetStudentByID  根据学生id获取学生信息
func GetStudentByID(id int) (Student, error) {

	var student Student

	sql := "select id,name,age,gender,address from student where id = ?"
	err := DB.QueryRow(sql, id).Scan(&student.ID, &student.Name, &student.Age, &student.Gender, &student.Address)
	if err != nil {
		log.Println("query student by id failed,err:" + err.Error())
		return student, err
	}

	return student, nil
}

//InsertStudent 添加学生记录
func InsertStudent(s Student) (int, error) {

	sql := "insert into student(name,age,gender,address) values(?,?,?,?)"
	_, err := DB.Exec(sql, s.Name, s.Age, s.Gender, s.Address)
	if err != nil {
		log.Println("insert student info failed,err:" + err.Error())
		return 0, err
	}

	return 0, nil
}

//UpdateStudent 修改学生信息
func UpdateStudent(s Student) (int, error) {

	sql := "update student set name = ?,age = ?,gender = ?,address = ? where id = ?"
	_, err := DB.Exec(sql, s.Name, s.Age, s.Gender, s.Address, s.ID)
	if err != nil {
		log.Println("update student info failed,err:" + err.Error())
		return 0, err
	}

	return 0, nil
}

//DeleteStudentByID 根据学生id删除学生信息
func DeleteStudentByID(id int) (int, error) {

	sql := "delete from student where id = ?"
	_, err := DB.Exec(sql, id)
	if err != nil {
		log.Println("delete student info by id failed,err:" + err.Error())
		return 0, err
	}

	return 0, nil
}
