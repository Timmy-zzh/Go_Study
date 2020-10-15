package slice_test

import (
	"fmt"
	"testing"
)

//切片声明与获取长度和容量
func TestSliceInit(t *testing.T) {
	var s0 []int
	fmt.Println(len(s0), cap(s0))
	//添加切片元素,添加元素后需要返回新的切片对象
	s0 = append(s0, 1)
	fmt.Println(len(s0), cap(s0))

	//切片第二种声明方式
	s1 := []int{1, 2, 3, 4}
	fmt.Println(len(s1), cap(s1))

	//第三种, []int 参数表示类型，3为len，5为cap
	s2 := make([]int, 3, 5)
	fmt.Println(len(s2), cap(s2))
	//这种情况下元素获取只能获取3个
	// fmt.Println(s2[0], s2[1], s2[2], s2[3], s2[4])
	fmt.Println(s2[0], s2[1], s2[2])
}

//切片可变长
//发现切片容量的是指数级别的增长，而且每次都会返回不同地址的对象
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Println(len(s), cap(s))
	}
}

//切片共享存储空间
//切片的容量是整个大切片的尾部容量
//切片部分元素改变会改变整体元素的值
//切片因为长度是变长的，所以不可进行两个切片的比较
func TestSliceShareMemory(t *testing.T) {
	year := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	tt := year[3:6]
	fmt.Println(len(tt), cap(tt))

	sum := year[5:8]
	fmt.Println(len(sum), cap(sum))
	fmt.Println(sum)
	fmt.Println(tt)
	sum[0] = "99"
	fmt.Println(tt)
}
