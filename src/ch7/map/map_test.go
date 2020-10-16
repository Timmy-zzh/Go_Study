package map_test

import (
	"fmt"
	"testing"
)

//Map初始化
func TestMapInit(t *testing.T) {
	//方式1
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	fmt.Println(m1[2])
	fmt.Println("len m1=", len(m1))

	//方式二
	m2 := map[int]int{}
	m2[4] = 16
	fmt.Println("len m2=", len(m2))

	//方式三 , 10代表容量cap
	m3 := make(map[int]int, 10)
	fmt.Println("len m3=", len(m3))
}

//不存在的key值
func TestNotExistingKey(t *testing.T) {
	fmt.Println("TestNotExistingKey")
	m1 := map[int]int{}
	fmt.Println(m1[1])
	m1[2] = 0
	fmt.Println(m1[2])

	//如何判断key值是否存在呢？
	m1[3] = 0
	if v, ok := m1[3]; ok {
		fmt.Println("key 3 exist value is", v)
	} else {
		fmt.Println("key 3 not exist")
	}
}

// map遍历 使用range关键字
func TestTravelMap(t *testing.T) {
	fmt.Println("TestTravelMap")
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
