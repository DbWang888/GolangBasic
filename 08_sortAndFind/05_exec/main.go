package main
import (
	"fmt"
)

func main(){
	var array [10]int = [10]int{5,1,2,6,8,3,10,23,34,2121}
	array_slice := array[:]
	array_slice = append(array_slice,333)
	temp := 0

	for i := 0; i < (len(array_slice) - 1); i++ {
		for i := 0; i < (len(array_slice) - 1); i++ {
			if array_slice[i] > array_slice[ i + 1] {
				temp = array_slice[i]
				array_slice[i] = array_slice[ i + 1]
				array_slice[i + 1] = temp
			}
		} 
	}

	fmt.Println(array_slice,len(array_slice))
}