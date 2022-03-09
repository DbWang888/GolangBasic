package main
import "fmt"

func main(){
	str := "sliceandstring"
	fmt.Printf("%v\n",str[8:])

	//修改字符串需要将字符串转成->[]byte 或者 ->[]rune 之后再修改
	//修改完成后重新转成string类型
	//byte按照字节处理，只能处理英文及符号，汉字占3个字节，用byte就会出现错误
	//处理中文，需要用到-> []rune即可  rune是按照字符处理的
	str_arr := []rune(str)
	str_arr[0] = '我'
	str = string(str_arr)
	fmt.Printf("%v\n",str)
}