package main
import "fmt"

func changeArr01(arr [3]int)(){
	arr[0] = 10
}

func changeArr02(arr *[3]int)(){
	arr[2] = 10
	(*arr)[1] = 90
}

func main(){
	var arr01 [3]int = [3]int{0,1,2}
	// for i := 0; i < len(arr01); i++ {
	// 	fmt.Printf("%v\n", arr01[i])
	// } 
	// var arr02 = [2]int{11,22}
	// var arr03 = [...]int{1:800, 2:200}
	// var arr04 = [...]int{122,33,44,55}

	// for index,value := range arr02 {
	// 	fmt.Printf("index=%v,value=%v\n",index,value)
	// }
	// for index,value := range arr03 {
	// 	fmt.Printf("index=%v,value=%v\n",index,value)
	// }
	// for index,value := range arr04 {
	// 	fmt.Printf("index=%v,value=%v\n",index,value)
	// }

	changeArr02(&arr01)
	fmt.Printf("%v\n",arr01[2])
	fmt.Printf("%v\n",arr01[1])
}