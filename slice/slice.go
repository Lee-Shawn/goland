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

func Slice1() {
	// 切片底层是数组，切片的第一个元素不一定是数组 的第一个元素
	mouths := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	Q2 := mouths[4:7]
	summer := mouths[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)
	// 查询两个切片中是否有相同的月份，这种方式性能比较低
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
}
