package main
import "fmt"

func main(){
	var arr [5]int = [5]int{11,2,22,4,5}
	slice1 := arr[0:2]
	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}

	var slice2 []int = make([]int, 5, 10)
	slice2[1] = 111
	slice2[4] = 222
	
	for i, v := range slice2 {
		fmt.Printf("i = %v , v = %v\n", i, v)
	}

	var slice3 []string = []string{"wo","de","xin"}
	for _, v := range slice3 {
		fmt.Printf("v = %v\n", v)
	}
}