package main
import "fmt"

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
	fmt.Printf("i2 = %d i3 = %d", i2, i3)
}