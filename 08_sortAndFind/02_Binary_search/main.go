package main
import (
	"fmt"
	"time"
)


func BinarySearch(arr [7]int, leftIndex int, rightIndex int, findVal int)(){
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	middle := (leftIndex + rightIndex) / 2
	fmt.Println("middle = ", middle)
	fmt.Println("leftIndex = ", leftIndex)
	fmt.Println("rightIndex = ", rightIndex)
	
	if (*arr)[middle] > findVal {
		BinarySearch(arr, leftIndex, middle - 1, findVal)
	} else if arr[middle] < findVal {
		BinarySearch(arr, middle + 1, rightIndex, findVal)
	} else {
		fmt.Println("找到了", arr[middle])
	}

}


func main(){
	start := time.Now().UnixNano()
	var arr [7]int = [7]int{2,4,5,6,7,10,111}
	BinarySearch(arr, 0, len(arr), 10)
	end := time.Now().UnixNano()
	fmt.Println(end -start)
}