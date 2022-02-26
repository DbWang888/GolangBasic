package main
import (
	"fmt"
)

func main(){
	//遍历字符串1  按字节来
	// var str string = "hello,word"
	// for i := 0; i < len(str); i++{
	// 	fmt.Printf("%c\n",str[i])
	// }

	//遍历字符串1  1个汉字3个字节  
	// str = "abc~~ok而且"
	// for index, val := range str {
	// 	fmt.Printf("index=%d, var=%c \n",index,val)

	// }

	//打印1-100之间是9的倍数的个数及总和

	// var sum_num,num int= 0,0
	// for i := 1; i <= 100; i++ {
	// 	if i % 9 ==0 {
	// 		fmt.Printf("%d能被9整除\n", i)
	// 		sum_num = sum_num +i
	// 		num++
	// 	} else {
	// 		fmt.Printf("%d不能被除\n", i)
	// 	}
	// }
	// fmt.Printf("个数为：%d,和为%d\n", sum_num, num)

	//和都等于6
	// var n1, n2 int = 0, 6
	// for i := 0; i <= 6; i++ {
	// 	fmt.Printf("%d + %d = %d \n", n1, n2, n1+n2)
	// 	n1++
	// 	n2--
	// }

	//统计3个班的成绩，每个班5名学生，求出各个班的平均分和所有班级的平均分
	// var banji int = 2
	// var stu int = 3
	// var sum_sc float64
	// for j := 1; j <= banji; j++ {

	// 	sum := 0.0
	// 	for i := 1; i <= stu; i++ {
	// 		var score float64
	// 		fmt.Printf("输入第%d个班请输入第%d个学生的成绩：\n", j, i)
	// 		fmt.Scanln(&score)
	// 		sum = sum + score

	// 	}
	// 	sum_sc += sum
	// 	fmt.Printf("第%d个班级总成绩为：%f 平均分为：%f \n", j, sum, sum / float64(banji))
	// }
	// fmt.Printf("班级总成绩为：%f 平均分为：%f \n", sum_sc, sum_sc / float64(banji * stu))


	//打印金字塔
	// var cengshu int 
	// fmt.Printf("请输入金字塔层数：\n")
	// fmt.Scanln(&cengshu)


	// for CS := 1; CS <= cengshu; CS++ {
	// 	//打印空格
	// 	for i := 1; i <= cengshu - CS; i++ {
	// 		fmt.Printf(" ")
	// 	}
	// 	//打印A
	// 	for j := 1; j <= 2 * CS - 1; j++ {
	// 		if j == 1 || j == 2 * CS -1 || CS == cengshu {
	// 			fmt.Printf("A")
	// 		} else {
	// 			fmt.Printf(" ")
	// 		}
	// 	}
	// 	fmt.Printf("\n")
		
	// }
	
	//九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d X %d = %d\t", j, i, i*j)
			
		}
		fmt.Printf("\n")
	}

}