package main
import "fmt"

func main(){

	a := 10
	b := 5

	a = a + b
	b = a - b
	a = a - b

	fmt.Println(a,b)

	c := 5
	d := "she"
	e := 6.6

	fmt.Scanln(&a)
	fmt.Scanf("%s %v", &d, &e)
	fmt.Println(c, d , e)

}