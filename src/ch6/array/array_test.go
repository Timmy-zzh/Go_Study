package array

import (
	"fmt"
	"testing"
)

func TestArrayInit(t *testing.T) {
	//数组声明1
	var a [3]int
	a[0] = 1
	fmt.Println(a)
	//数组声明2，二维数组声明
	b := [3]int{1, 2, 3}
	c := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(b)
	fmt.Println(c)
	//数组声明3
	d := [...]int{1, 2, 3, 4}
	fmt.Println(d)
}

//数组元素遍历
func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for idx, ele := range arr {
		fmt.Println(idx, ele)
	}

	//不需要数组下标
	for _, ele := range arr {
		fmt.Println(ele)
	}
}

//数组截取
func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	arrSec := arr[1:3] //左闭右开
	fmt.Println(arrSec)
}
