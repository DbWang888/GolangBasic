package main
import (
	"fmt"
	"math"
)

//得到数组中最大值及最小值对应的index
func getIndex(arr [8]float64) ( int, int) {
	max_array := arr
	min_array := arr
	var max,min float64
	var maxIndex, minIndex int

	for i := 0; i < 8; i++{
		if max > max_array[i] {
			continue
		} else {
			max = max_array[i]
			maxIndex = i
		}
	}
	min = min_array[0]
	for i := 0; i < 8; i++{
		if min < min_array[i] {
			continue
		} else {
			min = min_array[i]
			minIndex = i
		}
	}
	return maxIndex, minIndex

}
//得到数组的数的平均值
func getAverage (arr [8]float64) float64 {
	var sum float64 = 0
	for i := 0; i < 8; i++ {
		if i == 0 || i ==4 {
			continue
		}
		sum = sum + arr[i]
	}
	return sum / 6.0
}

//得到一个数组所有数减去同一个数后得到的新的数组
func getDiffArray(arr [8]float64, avg float64 ) [8]float64 {
	var diffarray [8]float64
	for i := 0; i < 8; i++ {
		diffarray[i] = math.Abs( arr[i] - avg )
	}
	return diffarray
}

func main(){

	var count_array [8]float64 = [8]float64{99.6,87.5,89.1,95.5,78.6,88.8,95.7,89.7}
	maxIndex, minIndex := getIndex( count_array )
	fmt.Printf("最高分的评委是第%d号，打最低分的评委是第%d号\n", maxIndex, minIndex)
	avg := getAverage(count_array)
	fmt.Printf("最终得分为%f\n", avg)
	diffarray := getDiffArray(count_array, avg)
	// fmt.Printf("差值数组为%v\n", diffarray)
	best, wores := getIndex(diffarray)
	fmt.Printf("最hao的评委是第%d号，zuichade的评委是第%d号\n", best, wores)
}