package main
import "fmt"


func num (n int) (bool){
	var num1 int = 0
	for i := 2; i < n; i++ {
		if n % i == 0 {
			num1++
		}
	}
	if num1 == 0 {
		return true
	} else {
		return false
	}
}

func main(){

	var sum int = 1
	var count int = 0

	for i := 2; i <= 100; i++ {
		if num(i) == true {
			sum = sum + i
			count++
			fmt.Printf("%d\t", i)
			if count % 5 == 0 {
				fmt.Printf("\n")
			}
		}
	}

	fmt.Println("和为", sum)
}