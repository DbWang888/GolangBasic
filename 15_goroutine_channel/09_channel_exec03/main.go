package main

func main() {

	// intChan := make(chan int, 2000)
	resChan := make(chan map[int]int, 2000)
	// exitChan := make(chan bool, 8)

	resMap := make(map[int]int, 50)

	resMap[1] = 50
	resChan <- resMap

}
