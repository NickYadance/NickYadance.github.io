---
title: '理解二分搜索的上下界问题'
date: '2023-11-10'
description: '尽管二分查找的基本思想相对简单，但细节可以令人难以招架...'
---
> 尽管二分查找的基本思想相对简单，但细节可以令人难以招架 — 高德纳

二分查找是一个迭代算法，基本思想是通过不断折半缩小查找范围，最终找到目标值。

如果只对二分查找的代码实现死记硬背，在编码过程中容易对上下界和终结条件发晕导致无法写出BugFree的代码。我们用数学的思想循环不变量来描述迭代过程，避免发晕。循环不变量的原理类似多米诺骨牌，包括四个要素
1. $Variant$: 定义
2. $Start$: 初始化
3. $Loop$: 保持
4. $End$: 终止

首先我们定义二分查找的循环不变量为: 对于有序数组$a[l,r)$和目标值$n$，在**数组区间$[l, r)$内，区间左边的元素小于$n$，区间右边的元素大于等于$n$，即**
$$ a[i]<n \ if \ i < l $$
$$ a[i]>=n \ if \ i >= r $$
算法开始时有$l=0,r=len(a)$，循环不变量成立。

不断缩小这个区间直到$l=r$，此时$l$左边的元素都小于$n$，$l$右边的元素(包括自身)都大于等于$n$，我们得到一个有用的数学性质，即$a[l]$是目标值n在数组$a$中的**下界**，循环不变量定义的区间就是**目标值下界所在的区间**。

> 数组中N的下界是第一个大于或等于N的元素下标

保持过程就是要寻找下一个下界所在的子区间。
1. 如果$n > a[mid]$，n的下界在mid右边且不包括mid，区间$[mid + 1, r)$满足循环不变量
2. 如果$n <= a[mid]$，n的下界在mid左边且包括mid，区间$[l, mid)$满足循环不变量

对照四要素，就可以写出**寻找目标值下界**的二分搜索算法。
```go
func BinarySearchLowerBound(arr []int, n int) int {
	low := 0
	high := len(arr)
	for low < high {
		mid := low + (high-low)>>1
		if n > arr[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
```

二分查找可以在二分搜索下界的基础上实现，在迭代过程中提前判断$a[mid]=n$可以减少迭代次数，但丢失掉了下界算法的**稳定性**，因为下界只会有一个值。
```go
func BinarySearch(arr []int, n int) int {
	lowerBound := BinarySearchLowerBound(arr, n)
	if lowerBound >= 0 && lowerBound < len(arr) && arr[lowerBound] == n {
		return lowerBound
	}
	return -1
}
```

与下界相对应的上界怎么求呢？数学定义里上下界对应的是一个数在排序数组的两端，例如下面例子里3的下界是a[2]，上界是a[4]。但是套用我们的下界算法并不能得到上界为4(你可以试试)，这时候可以做个处理，定义上界为**数学上界的右一位**，即例子里的a[5]。
```
a[i] 1 2 3(lower) 3 3(upper) 4(upper')
i    0 1 2        3 4        5
```

这种处理方式参考的是c++ std里[upper_bound](https://en.cppreference.com/w/cpp/algorithm/upper_bound)&[lower_bound](https://en.cppreference.com/w/cpp/algorithm/lower_bound)的定义，上下界分别是在数组中**插入目标值并保持有序的第一个和最后一个位置**。
- lower_bound: Searches for the first element in the partitioned range [first, last) which is **not ordered before** value.
- upper_bound: Searches for the first element in the partitioned range [first, last) which is **ordered after** value.

循环不变量只需要修改保持过程，将$n>a[mid]$修改为$n>=a[mid]$
1. 如果$n >= a[mid]$，n的上界在mid右边且不包括mid，区间$[mid + 1, r)$满足循环不变量
2. 如果$n < a[mid]$，n的上界在mid左边且包括mid，区间$[l, mid)$满足循环不变量

```go
func BinarySearchUpperBound(arr []int, n int) int {
	low := 0
	high := len(arr)
	for low < high {
		mid := low + (high-low)>>1
		if n >= arr[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
```

另外在以上的定义下，N的上界与N+1的下界位置相同，可以直接用下界换算。
```go
func BinarySearchUpperBound(arr []int, n int) int {
	return BinarySearchLowerBound(arr, n + 1)
}
```

总结一下二分查找的循环不变量四要素，对有序数组$a[l,r)$
- $Variant$: 目标值下界所在的区间
- $Start$: $l=0,r=len(a)$
- $Loop$: 如果$n>a[mid]$，则n的下界在mid右边且不包括mid，反之在左边
- $End$: $l=r$

## Reference
- [二分查找有几种写法？它们的区别是什么？ - Jason Li的回答 - 知乎](https://www.zhihu.com/question/36132386/answer/530313852)
- [算法导论](https://jingyuexing.github.io/Ebook/Algorithm/%E7%AE%97%E6%B3%95%E5%AF%BC%E8%AE%BA.pdf)