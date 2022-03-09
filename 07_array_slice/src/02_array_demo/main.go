package main
import (
	"fmt"
	"math/rand"
	"time"
)


func maxValue (arr [5]int)(int,int){
	max := arr[0]
	index := 0
	for i := 1; i < len(arr); i++ {
		if max > arr[i] {
			continue
		} else {
			max = arr[i]
			index = i
		}
	}
	return max, index
}

func fanzhuang(){
	rand.Seed(time.Now().UnixNano())
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = rand.Intn(500)
	}
	// fmt.Println(arr)
	for i := 5; i == 0; i-- {
		fmt.Printf("%v\n",arr[i-1])
	}
	
}

func main(){
	var char [26]byte
	char[0] = 'A'
	for i := 1; i < len(char); i++ {
		char[i] = char[i - 1] + 1
	}
	fmt.Printf("%c\n", char)

	var intArr = [5]int{1,99,111,55,11111}
	max, index := maxValue(intArr)
	fmt.Println(index)
	fmt.Println(max)
	fanzhuang()
	
}