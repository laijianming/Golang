package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

/*
	变量的数据类型
基本数据类型
1、数值型
  int, int8, int16, int32, int64
  uint, uint8, uint 16, uint32, uint64, byte
  float32, float64
2、字符型
  没有专门的字符型，使用byte来保存单个字母字符
3、布尔型
  bool
4、字符串
  string （官方将string归属到基本数据类型）

派生/复杂数据类型
1、指针
  Pointer
2、数组
3、结构体
  struct
4、管道
  Channel
5、函数（也是一种类型）
6、切片
  slice
7、接口
  interface
8、map

	基本数据类型默认值
整型 0
浮点型 0
字符串 ""
布尔类型 false

	格式化输出
%d 整型数字
%T 变量类型
%v 变量值输出
%t 单词true或false
%b 二进制
%c 该值对应的Unicode码值
%d 十进制
%o 八进制
%q 该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
%x 表示为16进制 a-f  %X  A-F
%b 无小数部分、二进制指数的科学计数法，如：-123456p-78
%e 科学计数法，如 -1234.456e+78  %E
%F 有小数部分但无指数部分  %f
%s 按字符串输出

	基本数据类型转换
Go在不同类型的变量之间赋值时需要显示转换，类型不能自动转换
基本语法： T(v)  T：数据类型  v：需要转换的变量

	指针
1、基本数据类型，变量存的就是值，也叫值类型
2、获取变量的地址，用&，比如：var num int,获取num的地址：&num
3、指针类型，变量存的是一个地址，这个地址指向的空间存的才是值，如：var ptr *int = &num
4、获取指针类型所指向的值，使用：*，比如：var *ptr int，使用*ptr获取p指向的值
*/
func main() {
	//intf()
	//floatf()
	//charf()
	//boolf()
	//stringf()
	//classChange()
	pointf()
}

/**
整数类型
*/
func intf() {
	// int = int8 == 一字节
	var a int = 110000
	fmt.Println(a)

	// 无符号 不能赋值负数
	var b uint = 1
	fmt.Println(b)

	var c byte = 255
	fmt.Println(c)
	// %T 可显示变量的数据类型  printf 格式化输出
	fmt.Printf("a的数据类型 %T ", a)

	// 查看某个变量的占用字节大小和数据类型
	fmt.Printf("\na 的类型 %T  占用的字节数是 %d \n", a, unsafe.Sizeof(a))

	// 使用整型变量时，遵守保小不保大的原则：在保证程序正确运行下，尽量使用占用空间小的数据类型
	var age byte = 100
	fmt.Println(age)
}

/**
1、Golang浮点类型有固定的范围和长度，不受OS的影响
2、Golang的浮点型默认是float64类型
3、浮点型常量有两种表示形式
  十进制形式： 5.12  .512 （必须要有小数点）
  科学计数法：5.1234e2 == 5.1234 * 10的2次方   5.12E-2 == 5.12/10的2次方
4、通常情况下，应该使用float64，因为它比float32更精确
*/
func floatf() {
	// 小数类型
	var a float32 = -123.9001012
	var b float64 = -123.9001012
	fmt.Println(a, " ", b)

	c := .123
	fmt.Printf("c 的数据类型 %T \n", c)

	// 科学计数法
	d := 5.1234e2
	fmt.Println(d)
	e := 51.234e-2
	fmt.Println(e)
}

/**
	字符类型（char）
Golang中没有专门的字符类型，如果要存储单个字符（字母），一般使用byte来保存。
字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是有单个字节连接起来的。也就是说对于传统的字符串是由字符组成的。
而Go的字符串不同，它是由字节组成的
Go统一utf-8的编码
*/
func charf() {
	var a byte = 'a'
	var b = 'b'
	// 当直接输出字符时，则输出字符对应的ascll码值
	fmt.Println("a = ", a)
	// 输出对应字符，使用格式化输出
	fmt.Printf("b = %c \n", b)
	var c = '冲'
	fmt.Printf("c = %c \n", c)
}

/*
	bool

*/
func boolf() {
	// bool 逻辑运算
	var a bool
	var b bool = true
	var c = true
	fmt.Println(a, b, c)
}

/*
	string 字符串
字符串就是一串固定长度的字符连接起来的字符序列。
Go的字符串是由单个字节连接起来的。
Go语言的字符串的字节使用UTF-8编码标识Unicode文本。(不会出现中文乱码)
*/
func stringf() {

	var hello string = "hello world !!!"
	fmt.Println(hello)
	// 字符串是不可变的

	// 双引号会识别转移字符
	str0 := "123\n123"
	fmt.Println(str0)

	// 反引号的使用（纯文本输出，可以实现防止攻击、输出源代码等效果）
	str1 := `func stringf() {
	var hello string = "hello \n world !!!"
	fmt.Println(hello)
	// 字符串是不可变的
	str1 := 
}`
	fmt.Println(str1)

	// 字符串的拼接方式 ( + 必须在每行后加)
	var str2 = "hello " + "world"
	str2 += " !!!"
	fmt.Println(str2)
	var str3 = "123 " +
		"456 " +
		"789"
	fmt.Println(str3)
}

/**
变量类型转换
被转换的是变量存储的数据（即值），变量本身的数据类型并没有发生变化

f:-ddd.ddd
b:-dddp+ddd 指数为二进制
e:-d.ddde+dd  指数为十进制

将string类型转成基本数据类型时，要确保string类型能够转成有效的数据，如果把非法数据转数字，则将其转成0
*/
func classChange() {
	var a int32 = 100
	var b float64 = float64(a)
	fmt.Printf("%T\n", b)
	var c int8 = int8(b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%v %v\n", b, c)

	// 高位转低位可能溢出，按溢出结果处理
	var d int32 = 100000
	var e int8 = int8(d)
	fmt.Println(e)

	// 字符串转数字，数字转字符串
	var num1 int = 99
	var num2 float64 = 123.456
	var bo bool = true
	var co byte = 'C'
	var str string

	// 数字转字符串
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("%T \n", str)
	fmt.Println(strconv.FormatInt(int64(num1), 10))
	fmt.Println(strconv.Itoa(num1))

	str = fmt.Sprintf("%f", num2)
	fmt.Println(str)
	// f 格式  10：小数位数  64：表示是float64的
	fmt.Println(strconv.FormatFloat(num2, 'f', 10, 64))

	str = fmt.Sprintf("%t", bo)
	fmt.Println(str)

	str = fmt.Sprintf("%c", co)
	fmt.Println(str)

	// string 转 基本数据类型
	str = "true"
	bo, _ = strconv.ParseBool(str)
	fmt.Printf("bo type = %T value = %v \n", bo, bo)

	str = "123456"
	var n int64
	n, _ = strconv.ParseInt(str, 10, 64)
	fmt.Printf("n type = %T value = %v \n", n, n)

	str = "123.123"
	num2, _ = strconv.ParseFloat(str, 64)
	fmt.Printf("num2 type = %T value = %v \n", num2, num2)
}

func pointf() {
	var i int = 10
	fmt.Println("i的指针地址=", &i)

	// 下面的 var ptr *int = &i
	// 1、ptr 是一个指针变量
	// 2、ptr 的类型 *int
	// 3、ptr 本身的值&i
	var ptr *int = &i
	fmt.Println(*ptr)
}
