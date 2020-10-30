package main

import (
	"fmt"
	"time"
)

//time下的函数和方法，比较重要的有
//Now(),Date(),Clock(),Parse(),ParseInLocation(),Unix(),UnixNano()
//Year(),Month(),Day(),Hour(),Minute(),Second(),YearDay(),WeekDay()
//Add(),AddDate(),Sub(),Format(),MarshalJSON()，UnmarshalJSON()等
func learn1() {

	now := time.Now()                                                                              //Now()函数：获取当前时间
	fmt.Println(now.Date())                                                                        //time.Date() 返回时间的年月日形式  @@
	parseTime, _ := time.Parse("2006-01-02 15:04:05", "2020-10-30 12:06:45")                       //Parse()将指定的字符串形式按照Go中的标准时间 @@
	fmt.Println(parseTime)                                                                         //2006-01-02 15:04:05格式化成一个time类型的对象
	todayLast, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-10-30 23:59:59", time.Local) //@@
	fmt.Println(todayLast)                                                                         //ParseInLocation()按照当地时区格式化字符串为一个Time对象
	unixTime := time.Unix(999999, 99999)                                                           //Unix()函数：返回一个距离1970年1月之后的xxx秒xxx纳秒的时间对象Time
	fmt.Println(unixTime)
	fmt.Println(now.Location()) //Location()返回当前时间的时区信息
	fmt.Println(now.Zone())     //Zone()返回时区的规范名
	fmt.Println(now.IsZero())   //检查时间是否是零点
	fmt.Println(now.Local(), now.UTC())
	fmt.Println(now.Unix())                                                                                               //时间从1970,1的秒数
	fmt.Println(now.UnixNano())                                                                                           //时间从1970,1的纳秒数
	fmt.Println(now.Clock())                                                                                              //返回时间的时分秒
	fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.YearDay(), now.Weekday()) //年月日时分秒，当年的第几天
	fmt.Println(now.Add(1 * time.Hour))                                                                                   //Add()时间的基础上增加时间，返回一个Time对象
	fmt.Println(now.AddDate(1, 1, 1))                                                                                     //具体增加时间，年月日
	fmt.Println(now.Sub(todayLast))                                                                                       //求两个Time的差值
	fmt.Println(now.Format("20060102"))                                                                                   //Format()将时间格式化成指定的形式
	data, _ := now.MarshalJSON()                                                                                          //将时间序列化为JSON格式
	fmt.Println(string(data))
}
