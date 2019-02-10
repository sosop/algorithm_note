# 算法理论

[贪心算法](#贪心算法)



# 算法实践

[排序](#排序)

## 贪心算法

**道：做出当前最好的选择，通过局部最优选择达到全局最优的解决方案；选择不可逆、最优解也许是近似解、贪心策略决定算法质量**

**法：贪心选择、最优子结构**

**术：贪心策略、局部最优解、全局最优解**







## 排序

### 冒泡排序

**每次比较两个相邻元素，如果顺序有误就交换** 

```go
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
```

**复杂度： (num - 1) * (num - i) ~ num^2. ==> O(n^2)**

### 快速排序

**确定基准数，先从左到右，在从右到左的比较，二分思想**

```go
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
```

