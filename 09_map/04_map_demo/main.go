package main

import "fmt"

func main(){
	var map_slice []map[string]int
	map_slice = make( []map[string]int, 2)
	fmt.Println(map_slice)
	if map_slice[0] == nil {
		map_slice[0] = make(map[string]int)
		map_slice[0]["wang"] = 132
		map_slice[0]["li"] = 123
	}

	if map_slice[1] == nil {
		map_slice[1] = make(map[string]int)
		map_slice[1]["zhang"] = 132
		map_slice[1]["song"] = 123
	}

	name := map[string]int{
		"网" : 123,
		"张" : 321,
	}

	map_slice = append(map_slice, name)

	fmt.Println(map_slice)
}