package utils
import "fmt"

func Cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("符号有误")

	} 
	GOPATH=C:\Users\wangli\go
	set GOPRIVATE=
	set GOPROXY=https://goproxy.io,direct
	set GOROOT=D:\Golangenvr
	return res
}