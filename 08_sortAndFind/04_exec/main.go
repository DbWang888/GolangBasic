package main
import (
	"fmt"
	"math/rand"
	"time"
)



// 得到一个由0-200之间的随机数构成的长度为10的数组
func getArray()( [10]int ){
	rand.Seed(time.Now().Unix())
	var array [10]int
	for i := 0; i < 10; i++ {
		array[i] = rand.Intn(200)
	}
	return array
}

// //对数组进行倒序排序
func getReverseorder(arr [10]int) ([10]int){
	temp := 0
	for i := 0; i < 9; i++{
		for i := 0; i < 8; i++{
			if arr[i] < arr[i+1] {
				temp = arr[i]
				arr[i] = arr[ i + 1 ]
				arr[i + 1] = temp
		}	
	}
	
}   
return arr
} 
 
//返回最大值所在的下标和返回最小值所在的下标
//4 9 6 8 7
//46987
//0 1
func getMaxVal( arr [10]int) (int, int) {
	max := arr[0]
	maxIndex := 0
	for i := 0; i < 9; i++ {
		if max > arr[i] {
			continue
		} else {
			max = arr[i]
			maxIndex = i
		}
	}
	return max, maxIndex
}

//得到最小值以及对应的下标
func getMinVal( arr [10]int) (int, int) {
	min := arr[0]
	minIndex := 0
	for i := 0; i < 9; i++ {
		if min < arr[i] {
			continue
		} else {
			min = arr[i]
			minIndex = i
		}
	}
	return min, minIndex
}

//倒序打印
func reversePrint(arr [10]int){
	for i := (len(arr) - 1); i >= 0; i-- {
		fmt.Println(arr[i])
	}
}

func main(){
	arr := getArray()
	reversePrint(arr)
	fmt.Println(arr)
	fmt.Println(getReverseorder(arr))
	maxValue,maxIndex := getMaxVal(arr) 
	minValue,minIndex := getMinVal(arr)
	fmt.Println(maxValue,maxIndex)
	fmt.Println(minValue,minIndex)
}

