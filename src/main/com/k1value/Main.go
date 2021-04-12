package main

import "fmt"

/*
	变量
	变量使用的注意事项
1、变量表示内存中的一个存储区域
2、该区域有自己的名称（变量名）和类型（数据类型）
3、Golang变量使用的三种方式
  一：指定变量类型，声明后不赋值，使用默认值
  二：根据值自行判定变量类型（类型推导）
  三：省略var（ := 左侧的变量不应该是已经声明过的，否则编译错误）
4、多变量声明：用 , 可声明多个变量
5、该区域的数据值可以在同一类型范围内不断变化
6、变量在同一个作用域内不能同名
7、变量=变量名+值+数据类型
8、变量如果没赋初值，则使用默认值
*/
func main() {
	// 一：指定变量类型，声明后不赋值，使用默认值
	var i int
	i = 10
	fmt.Print(i)
	// 二：根据值自行判定变量类型（类型推导）
	var a = 20
	fmt.Print("\n", a)
	// 三：省略var（ := 左侧的变量不应该是已经声明过的，否则编译错误）
	b := 30
	fmt.Print("\n", b)

	var n1, n2, n3 int = 11, 12, 13
	fmt.Print("\n", "n1 = ", n1, "\t n2 = ", n2, "\t n3 = ", n3)
	n4, n5, n6 := "hello", 100, 'a'
	fmt.Print("\n", "n4 = ", n4, "\t n5 = ", n5, "\t n6 = ", n6)

	fmt.Print("\n", "n7 = ", n7, "\t n8 = ", n8, "\t n9 = ", n9)
	fmt.Println("\n", "n10 = ", n10, "\t n11 = ", n11, "\t n12 = ", n12)
	sum()
}

var n7 = 11
var n8 = "world"
var n9 = 'b'

var (
	n10 = 12
	n11 = "!!!"
	n12 = 'c'
)

/**
加号的使用
*/
func sum() {
	// 当左右两边都是数值型时，则做加法运算
	fmt.Println(12 + 11)
	// 当左右两边都是字符串，则做字符串拼接
	fmt.Println("hello " + "world")
}