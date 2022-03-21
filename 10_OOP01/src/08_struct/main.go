package main
import "fmt"

type Dog struct {
	Name string
	Age int
}

func main(){
	var dog *Dog = &Dog{
		Name:"小明",
		Age:2,
	}

	fmt.Printf("%+v", *dog)
}