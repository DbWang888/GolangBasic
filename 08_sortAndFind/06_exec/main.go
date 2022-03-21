package main
import (
	"fmt"
)

func main(){

	var two_arr [3][4]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++{
			fmt.Printf("输入第%d行，第%d列", i, j)
			fmt.Scanln(&two_arr[i][j])
		}
	}

	

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++{
			fmt.Printf("%v\t",two_arr[i][j])
		}
		fmt.Printf("\n")
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++{
			if i == 0 || i == 2 {
				two_arr[i][j] = 0
			} else {
				if j == 0 || j == 3 {
					two_arr[i][j] = 0
				}
			}
		}
	}
	fmt.Println()
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++{
			fmt.Printf("%v\t",two_arr[i][j])
		}
		fmt.Printf("\n")
	}
	
	
}