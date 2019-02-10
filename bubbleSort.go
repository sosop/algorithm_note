package main

import "fmt"

func main() {
    s := []int{10, 7, 100, 99, 8, 66, 78, 33, 57, 24}
    fmt.Println(s)
    bubbleSort(s)
	fmt.Println(s)
}

// 冒泡从小到到排序
func bubbleSort(ori []int) {
    num := len(ori)
    // num个数只需要num - 1趟的两两比较
    for i := 0;i < num - 1; i++ {
        // 每趟需要比较的次数
        for j := 0; j < num - i - 1; j ++ {
            if ori[j] > ori[j+1] {
                ori[j], ori[j+1] = ori[j+1], ori[j]
            }
        }
    }
}
