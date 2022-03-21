package main
import "fmt"

type Two_Array struct {
	Two_arr [3][3]int
} 

func main(){
	var arr Two_Array
	arr.Two_arr[0][0] =1
	arr.Two_arr[1][0] =3
	arr.Two_arr[2][0] =5
	arr.Two_arr[0][2] =7
	arr.Two_arr[0][1] =4

	for _,v := range arr.Two_arr {
		fmt.Println(v)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	n := 0
	for j := 0; j < 3; j++ {
		n = arr.Two_arr[0][j]
		arr.Two_arr[0][j] = arr.Two_arr[j][0]
		arr.Two_arr[j][0] = n
	}

	for _,v := range arr.Two_arr {
		fmt.Println(v)
	}

}