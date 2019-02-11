package main

import "fmt"

func main() {
    s := []int{10, 7, 100, 99, 8, 66, 78, 33, 57, 24}
    fmt.Println(s)
    quickSort(s)
    fmt.Println(s)
}

// 由小到大的快速排序
func quickSort(ori []int) {
    num := len(ori)
    leftIndex, rightIndex := 1, num - 1

    if leftIndex > rightIndex {
        return
    }

    if leftIndex == rightIndex && ori[0] > ori[leftIndex] {
        ori[0], ori[leftIndex] = ori[leftIndex], ori[0]
    }

    // leftindex等于rightindex，一轮比较完成
    for leftIndex != rightIndex {
        // 先从右开始比较
        for rightIndex > leftIndex && ori[0] <= ori[rightIndex] {
            rightIndex --
        }
        // 再从右到左比较
        for leftIndex < rightIndex && ori[0] >= ori[leftIndex] {
            leftIndex ++
        }
        // 交换
        if leftIndex < rightIndex {
            ori[leftIndex], ori[rightIndex] = ori[rightIndex], ori[leftIndex]
        }
    }

    // 相等后，将基准与当前index交换，再利用递归二分
    ori[0], ori[leftIndex] = ori[leftIndex], ori[0]
    quickSort(ori[:leftIndex])
    quickSort(ori[leftIndex+1:])
}
