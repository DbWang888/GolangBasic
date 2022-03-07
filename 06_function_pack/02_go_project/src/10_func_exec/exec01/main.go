package main
import (
	"fmt"
)

func mon_day (mon int, year int) int {
	if mon == 2 {
		if year % 4 == 0 {
			fmt.Println("闰年")
			return 29
		} else {
			fmt.Println("不是闰年")
			return 28
		}
	} else if mon == 4 || mon == 6 || mon == 9 || mon == 11 {
		fmt.Println("月小")
		return 30
	} else {
		fmt.Println("月大")
		return 31
	}
}

func main(){

	var mon, year, day int
	for {
		fmt.Println("输入年")
		fmt.Scanln(&year)
		fmt.Println("输入月")
		fmt.Scanln(&mon)

		if mon <= 0 || mon > 12 {
			fmt.Println("输入有误")
			continue
		}
		if mon > 0 && mon <= 12 {
			day = mon_day(mon, year)
			fmt.Printf("输入的是%d年%d月，本月共%d天\n", year, mon, day)
			continue
		}
		
	}
	




}