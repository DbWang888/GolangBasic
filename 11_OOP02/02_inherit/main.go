package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GetResult() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(99) + 1
}

type Student struct {
	Name           string
	Age            int
	Student_number int
	Achievement    int
}

type Pupil struct {
	Student
	InterestClass string
}

func (p *Pupil) test() int {
	fmt.Println("小学生考试完成")
	return GetResult()
}

type Graduate struct {
	Student
	Professional_course string
}

func (g Graduate) test() int {
	fmt.Println("大学生考试完成")
	return GetResult()
}

type Goods struct {
	Area string
	Sort string
}

type Books struct {
	area_sort Goods
	Name      string
	Price     float64
}

func (g *Goods) SetAreaAndSort(a string, s string) {
	g.Area = a
	g.Sort = s
}

func main() {
	var jack Pupil
	jack.Student.Name = "小学生"
	jack.Achievement = jack.test()
	jack.InterestClass = "足球"
	fmt.Println(jack)
	fmt.Println()
	fmt.Println()

	var book1 Books
	book1.area_sort.SetAreaAndSort("山东", "书籍")
	book1.Name = "三字经"
	book1.Price = 22.9
	fmt.Printf("%+v\n", book1)
}
