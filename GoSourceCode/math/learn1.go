package main

import (
	"fmt"
	"math"
)

//关于math包，其实没有什么可说的，只要你用过其他语言，那么应该都有一个类似的包
//提供一些基本的数学操作，例如最大值，最小值，绝对值，二次方，三次方等等
func learn1() {
	fmt.Println(math.NaN()) //返回NaN
	var a = 3.14
	fmt.Println(math.IsNaN(a))       //检查float类型参数是否是NaN
	fmt.Println(math.Inf(1))         //参数>=0返回正无穷大，否则返回负无穷大
	fmt.Println(math.IsInf(3.14, 1)) //用的应该比较少，不用过多考虑，需要的时候看文档
	fmt.Println(math.Ceil(3.14))     //返回不小于float类型参数的整数，例如计算分页时的总页数
	fmt.Println(math.Floor(3.14))    //返回不大于float类型参数的整数
	fmt.Println(math.Trunc(3.14))    //返回浮点数的整数部分
	n, f := math.Modf(3.14)          //返回浮点数的整数部分和小数部分
	fmt.Println(n, f)
	fmt.Println(math.Abs(-5))                   //返回绝对值
	fmt.Println(math.Max(2, 4), math.Min(2, 4)) //返回最大值，最小值
	fmt.Println(math.Mod(6.4, 3.2))             //取余
	fmt.Println(math.Sqrt(4))                   //二次方根
	fmt.Println(math.Cbrt(27))                  //三次方根
}
