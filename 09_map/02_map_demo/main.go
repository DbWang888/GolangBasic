package main
import "fmt"

func main(){

	//map删除及查找
	heroes := map[string]int{
		"jack" : 33,
		"wangfei" : 100,
		"limeng" : 122,
	}

	fmt.Println(heroes)
	fmt.Println()
	//删除map的key及对应的value值
	delete(heroes,"wangfei")

	fmt.Println(heroes)

	// heroes = make(map[string]int)
	// fmt.Println(heroes)

	
	//map查找
	val, findRes := heroes["limeng"]

	if findRes {
		fmt.Println(val)
	} else {
		fmt.Println("xxx")
	}

	
}