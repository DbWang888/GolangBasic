//工厂模式测试，在别的包也能使用小写字母开头的内容

package model

type student struct{
	Name string
	age int
}

func NewStudent (name string, age int) *student {
	return &student {
		Name : name,
		age : age,
	}
}

func (s *student) GetScore() int {
	return s.age
}