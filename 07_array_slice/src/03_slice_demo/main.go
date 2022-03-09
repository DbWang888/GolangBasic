package main
import "fmt"

func main(){
	var arr [5]int = [5]int{11,2,22,4,5}
	slice1 := arr[0:3]
	

	sliceslice := slice1[0:2]
	fmt.Println(sliceslice)
	ptr1 := &arr[0]
	ptr2 := &slice1[0]
	ptr3 := &sliceslice[0]
	sliceslice[0] = 10000000
	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}
	fmt.Printf("%v\n%v\n%v", ptr1, ptr2, ptr3)
	// var slice2 []int = make([]int, 5, 10)
	// slice2[1] = 111
	// slice2[4] = 222
	// for i, v := range slice2 {
	// 	fmt.Printf("i = %v , v = %v\n", i, v)
	// }

	// var slice3 []string = []string{"wo","de","xin"}
	// for _, v := range slice3 {
	// 	fmt.Printf("v = %v\n", v)
	// }
	// fmt.Printf("cap= %v\n", cap(slice3))
	// fmt.Printf("len= %v\n", len(slice3))
}