package process

import "fmt"

//UserMgr实例在服务端只有一个，在很多地方都会使用到
//定义为全局变量
type userMgr struct {
	onlineUsers map[int]*UserProcess
}

var (
	UserMgr *userMgr
)

//完成对UserMgr初始化工作
func init() {
	UserMgr = &userMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//完成对onlineUsers添加
func (um *userMgr) AddOnlineUser(up *UserProcess){
	um.onlineUsers[up.UserId] = up
}

//完成对onlineUsers删除
func (um *userMgr) DelOnlineUser(userId int){
	delete(um.onlineUsers,userId)
}

//返回当前所有在线的用户 
func (um *userMgr) GetAllOnlineUser() map[int]*UserProcess {
	return um.onlineUsers
}

//根据id返回对应的值
func (um *userMgr) GetOnlineUserById(userId int)(up *UserProcess,err error){
	//从map中取出一个值，带检测方式
	up, ok := um.onlineUsers[userId]
	if !ok { //说明查找的这个用户当前不在线
		err = fmt.Errorf("用户 %d 不存在", userId)
		return
	}
	return
}