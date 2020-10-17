package string_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestStringInit(t *testing.T) {
	var s string
	fmt.Println(len(s)) //初始化默认值为""

	// s[1] = '3' //string 是不可变的byte slice
	s = "hello"
	fmt.Println(len(s))

	s = "\xE4\xBB\xA5"
	fmt.Println(s)
	fmt.Println(len(s))

	fmt.Println("---------------")
	s = "中"
	fmt.Println(len(s)) //是byte数

	c := []rune(s)
	fmt.Printf("中 unicode %x", c[0])
	fmt.Println()
	fmt.Printf("中 UTF8 %x", s)
	fmt.Println()
}

//遍历字符串
func TestStringToRune(t *testing.T) {
	fmt.Println("-------------------")
	s := "中华人民共和国"
	for _, c := range s {
		fmt.Printf("%[1]c , %[1]x", c)
		fmt.Println()
	}
}

//字符串分割
func TestStringFn(t *testing.T) {
	fmt.Println("TestStringFn -------------------")
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		fmt.Println(part)
	}

	//字符合并
	fmt.Println(strings.Join(parts, "-"))
}

//类型转换，string转换成整数
func TestConv(t *testing.T) {
	fmt.Println("TestConv -------------------")
	//整数转换为字符串
	s := strconv.Itoa(10)
	fmt.Println("str:" + s)

	//字符串转换为整型
	if i, err := strconv.Atoi("10"); err == nil {
		fmt.Println(10 + i)
	}
}
