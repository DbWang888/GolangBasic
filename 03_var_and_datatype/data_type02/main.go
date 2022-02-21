package main
import (
	"fmt"
	"strconv"

)

//基本数据类型的转换
func main(){
	//被转换的时 变量存储的数据（值），变量本身的数据类型没有变化
	var i int = 100
	var i1 float64 = float64(i)
	fmt.Printf("i1 = %v i = %T i1 = %T\n", i1, i, i1)

	//大范围类型转到小范围类型，编译不报错，但是结果按溢出处理
	//可能和希望的结果不一样

	var i2 int64 = 64440
	var i3 int8 = int8(i2)
	fmt.Printf("i2 = %d i3 = %d\n", i2, i3)

	//基本数据类型转string
	
	var num1 int = 99
	var num2 float64 =23.456
	var b bool = true
	// var mychar byte = 'h'
	var str string

	//方式  fmt.Sprintf()
	str = fmt.Sprintf("%d" ,num1)
	fmt.Printf("%T %q\n",str ,str)
	

	//使用strconv包函数
	str = strconv.FormatInt(int64(num1),10)
	fmt.Printf("%T %q\n",str ,str)

	str = strconv.FormatFloat(num2, 'f', 10, 64) //f表示格式，10表示小数点后保留10位，64表示float64
	fmt.Printf("%T %q\n",str ,str)

	str = strconv.FormatBool(b)
	fmt.Printf("%T %q\n",str ,str)

	str = strconv.Itoa(num1)  //只能转int类型，int64及其他类型需要转为int
	fmt.Printf("%T %q\n",str ,str)


	//string转基本数据类型
	//strconv.ParseXX相关函数
	// 使用 _ 来丢弃不需要的返回值
	var str2 string = "123456"
	var n1 int64
	var str3 string = "true"
	var n2 bool
	var str4 string = "456.789"
	var n3 float64
	
	n1, _ =strconv.ParseInt(str2, 10, 64)
	fmt.Printf("%T %v\n",n1 ,n1)

	n2, _ = strconv.ParseBool(str3)
	fmt.Printf("%T %v\n",n2 ,n2)

	n3, _ = strconv.ParseFloat(str4, 64)
	fmt.Printf("%T %v\n",n3 ,n3)

	//注意：如果是无法转换的值，则返回0  如将hello转换成int

	var str5 string = "hello"
	n1, _ =strconv.ParseInt(str5, 10, 64)
	fmt.Printf("%T %v\n",n1 ,n1)


}