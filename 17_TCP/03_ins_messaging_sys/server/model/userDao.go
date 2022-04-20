package model

import (
	"03_ins_messaging_sys/03_ins_messaging_sys/common/message"
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//服务器启动后，就初始化一个UserDao的实例
//做成全局的变量，需要和redis操作时，直接使用即可
var (
	MyUserDao *UserDao
)

//定义一个UserDao结构体，完成对User结构体的各种操作
type UserDao struct {
	Pool *redis.Pool
}

//使用工厂模式，创建UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		Pool: pool,
	}
	return
}

//根据一个用户Id 返回要给 User 实例 + err
func (u *UserDao) GetUserById(conn redis.Conn, id int) (user *User, err error) {
	//通过给定Id 去redis查询这个用户
	res, err := redis.String(conn.Do("Hget", "users", id))
	if err != nil {
		if err == redis.ErrNil { //表示在user哈希中，没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	//需要把res反序列化成User实例（取出的res为json字符串）
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}
	return
}

//完成登录校验
//Login完成对用户的验证，如果Id和pwd都正确，则返回一个User实例
//如果用户的Id或者pwd错误，则返回对应的错误信息
func (u *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	//先从UserDao的连接池中取出一个连接
	conn := u.Pool.Get()
	defer conn.Close()
	user, err = u.GetUserById(conn, userId)
	if err != nil {
		return
	}

	//此时证明用户获取到，对信息进行校验
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (u *UserDao) Register(user *message.User) (err error) {

	//先从UserDao的连接池中取出一个连接
	conn := u.Pool.Get()
	defer conn.Close()
	_, err = u.GetUserById(conn, user.UserId)
	if err == nil { //根据id从数据库中取出对应的信息，如果不报错，则该Id对应的用户存在
		err = ERROR_USER_EXISTS
		return
	}

	//此时 说明id不存在，可以完成注册
	fmt.Println("json.Marshal(user) = ", user)
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal(user) err = ", err)
		return
	}

	//将用户信息写入数据库
	_, err = conn.Do("hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("写入数据库错误 err=", err)
		return
	}

	return
}
