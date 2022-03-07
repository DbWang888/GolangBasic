package main
import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	//统计字符串长度，len()，中文字符占3个字节
	str1 := "Golang基础练习！"
	fmt.Println("str1长度为：", len(str1))

	//遍历字符串
	for i := 0; i < len(str1); i++ {
		fmt.Printf("str[%d] = %c\n", i, str1[i])
	}
	//转为切片，可遍历中文
	r := []rune(str1)
	for i := 0; i < len(r); i++ {
		fmt.Printf("r[%d] = %c\n", i, r[i])
	}

	//字符串转整数

	var str2 string = "123456"
	var str3 string = "abc"

	n1, err1 := strconv.Atoi(str2)
	if err1 != nil {
		fmt.Println("错误，err = ", err1)
	} else {
		fmt.Println(n1)
	}

	n2, err2 := strconv.Atoi(str3)
	if err2 != nil {
		fmt.Println("错误，err = ", err2)
	} else {
		fmt.Println(n2)
	}

	//整数转字符串

	n3 := strconv.Itoa(12345)
	fmt.Printf("n3 = %v, n3类型为%T\n", n3, n3)

	//字符串转[]byte 字符切片，返回对应的assic码

	var bytes = []byte("hello go")
	fmt.Printf("%v\n", bytes)

	//[]byte转字符串

	str4 := string([]byte{97,98,99})
	fmt.Printf("%v\n", str4)


	//10进制转2，8，16进制  ,返回对应的字符串 ,2可换成8，16

	str5 := strconv.FormatInt(123,2) 
	fmt.Printf("%v\n", str5)
	str6 := strconv.FormatInt(123,8) 
	fmt.Printf("%v\n", str6)
	str7 := strconv.FormatInt(123,16) 
	fmt.Printf("%v\n", str7)

	//查找子串是否在指定字符串中， 返回true 或者 false

	b1 := strings.Contains("你今天是很好还是不好", "很好")
	fmt.Printf("%v\n", b1)

	//查找一个字符串有几个指定的字串，返回int类型

	num1 := strings.Count("你今天是很好还是不好","很好")
	fmt.Printf("%v\n,类型是%T\n", num1, num1)

	//不区分字符串大小写比较   “==” 区分大小写
	
	b2 := strings.EqualFold("abc", "ABc")
	fmt.Printf("%v\n", b2)

	//指定字符串替换
	str8 := "我是帅哥吗？是美女?7788"
	str9 := strings.Replace(str8, "美女", "帅哥", 1)
	fmt.Printf("%v\n", str9)

	//按照指定的字符为分割标识，将字符串拆分成字符串数组

	strArr1 := strings.Split(str8, "？")
	fmt.Printf("%v\n", strArr1)
	for i := 0; i < len(strArr1); i++ {
		fmt.Printf("%v\n", strArr1[i])
	}

	//将字符进行大小写转化

	str10 := "AAABBBcccddd"

	str11 := strings.ToLower(str10)
	str12 := strings.ToUpper(str10)
	fmt.Printf("%v-----%v\n", str11, str12)


	//判断字符串以指定字符串结尾或开头

	str13 := "http:xxxxxxxxx.com"
	b3 := strings.HasPrefix(str13, "http")
	b4 := strings.HasSuffix(str13, ".com")

	fmt.Printf("b3 = %v\n b4 = %v", b3, b4)




	

}