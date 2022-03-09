package main
import "fmt"

func main(){
	
	var slice []int = make([]int, 5, 50)
	slice[0] = 1111
	// var slice []int = []int{100, 200,300}
	//append函数实现切片动态扩容
	fmt.Printf("len = %v\n cap= %v\n", len(slice), cap(slice))
	//动态扩容若是超出原来切片的容量 底层会生成新的数组，引用地址会改变
	//不超出原来切片的容量，则是原来数组的重新切片，引用地址不会改变
	fmt.Printf("%v\n",&slice[0])
	slice =append(slice,500,600)
	fmt.Println(slice)
	fmt.Printf("%v\n",&slice[0])
	fmt.Printf("len2 = %v\n cap2= %v\n", len(slice), cap(slice))
	//append追加切片，在切片名称后加...为固定写法

	slice = append(slice, slice...)
	fmt.Println(slice)
	fmt.Printf("len2 = %v\n cap2= %v\n", len(slice), cap(slice))
	fmt.Printf("%v\n",&slice[0])

	var slcie2 []int = []int{1,2,3,4,5,6}
	var slcie3 []int = []int{11,22,33}

	copy(slcie2,slcie3)
	fmt.Println(slcie2)
	fmt.Println(slcie3)
}