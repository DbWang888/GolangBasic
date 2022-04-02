package testdemo02

import (
	"fmt"
	"os"
	"testing"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err != nil {
		fmt.Println("存在")
		return true, nil
	}

	if os.IsNotExist(err) {
		fmt.Println("不存在")
		return false, nil
	}
	return false, err
}

func TestMonStore(t *testing.T) {
	filePath := "D:\\ProgramPro\\GolangBasic\\GolangBasic\\14_unit_testing\\02_testdemo02\\montser.txt"
	UseStore()
	sf, err := PathExists(filePath)
	fmt.Println(sf)
	if err != nil || sf == false {
		t.Fatalf("序列化程序错误")
	}
	t.Logf("文件创建成功")
}

func TestResotre(t *testing.T) {
	UseRestore()
}
