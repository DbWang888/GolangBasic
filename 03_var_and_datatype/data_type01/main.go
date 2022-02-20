package main
import (
	"fmt"
	"unsafe"
)
func main(){
	
	//fmt.Printf() 用于格式化输出
	// %T 表示输出变量的数据类型,  i:=100,默认i为int
	i := 100
	fmt.Printf("i的变量类型为：%T\n",i)

	//查看变量的字节大小以及数据类型
	var n int64 = 1000
	fmt.Printf("n的字节大小为%d\n数据类型为%T\n", unsafe.Sizeof(n), n)

	//Golang程序中整型变量在使用时，遵守保小不保大原则，
	//程序正确运行的情况下，根据实际情况，尽量使用占空间小的数据类型
	//例：人的年纪，用byte就足够
	var age byte = 30
	fmt.Printf("人的年纪：%d\n", age)

	//浮点型的默认声明为float64
	m := 5.12
	fmt.Printf("m的类型为：%T\n", m)


	//字符类型的使用 golang中无字符类型变量，常用byte保存
	//只能用单引号''表示，单字符不能用双引号 "
	var char1 byte = '1' //字符的1
	var char2 byte = 'a'
	//直接输出，则输出的内容为字符所对应的ASCII码值
	fmt.Println("char1=", char1, "char2=", char2) 
	//需要使用格式化输出来输出赋给变量的字符
	fmt.Printf("char1=%c char2=%c\n", char1, char2)
	//若赋值的字符码值大于byte的范围，可用int代替
	var char3 int = '北'
	fmt.Printf("char3=%c char3对应的码值%d\n", char3, char3)

	//字符类型可以运算，本质是对应的码值的运算
	var char4 = 10 + 'a' //a的码值为97  相当于 10+97 = 107
	fmt.Println("char4=", char4)

	//字符串一旦赋值，赋值后字符串内容不可更改
	var str = "hello"
	//str[0] = 'a' 会报错，不可更改
	//使用反引号 ``可输出带特殊符号的文本 比如：代码
	var str2 = `
	var char3 int = '北'
	fmt.Printf("char3=%c char3对应的码值%d\n", char3, char3)
	`
	fmt.Println(str2,str)

	//字符串拼接操作 太长时 可换行  “+”需要放在上一行行末
	str3 := "123123" + "1123123" +
	"456"+"456"
	fmt.Println(str3)

	//基本数据类型默认值，未赋值时则为默认值，又叫零值
	var inta int  //0
	var flo1 float32   //0.00
	var flo2 float64  //0.00
	var bool1 bool  //false
	var strs string  // ""
	fmt.Printf("%d\n%f\n%f\n%v\n%v",inta,flo1,flo2,bool1,strs)  //%v表示按照变量原值输出

}