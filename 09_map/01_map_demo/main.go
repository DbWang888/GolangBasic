package main
import "fmt"

func main(){
	//map声明方式1 var name map[type]type
	var student1 map[string]int
	student1 = make(map[string]int, 3)

	student1["jack"] = 10
	student1["marry"] = 15
	fmt.Println(student1)

	//第二种方式  name := make(map[type]type)

	var student2 map[string]int = make(map[string]int)
	student2["jack"] = 222222
	fmt.Println(student2)
	//第三种方式，声明的同时赋值  
	heroes := map[string]int{
		"jack" : 33,
		"wangfei" : 100,
		"limeng" : 333,
	}

	fmt.Println(heroes)


	//test

	student3 := map[string]map[string]string{
		"stu1" : {
			"name" : "lilei",
			"sex" : "man",
		},
		"stu2" : {
			"name" : "hanmei",
			"sex" : "woman",
		},

	}

	student3["stu3"] = make(map[string]string)
	student3["stu3"]["name"] = "liqing" 
	fmt.Println(student3)

}