package slice

import "fmt"

func SubSlice() {
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	// 从第一个位置截取到第三个位置
	slice := underlyingArray[0:3]
	fmt.Println(slice)
	// 从第一个位置开始截取
	slice1 := underlyingArray[:3]
	fmt.Println(slice1)
	// 截取到最后一个位置
	slice2 := underlyingArray[3:]
	fmt.Println(slice2)
}
