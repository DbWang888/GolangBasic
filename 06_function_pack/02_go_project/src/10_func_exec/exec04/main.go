package main
import (
	"fmt"
	"time"
)

func get_days(year int, mon int, day int) float64 {
	t := time.Date(year,time.Month(mon),day,00,00,00,00,time.UTC)
	u := time.Date(1990,01,01,00,00,00,00,time.UTC)
	sbs := t.Sub(u).Hours()
	days := sbs / 24 
	return days
}

func main(){
	//三天打鱼两天晒网
	var year, mon, day int
	for {
		fmt.Println("请输入年")
		fmt.Scanln(&year)
		fmt.Println("请输入月")
		fmt.Scanln(&mon)
		fmt.Println("请输入日")
		fmt.Scanln(&day)
		days := get_days(year,mon,day)
		if days >= 0 && days <= 5 {
			if days == 0 || days == 1 {
				fmt.Println("晒网")
			} else {
				fmt.Println("打鱼")
			}
		} else {
			sbs_day := int64(days) % 5
			if sbs_day == 1 || sbs_day == 2 {
				fmt.Println("晒网")
			} else {
				fmt.Println("打鱼")
			}
		} 	
		
	}
	
}
