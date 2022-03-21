package main

import (
	"fmt"
)

//定义一个cat结构体数据类型
type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
	Skill map[int]string //map slice *int等引用类型使用前需要先make 分配内存空间再使用
	Ptr   *int
}

type Person struct {
	Name  string
	Age   int
	Score float32
}

func main() {

	//第一种使用方法 先声明，后赋值
	var num int = 500
	var cat_jack Cat
	cat_jack.Name = "jack"
	cat_jack.Age = 18
	cat_jack.Color = "white"
	cat_jack.Hobby = "吃鱼"
	cat_jack.Skill = make(map[int]string, 3)
	cat_jack.Skill[1] = "抓鱼"
	cat_jack.Skill[2] = "杀鱼"
	cat_jack.Skill[3] = "摸鱼"
	// cat_jack.Ptr = new(int)
	cat_jack.Ptr = &num

	//打印
	fmt.Println("cat_jack的内容", cat_jack)

	//第二种使用方法  声明的同时赋值
	person1 := Person{"tom", 15, 99.5}
	fmt.Println("person的内容", person1)

	//第三种方法使用指针，new->  var person *Person = new(Person)
	var person2 *Person = new(Person)
	(*person2).Name = "smith" //等价于person.Name = "smith"
	(*person2).Age = 55
	(*person2).Score = 12.4
	fmt.Println(*person2)

	//第四种方法使用指针，new->  var person *Person = &Persion{}
	var person3 *Person = &Person{}
	person3.Name = "marry" //等价于person.Name = "smith"
	person3.Age = 32
	person3.Score = 19.9
	fmt.Printf("%v\n", person3) //person3 会输出&{}-->&{marry 32 19.9}，*person3直接输出{}-->{marry 32 19.9}

}
