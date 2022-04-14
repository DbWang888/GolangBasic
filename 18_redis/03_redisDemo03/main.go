//redis 连接池
package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//定义全局的pool
var pool *redis.Pool

//使用init函数在启动程序时 初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	//先从pool取出连接
	conn := pool.Get()
	defer conn.Close()

	//存入数据
	_, err := conn.Do("set", "pool_name", "tom猫")
	if err != nil {
		fmt.Println("存储错误 err = ", err)
		return
	}
	//读取数据
	r, err := redis.String(conn.Do("get", "pool_name"))
	if err != nil {
		fmt.Println("读取失败 err = ", err)
		return
	}
	fmt.Println("读取内容为 r = ", r)

}
