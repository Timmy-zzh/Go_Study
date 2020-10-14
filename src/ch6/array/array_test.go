package array

import (
	"fmt"
	"testing"
)

func TestArrayInit(t *testing.T) {
	var a [3]int
	a[0] = 1
	fmt.Println(a)

	b := [3]int{1, 2, 3}
	c := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(b)
	fmt.Println(c)

	d := [...]int{1, 2, 3, 4}
	fmt.Println(d)
}

//数组元素循环
func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for idx, ele := range arr {
		fmt.Println(idx, ele)
	}
}
