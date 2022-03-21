package main
import "fmt"

func main(){
	var arr [4][4]int = [4][4]int{ {1,2,3,4} , {1,3,2,4}, {2,3,2,1}, {2,3,1,4} }

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++{
			fmt.Printf("%v\t",arr[i][j])
		}
		fmt.Printf("\n")
	}
	
	var arr1 [4]int 

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++{
			if i == 0 {
				arr1[j] = arr[0][j]
				arr[0][j] = arr[3][j]
				arr[3][j] = arr1[j]
			}  else if i == 1 {
				arr1[j] = arr[1][j]
				arr[1][j] = arr[2][j]
				arr[2][j] = arr1[j]
			} else {
				continue
			}
		}
	}
	fmt.Printf("\n\n")
	fmt.Println(arr1)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++{
			fmt.Printf("%v\t",arr[i][j])
		}
		fmt.Printf("\n")
	}

}
