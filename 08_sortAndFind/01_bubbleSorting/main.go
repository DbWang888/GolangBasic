package main
import "fmt"
func main () {
//JSRUN引擎2.0，支持多达30种语言在线运行，全仿真在线交互输入输出。
fmt.Println("Hello JSRUN! \n\n - from Golang .")
var arr [5]int = [5]int{11,2,23,1,13}
// var arr [3]int = [3]int{23,1,13}
var n int
for i := 0; i < (len(arr) - 1); i++{
    for i := 0; i <= (len(arr) - 2); i++{
        if arr[i] > arr[i+1] {
        n = arr[i] 
        arr[i] = n
        arr[i] = arr[ i + 1]
        arr[i + 1] = n
        }
    }
}
fmt.Println(arr)
}