#### Map

- map声明

~~~go
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
~~~

- map访问不存在的key值
  - 在访问的key不存在时，仍会返回零值，不能通过返回nil来判断元素是否存在

~~~go
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
~~~

- map元素遍历

~~~go
// map遍历 使用range关键字
func TestTravelMap(t *testing.T) {
	fmt.Println("TestTravelMap")
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
~~~



#### Map扩展功能

- map的vaule可以用函数表示

~~~go
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

2 4 8
~~~

- 通过map实现Set功能

  - Go内置集合中没有Set实现，可以通过map[type]bool实现

  1. 元素的唯一性
  2. 基本操作
     1. 添加元素
     2. 判断元素是否存在
     3. 删除元素
     4. 元素个数

~~~go
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
~~~

