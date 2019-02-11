# 算法理论

[贪心算法](#贪心算法)

# 数据结构

[线性表](#线性表)

# 算法实践

[排序](#排序)



## 贪心算法

**道：做出当前最好的选择，通过局部最优选择达到全局最优的解决方案；选择不可逆、最优解也许是近似解、贪心策略决定算法质量**

**法：贪心选择、最优子结构**

**术：贪心策略、局部最优解、全局最优解**



## 数据结构

### 线性表

**零个或多个数据元素的有限序列**

#### 顺序存储

```go
type List struct {
   	array 		unsafe.Pointer // 存储数据，指向一块连续的地址
    Len 		int	  // 长度
    Cap 		int   // 容量
}
// 与slice相同
```

*优缺点：能够实现快速存储，但是插入删除会有数据移动，而且存在扩容问题以及磁盘碎片问题*

#### 链式存储

基本概念：节点、头节点、头指针、数据域、指针域

分类：单链表、静态链表(第一个元素存放备用链表的第一个位置，最后一个元素作为头节点使用)、循环链表(将终端节点的指针指向头节点)、双向链表(在原节点上增加前驱指针域，与后继指针域结合起来)

```go
// 单链表
type Node struct {
    Value	dataType
    Next 	*Node
}
// 静态链表
type StaticNode struct {
    Value	dataType
    Current	int
}
```

优缺点：插入删除快速（已找到元素前提下），节约存储空间，但是查询是需要遍历



### 栈

**表尾进行插入和删除的线性表**

```go
type Stack struct {
    Value	dataType
    Top		int
}
```

### 队列

**一端进行插入另一端进行删除的线性表**

```go
type Queue struct {
    Value	dataType
    Head	int
    Tail	int
}
```



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

