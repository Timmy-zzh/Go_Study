package map_ext

import (
	"fmt"
	"testing"
)

//map扩展功能
//map的value可以是函数
func TestMapWithFunValue(t *testing.T) {
	//1.map定义
	// key值为int型
	//value为func函数，参数为int，返回值为int
	m1 := map[int]func(op int) int{}
	m1[1] = func(op int) int { return op }
	m1[2] = func(op int) int { return op * op }
	m1[3] = func(op int) int { return op * op * op }
	fmt.Println(m1[1](2), m1[2](2), m1[3](2))
}

//使用map实现Set功能
func TestMapForSet(t *testing.T) {
	fmt.Println("TestMapForSet")
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		fmt.Println(n, "is existing")
	} else {
		fmt.Println(n, "is not existing")
	}
	mySet[3] = true
	fmt.Println(len(mySet))

	//删除一个元素
	delete(mySet, 1)
	fmt.Println(len(mySet))
}
