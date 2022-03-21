package main
import (
	"fmt"
	"math/rand"
	"time"
)

func getArr()( [10]int ){
	var arr [10]int
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(101)
	}
	return arr
}

//二分法查找
// func binary_search(arr *[10]int, leftindex, rightindex, findval ){
// 	var middleVal int
// 	var middleVal int = arr[(leftindex + rightindex)/2]
// 	for {
// 		middleVal := arr[middleIndex]
// 		if findval > middleVal {
// 			leftindex = middleIndex
// 			middleIndex = ( leftindex + rightindex ) / 2
// 			middleVal = arr[middleIndex]
// 		} else if findval == middleVal {
// 			fmt.Println("找到了",middleIndex)
// 			return middleIndex
// 		} else if findval < middleVal {
// 			rightindex = middleIndex
// 			middleIndex = ( leftindex + rightindex ) / 2
// 			middleVal = arr[middleIndex]
// 		} else if leftindex > rightindex {
// 			fmt.Println("没有找到")
// 			return -1
// 		}
// 	}
	
// }
 
func main(){
	arr := getArr()
	temp := 0
	for i := 0; i < (len(arr) - 1); i++{
		for i := 0; i < len(arr) - 1; i++ {
			if arr[i] > arr[ i+1 ]{
				temp = arr[i]
				arr[i] = arr[ i+1 ]
				arr[i + 1] = temp
			}
		}
	}
	fmt.Println(arr)
	var leftindex,rightindex,findval,middleIndex int
	findval = 22
	leftindex = 0
	rightindex = 9
	middleIndex = 4
	var middleVal int = arr[4]
	for {
		if leftindex > rightindex {
			fmt.Println("没有找到")
			break
		}
		middleVal := arr[middleIndex]
		fmt.Println("md")
		if findval > middleVal {
			leftindex = middleIndex + 1
			middleIndex = ( leftindex + rightindex ) / 2
			middleVal = arr[middleIndex]
			fmt.Println("11111111")
		} else if findval == middleVal {
			fmt.Println("找到了",middleIndex,findval)
			break
		} else if findval < middleVal {
			rightindex = middleIndex - 1
			middleIndex = ( leftindex + rightindex ) / 2
			middleVal = arr[middleIndex]
			fmt.Println("2222222")
		} 
	}

	fmt.Println(middleVal)


	

}	