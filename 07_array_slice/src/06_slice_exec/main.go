package main
import "fmt"


func fbnq(n int)( []int ){
	var fbnq_slice []int = make([]int,n)
	if n == 1 {
		fbnq_slice[0] = 1
		return fbnq_slice
	} else if n == 2{
		fbnq_slice[0] = 1
		fbnq_slice[1] = 1
		return fbnq_slice
	} else {
		fbnq_slice[0] = 1
		fbnq_slice[1] = 1
		for i := 2; i < n; i++ {
			fbnq_slice[i] = fbnq_slice[i-1] + fbnq_slice[ i -2]
		}
		return fbnq_slice
	}	

	
}
func main(){
	
	fb := fbnq(10)
	fmt.Println(fb)
}