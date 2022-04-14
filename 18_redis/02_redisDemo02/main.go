package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	//链接数据库
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("数据库连接失败 err = ", err)
		return
	}
	defer conn.Close()

	//批量写入数据
	_, err = conn.Do("hmset", "gouser02", "name", "john", "age", 18)
	if err != nil {
		fmt.Println("数据写入失败 err = ", err)
		return
	}

	//读取数据
	//string
	// r, err := redis.String(conn.Do("hget", "gouser02", "name"))
	// if err != nil {
	// 	fmt.Println("数据读取失败 err = ", err)
	// 	return
	// }
	// //int
	// age, err := redis.Int(conn.Do("hget", "gouser02", "age"))
	// fmt.Printf("name = %v， age = %v", r, age)

	//批量读取数据
	r, err := redis.Strings(conn.Do("HMGet", "gouser02", "name", "age"))
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Printf("%v\n", r)
	for i, v := range r {
		fmt.Printf("i = %v, v = %v\n", i, v)
	}

}
