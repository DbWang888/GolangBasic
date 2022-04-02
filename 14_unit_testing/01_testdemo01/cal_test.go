package testdemo01

import (
	"testing"
)

func TestAdd(t *testing.T) {

	//调用Add函数

	res := Add(20, 30)
	if res != 40 {
		//t.Fatalf输出自定义错误信息
		t.Fatalf("Add执行错误，错误值为%v，正确是%v", res, 40)
	}

	//如果正确，输出日志
	t.Logf("Add执行正确")
}

func TestGetSum(t *testing.T) {
	res := GetSum(10)
	if res != 55 {
		t.Fatalf("getsum执行错误当前值为： %v", res)
	}

	t.Logf("执行正确  %v", res)
}
