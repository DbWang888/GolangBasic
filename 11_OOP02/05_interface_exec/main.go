package main

import (
	"fmt"
	"sort"
)

type Hero struct {
	Name  string
	Age   int
	Score float64
}

type Heroes []Hero //定义Hero切片类型

//定义Heroes切片的Less方法
func (h Heroes) Len() int {
	return len(h)
}

//定义Heroes切片的Less方法

func (h Heroes) Less(i, j int) bool {
	return h[i].Name < h[j].Name
}

//定义Heroes切片的Swap方法
func (h Heroes) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {

	//定义3个英雄
	hero1 := Hero{
		Name:  "小明",
		Age:   19,
		Score: 88.8,
	}

	hero2 := Hero{
		Name:  "大明",
		Age:   25,
		Score: 98.8,
	}

	hero3 := Hero{
		Name:  "中明",
		Age:   22,
		Score: 66.6,
	}

	var heroes Heroes

	heroes = append(heroes, hero1)
	heroes = append(heroes, hero2)
	heroes = append(heroes, hero3)

	for _, v := range heroes {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	sort.Sort(heroes)

	for _, v := range heroes {
		fmt.Printf("%+v\n", v)
	}
}
