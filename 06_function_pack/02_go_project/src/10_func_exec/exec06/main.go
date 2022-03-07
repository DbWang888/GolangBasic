package main
import (
	"fmt"
	"strings"
)

func main(){
	//打印a-z以及Z-A
	for i := 65; i <= 90; i++ {
		fmt.Printf("%c", i)
	}
	for i := 90; i >= 65; i-- {
		char := fmt.Sprintf("%c", i)
		a_char := strings.ToLower(char)
		fmt.Printf("%v", a_char)
	}
}