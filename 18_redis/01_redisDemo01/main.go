package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {

	//go链接数据库redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接失败， err = ", err)
		return
	}
	defer conn.Close() //用完后关闭

	//通过go写入string数据  [key - value]
	_, err = conn.Do("Set", "name", "jack猫猫")
	if err != nil {
		fmt.Println("string写入失败 err = ", err)
		return
	}

	//通过go 向 redis读取数据 string[key-val]
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("string读取失败 err = ", err)
		return
	}
	//返回的r为interface{}，需要类型断言  redis.String直接转换
	// nameString := r.(string)
	fmt.Println("string r = ", r)

	//通过go向redis写入哈希数据类型

	_, err = conn.Do("hset", "gouser01", "name", "john")
	if err != nil {
		fmt.Println("hash写入失败 err = ", err)
		return
	}

	//通过go读取hash数据

	hr, err := redis.String(conn.Do("hget", "gouser01", "name"))
	if err != nil {
		fmt.Println("hash读取失败 err = ", err)
		return
	}
	fmt.Println("HASH HR = ", hr)
}
