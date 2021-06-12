package array

import "fmt"

func PrintArray() {
	var a [3]int

	// 打印数组的长度
	fmt.Println(len(a))

	// 打印数组的下标和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// 打印数组的元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}

func InitArray() {
	var a [3]int = [3]int{1, 2, 3}
	var b [3]int = [3]int{1, 2}
	fmt.Println(a[0])
	// 默认情况下，数组的每个元素都被初始化为元素类型对应的零值
	fmt.Println(b[2])

	// 数组的长度根据初始值的个数来确定
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", c)
}
