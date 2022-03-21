package main
import "fmt"

func modifyUser( users map[string]map[string]string, name string){
	if users[name] == nil {
		users[name] = make(map[string]string,2)
		users[name]["nickname"] = "xiaoming"
		users[name]["pwd"] = "123456"
	} else {
		users[name]["pwd"] = "888888"
	}
}


func main(){
	users := make(map[string]map[string]string, 10)
	users["jack"] = make(map[string]string, 2)
	users["jack"]["nickname"] = "dawamg"
	users["jack"]["pwd"] = "aaa"
	fmt.Println(users)
	name := "jack"
	modifyUser(users, name)
	fmt.Println(users)
}
