package main
import "fmt"

//转义字符的使用 "\"
/*
	\t 制表符
	\n 换行
	\\ 表示符号\
	\" 表示符号"
	\r 表示一个回车 1.16.14版本意思为回车后的字符替换最左边字符
*/
func main(){
	fmt.Println("制\t表\t符")
	fmt.Println("换行\n换一行")
	fmt.Println("表示符号，可用于表示路径 c:\\go")
	fmt.Println("表示一个回车\r一个回车")

	//小练习：用一个语句排版输出姓名年龄籍贯住址，并对应一个数据

	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t15\t北京\t成都")
}