package main
import "fmt"

func main(){

	var two_arr [4][6]int
	two_arr[1][2] = 5
	two_arr[2][1] = 6
	for i := 0; i < len(two_arr); i++ {
		for j := 0; j < len(two_arr[i]); j++ {
			fmt.Printf("%v", two_arr[i][j])
		}
		fmt.Printf("\n")
	}

	var two_arr2 [2][2]int = [2][2]int{{2,2},{1,2}}
	fmt.Println()
	fmt.Println(two_arr2)

	for _, v := range two_arr {
		for _, v2 := range v {
			fmt.Printf("%v",v2)
		}
		fmt.Printf("\n")
	}
}