package condition_test

import (
	"fmt"
	"testing"
)

// if条件判断语句
func TestIfCondition(t *testing.T) {
	fmt.Println("TestIfCondition")
	n := 2
	if a := n == 1; a {
		fmt.Println("1==1")
	} else {
		fmt.Println("1!=1")
	}
}

//switch判断语句
func TestSwitch(t *testing.T) {
	fmt.Println("TestSwitch")
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			fmt.Println(i, "Even")
		case 1, 3:
			fmt.Println(i, "Odd")
		default:
			fmt.Println(i, "it is not 0-3")
		}
	}
}

// 在Switch语句的case中进行条件判断
func TestSwitchCondition(t *testing.T) {
	fmt.Println("TestSwitchCondition")
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "Even")
		case i%2 == 1:
			fmt.Println(i, "Odd")
		default:
			fmt.Println(i, "unknow")
		}
	}
}
