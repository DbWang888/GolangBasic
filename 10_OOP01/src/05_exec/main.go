package main
import (
	"fmt"
)

type MethodUtils struct {
	Lenth int
	Width int
}

func (m *MethodUtils) printRectangle (key string) int {
	for i := 1; i <= m.Width; i++ {
		for j := 1; j <= m.Lenth; j++ {
			if i == 1 || i == m.Width {
				fmt.Printf("%v", key)
			} else {
				if j == 1 || j == m.Lenth {
					fmt.Printf("%v", key)
				} else {
					fmt.Printf(" ")
				}
			}
		}
		fmt.Printf("\n")
	}
	
	return m.Width * m.Lenth
}

func (m *MethodUtils) areas (w float64, l float64) float64 {
	return w * l
}

func main(){
	rec := MethodUtils {8, 5}
	rec.printRectangle("x")
	area := (&rec).areas(3.0, 4.0)
	fmt.Printf("area = %.2f", area)

}