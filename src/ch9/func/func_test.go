package func_test

import (
	"fmt"
	"math/rand"
	// "math/rand"
	"testing"
	"time"
)

//1.函数有多个返回值
//该函数名为returnMulValus； 没有入参，返回值为两个int类型
func returnMulValus() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnMulValus()
	fmt.Println(a, b)

	tsSf := timeSpent(slowFun)
	fmt.Println(tsSf(10))

	fmt.Println("sum-------------------")
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))
}

//函数作为参数传入，并计算函数耗时
//函数参数：inner func(op int) int
//返回值：func(op int )int
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend:", time.Since(start).Seconds())
		return ret
	}
}

//耗时操作
func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

//3.函数的可变参数
func sum(ops ...int) int {
	s := 0
	for _, op := range ops {
		s += op
	}
	return s
}

func Clear() {
	fmt.Println("Cleanr resources...")
}

//defer函数--defer函数会延迟执行
func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	//报错的意思
	panic("err")
}
