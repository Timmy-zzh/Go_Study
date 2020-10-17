#### String字符串

- Go语言字符串与其他编程语言差异
  - string是数据类型，不是引用或指针类型
  - string是只读的byte slice，len函数得到是他所包含的byte数
  - string的byte数组可以存放人恶化数据

- 数组声明与使用

~~~go
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
~~~

- Unicode UTF8
  - Unicode是一种字符集（code point）
  - UTF8是unicode的存储实现（转换为字节序列的规则）

- 编码与存储

|     字符      |        中        |
| :-----------: | :--------------: |
|    Unicode    |      0x4E2D      |
|     UTF-8     |     0xE4B8AD     |
| string/[]byte | [0xE4,0xb8,0xAD] |

- 遍历字符串

~~~go
//遍历字符串
func TestStringToRune(t *testing.T) {
	fmt.Println("-------------------")
	s := "中华人民共和国"
	for _, c := range s {
		fmt.Printf("%[1]c , %[1]x", c)
		fmt.Println()
	}
}
~~~

- 字符串分割与合并

~~~go
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
~~~

- 字符串类型转换

~~~go
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
~~~









