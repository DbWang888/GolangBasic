package main
import "fmt"

func main(){
	//遍历map  for--range
	heroes := map[string]int{
		"jack" : 33,
		"wangfei" : 100,
		"limeng" : 122,
	}

	for k, v := range heroes {
		fmt.Println(k, v)
	}

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

	for k1, v1 := range student3 {
		fmt.Printf("%v\n", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t\t")
			fmt.Println(k2, v2)
		}
	}
	
	//map长度 len

	fmt.Printf("长度是 %d", len(student3))
}
