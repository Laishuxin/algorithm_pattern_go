# 排序算法

| 排序算法 | 复杂度                               | 稳定性 |
| -------- | ------------------------------------ | ------ |
| 冒泡排序 | $O(n^2)$                             | 稳定   |
| 插入排序 | $O(n^2)$                             | 稳定   |
| 选择排序 | $O(n^2)$                             | 不稳定 |
| 奇偶排序 | $O(n^2)$                             | 稳定   |
| 堆排序   | $O(nlog ~n)$                         | 不稳定 |
| 快速排序 | $O(nlog ~n )$                        | 不稳定 |
| 归并排序 | $O(n log ~n)$                        | 稳定   |
| 希尔排序 | $O(n^{3/2- 2} )$（取决于步长的选取） | 不稳定 |
| 桶排序   |                                      |        |





## Heap Sort

1. 构建堆
2. 排序

### 构建堆

1. 从最后一个父节点（`(length - 2) / 2`）出发，自下而上构建一个堆

2. 比较对应两个（可能一个）子节点，选择三者其中最小/大的一个。（**小心数组越界**）

   ```go
   if childIndex+1 < length && arr[childIndex+1] < arr[childIndex] {
       childIndex++
   }
   if temp <= arr[childIndex] {
       break
   }
   ```

3. 更新节点。

   ```go
   arr[parentIndex] = arr[childIndex]
   parentIndex = childIndex
   childIndex = 2*parentIndex + 1
   ```

4. 安置顶层父节点

   ```go
   arr[parentIndex] = temp
   ```

整个过程类似于沿着顶层父节点的插入排序。



### 排序

1. 堆顶挪到数组最后面位置。堆顶就是最大/小值。

   ```go
   Sort.Swap(arr, 0, i)
   ```

2. 从新调整堆顶

   ```go
   _downAdjust(arr, 0, i)
   ```



### 算法复杂度

1. 构建最大/最小堆

   最坏情况为整树的高度$n\times O( log ~n) = O(nlog~n)$

2. 排序

   $n \times O(log ~n) = O(nlog ~n)$

综上：堆排序的算法复杂度为$O(n log ~n)$





## 快速排序

### 双指针快速排序

#### pivot

将排序数组的左区间设置为pivotIndex，然后随机地与数组中的元素调换，避免出现pivot全局最大/小值



#### 快速移动

右指针向左移动，左指针向右移动。确保数组右边是大于Pivot的数，左边是小于Pivot的数。

这就要求在移动的时候，

+ 右指针遇到比Pivot小的数，需要停下来
+ 左指针遇到比Pivot大的数，需要停下来
+ 如果到最后左右指针不相遇，则调换两个指针的索引，然后执行下一趟扫描

```go
for l < r {
    for l < r && arr[r] < pivot {
        r--
    }
    for l < r && pivot <= arr[l] {
        l++
    }
    if l < r {
        Sort.Swap(arr, l, r)
    }
}
```



完成左右分类后，左指针的值一定是小于或等于pivot，交换pivot 与 左指针对应的值。

然后分而治之。



### 简明版本

1. 设置pivot
2. 将小于pivot的存放在一个数组中
3. 将大于pivot的存放在一个数组中
4. 将等于pivot的存放在一个数组中
5. 分而治之
6. 合并

```go
if (curr < pivot) {
    less = append(less, curr)
} else if (curr > pivot) {
    great = append(great, curr)
} else {
    equal = append(equal, curr)
}
less = QuickSortSimplicityVersion(less)
great = QuickSortSimplicityVersion(great)
return append(append(less, equal...), great...)
```



## 归并排序

### 排序

1. 将数组划分为左右两部分，分而治之
2. 排序左边部分
3. 排序右边部分
4. 合并左边部分和右边部分

### 合并

设置三个指针，1个为指向左边部分的索引，1个为指向右边部分的索引，1个指向当前索引

+ 如果右边部分>=左边部分，移动右边部分指针，同时将右边数组中的值赋给当前数组

+ 如果左边部分>右边部分，移动左边部分指针，同时将左边数组中的值赋给当前数组
+ 移动当前数组指针

```go
var (
    l = 0
    r = 0
    c = 0
)
for l < lLen && r < rLen {
    if rPart[r] <= lPart[l] {
        currArr[c] = lPart[l]
        l++
    } else {
        currArr[c] = rPart[r]
        r++
    }
    c++
}
```



判断左边/右边指针是否到达边界，没有达到则说明其中还有部分没有归并。

**注意：合并排序需要借助一个中间数组，才能实现sort in place**





## 奇偶排序

奇偶排序是将数组划分为奇数位和偶数位，分别对奇数、偶数排序从而实现整体排序的效果。



## 希尔排序

希尔排序是插入排序的变种，它将插入排序划分为n step，然后不断缩小步长从而实现数组有序。

值得注意的是**希尔排序**更多地是为了减少插入排序的工作量，希尔排序对输入敏感，这就对步长要求。

综合下来，希尔排序的复杂度为$O(n^{3/2} - 2)$